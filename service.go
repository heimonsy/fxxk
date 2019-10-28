package fxxk

import (
	"sync"

	grpc "google.golang.org/grpc"
)

type connectClient struct {
	done   chan struct{}
	stream Fxxk_ConnectServer
}

type Service struct {
	mu      *sync.Mutex
	clients []*connectClient
}

func NewService() *Service {
	return &Service{mu: &sync.Mutex{}}
}

func (s *Service) Register(gs *grpc.Server) {
	RegisterFxxkServer(gs, s)
}

func (s *Service) Connect(
	req *ConnectRequest,
	stream Fxxk_ConnectServer,
) error {

	client := &connectClient{done: make(chan struct{}), stream: stream}

	s.mu.Lock()
	s.clients = append(s.clients, client)
	s.mu.Unlock()

	<-client.done
	return nil
}
