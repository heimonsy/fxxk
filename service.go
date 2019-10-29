package fxxk

import (
	fmt "fmt"
	"sync"

	grpc "google.golang.org/grpc"
)

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

func (s *Service) newTunelOfClientID(clientID string, tunel tunelStream) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	client, exists := s.clients[clientID]
	if !exists {
		return fmt.Errorf("client id %s not exists", clientID)
	}
	client.idleTunel = append(client.idleTunel, tunel)

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
	ts := tunelStream{Fxxk_ProxyTunelServer: stream, done: make(chan struct{})}

	if err := s.newTunelOfClientID(id.ClientId, ts); err != nil {
		return err
	}

	<-ts.done
	return nil
}
