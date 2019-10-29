package fxxk

import (
	"bytes"
	"io"
	"sync"
)

type connectClient struct {
	id          string
	done        chan struct{}
	mu          *sync.Mutex
	stream      Fxxk_ConnectServer
	idleTunnels []*tunnelServerStream
}

func newConnectClient(id string, stream Fxxk_ConnectServer) *connectClient {
	return &connectClient{
		id:     id,
		stream: stream,
		done:   make(chan struct{}),
		mu:     &sync.Mutex{},
	}
}

func (c *connectClient) Send(cmd *Command) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.stream.Send(cmd)
}

type tunnelServerStream struct {
	Fxxk_TunelServer
	done chan struct{}
}

func (ts *tunnelServerStream) Close() {
	close(ts.done)
}

func (ts *tunnelServerStream) ReadWriter() io.ReadWriter {
	return &wrapStreamIOReadWriter{
		stream: &tunnelServer{stream: ts.Fxxk_TunelServer},
		buf:    bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}
