package fxxk

import (
	"bytes"
	"io"
	"log"
	"sync"
)

type connectClient struct {
	id     string
	done   chan struct{}
	mu     *sync.Mutex
	stream Fxxk_ConnectServer

	// TODO 重新实现 idle tunnels，方便优雅的关闭
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

func (c *connectClient) Close() {
	select {
	case <-c.done:
	default:
		c.mu.Lock()
		defer c.mu.Unlock()
		for _, v := range c.idleTunnels {
			v.Close()
		}
		close(c.done)
		log.Println("[server] client closed:", c.id)
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
	select {
	case <-ts.done:
	default:
		close(ts.done)
	}
}

func (ts *tunnelServerStream) ReadWriter() io.ReadWriter {
	return &wrapStreamIOReadWriter{
		stream: &tunnelServer{stream: ts.Fxxk_TunelServer},
		buf:    bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}
