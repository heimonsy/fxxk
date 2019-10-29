package fxxk

import (
	fmt "fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
)

// ConnectionHandler 处理进来的连接
type ConnectionHandler interface {
	ProxyAndListen(addr string)
}

type Service struct {
	mu      *sync.Mutex
	clients map[string]*connectClient

	newTunelSig *sync.Cond
}

// NewService
func NewService() *Service {
	return &Service{
		mu:          &sync.Mutex{},
		clients:     make(map[string]*connectClient),
		newTunelSig: sync.NewCond(&sync.Mutex{}),
	}
}

func (s *Service) Register(gs *grpc.Server) {
	RegisterFxxkServer(gs, s)
}

func (s *Service) getTunnel() (stream *tunnelServerStream) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, c := range s.clients {
		if len(c.idleTunnels) > 0 {
			stream = c.idleTunnels[len(c.idleTunnels)-1]
			c.idleTunnels = c.idleTunnels[:len(c.idleTunnels)-1]
			return
		}
	}
	return
}

func (s *Service) acquireTunnel() (succeed bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, c := range s.clients {
		if err := c.Send(&Command{Type: Command_NEW_TUNEL}); err != nil {
			log.Printf("send command to client %s error: %s", c.id, err)
			c.Close()
			continue
		}
		succeed = true
	}

	if !succeed {
		log.Println("[server] no clients")
	}
	return
}

func (s *Service) newTunelOfClientID(clientID string, tunnel *tunnelServerStream) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	client, exists := s.clients[clientID]
	if !exists {
		return fmt.Errorf("client id %s not exists", clientID)
	}
	client.idleTunnels = append(client.idleTunnels, tunnel)

	s.newTunelSig.L.Lock()
	s.newTunelSig.Signal()
	s.newTunelSig.L.Unlock()
	return nil
}

func (s *Service) Connect(
	req *ConnectRequest,
	stream Fxxk_ConnectServer,
) error {
	client := newConnectClient(req.ClientId, stream)

	s.mu.Lock()
	if _, exists := s.clients[client.id]; exists {
		s.mu.Unlock()
		return fmt.Errorf("client id exists")
	}
	s.clients[client.id] = client
	s.mu.Unlock()
	log.Println("[server] new client: ", client.id)

	if err := client.Send(&Command{Type: Command_PING}); err != nil {
		return err
	}

	go s.keepalive(client)
	<-client.done

	s.mu.Lock()
	delete(s.clients, client.id)
	s.mu.Unlock()
	return nil
}

func (s *Service) keepalive(client *connectClient) {
	t := time.NewTicker(time.Second * 10)
	for range t.C {
		if err := client.Send(&Command{Type: Command_PING}); err != nil {
			client.Close()
			break
		}
	}

}

func (s *Service) Tunel(
	stream Fxxk_TunelServer,
) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	id, ok := req.Req.(*TunnelRequest_ClientId)
	if !ok {
		return fmt.Errorf("need sent client id first")
	}
	ts := &tunnelServerStream{Fxxk_TunelServer: stream, done: make(chan struct{})}

	if err := s.newTunelOfClientID(id.ClientId, ts); err != nil {
		return err
	}
	log.Println("[server] new tunnel from client:", id.ClientId)

	<-ts.done
	return nil
}

func (s *Service) ProxyAndListen(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("[server]", tcpConn.RemoteAddr(), "new connection from handle")
		go s.controls(tcpConn)
	}
}

func (s *Service) waitAndGetTunnel() (stream *tunnelServerStream) {
	for {
		stream := s.getTunnel()
		if stream != nil {
			return stream
		}

		if !s.acquireTunnel() {
			return nil
		}
		s.newTunelSig.L.Lock()
		s.newTunelSig.Wait()
		s.newTunelSig.L.Unlock()
	}
}

func (s *Service) controls(conn net.Conn) {
	stream := s.waitAndGetTunnel()
	if stream == nil {
		conn.Close()
		log.Println("[server] cant get tunnel for conn:", conn.RemoteAddr())
		return
	}

	// 与客户端约定发送一个空包，表示准备开始使用这个 stream
	if err := stream.Send(&TunnelPackage{}); err != nil {
		conn.Close()
		log.Println("[server] send empty package error", err)
		return
	}
	log.Println("[server]", conn.RemoteAddr(), "empty package sended")

	done1 := make(chan struct{})
	done2 := make(chan struct{})
	rw := stream.ReadWriter()
	go atob(rw, conn, done1)
	go atob(conn, rw, done2)

	select {
	case <-done1:
	case <-done2:
	}
	stream.Close()
	conn.Close()
}

func atob(a io.Reader, b io.Writer, done chan struct{}) {
	defer func() {
		close(done)
	}()

	var buf [1024 * 4]byte
	for {
		var err error
		n, err := a.Read(buf[:])
		if err != nil {
			if err != io.EOF {
				log.Println("read conn error", err)
			}
			return
		}

		var writed int
		for writed != n {
			w, err := b.Write(buf[writed:n])
			if err != nil {
				log.Println("write conn error", err)
				return
			}
			writed += w
		}
	}
}
