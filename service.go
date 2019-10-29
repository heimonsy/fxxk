package fxxk

import (
	fmt "fmt"
	"io"
	"log"
	"net"
	"sync"

	grpc "google.golang.org/grpc"
)

// ConnectionHandler 处理进来的连接
type ConnectionHandler interface {
	HandleAndListen(addr string)
}

type Service struct {
	mu      *sync.Mutex
	clients map[string]*connectClient

	newTunelSig *sync.Cond
}

// NewService
func NewService() *Service {
	return &Service{mu: &sync.Mutex{}, newTunelSig: sync.NewCond(&sync.Mutex{})}
}

func (s *Service) Register(gs *grpc.Server) {
	RegisterFxxkServer(gs, s)
}

func (s *Service) getTunnel() (stream *tunnelStream) {
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
		if err := c.stream.Send(&Command{Type: Command_NEW_TUNEL}); err != nil {
			log.Printf("send command to client %s error: %s", c.id, err)
			// TODO close client
			continue
		}
		succeed = true
	}

	if !succeed {
		log.Println("no clients")
	}
	return
}

func (s *Service) newTunelOfClientID(clientID string, tunnel *tunnelStream) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	client, exists := s.clients[clientID]
	if !exists {
		return fmt.Errorf("client id %s not exists", clientID)
	}
	client.idleTunnels = append(client.idleTunnels, tunnel)

	s.newTunelSig.Signal()
	return nil
}

func (s *Service) Connect(
	req *ConnectRequest,
	stream Fxxk_ConnectServer,
) error {
	client := &connectClient{id: req.ClientId, done: make(chan struct{}), stream: stream}

	s.mu.Lock()
	if _, exists := s.clients[client.id]; exists {
		s.mu.Unlock()
		return fmt.Errorf("client id exists")
	}
	s.clients[client.id] = client
	s.mu.Unlock()

	<-client.done
	return nil
}

func (s *Service) ProxyTunel(
	stream Fxxk_ProxyTunelServer,
) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	id, ok := req.Req.(*TunnelRequest_ClientId)
	if !ok {
		return fmt.Errorf("need sent client id first")
	}
	ts := &tunnelStream{Fxxk_ProxyTunelServer: stream, done: make(chan struct{})}

	if err := s.newTunelOfClientID(id.ClientId, ts); err != nil {
		return err
	}

	<-ts.done
	return nil
}

func (s *Service) HandleAndListen(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("new connection from handle:", tcpConn.RemoteAddr())
		go s.controls(tcpConn)
	}
}

func (s *Service) waitAndGetTunnel() (stream *tunnelStream) {
	for {
		stream := s.getTunnel()
		if stream == nil && !s.acquireTunnel() {
			return nil
		}
		s.newTunelSig.Wait()
	}
}

func (s *Service) controls(conn net.Conn) {
	stream := s.waitAndGetTunnel()
	if stream == nil {
		conn.Close()
		fmt.Println("cant get tunnel for conn:", conn.RemoteAddr())
		return
	}
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	rw := stream.ReadWriter()
	s.atob(rw, conn, done1)
	s.atob(conn, rw, done2)

	select {
	case <-done1:
	case <-done2:
	}
	stream.Close()
	conn.Close()
}

func (s *Service) atob(a io.Reader, b io.Writer, done chan struct{}) {
	defer func() {
		close(done)
	}()

	var buf [1024 * 4]byte
	for {
		var err error
		n, err := a.Read(buf[:])
		if err != nil {
			fmt.Println("read conn error")
			return
		}

		var writed int
		for writed != n {
			w, err := b.Write(buf[writed:n])
			if err != nil {
				fmt.Println("write conn error")
				return
			}
			writed += w
		}
	}
}
