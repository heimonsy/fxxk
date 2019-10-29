package fxxk

import (
	context "context"
	fmt "fmt"
	"log"
	"net"

	"github.com/rs/xid"
	grpc "google.golang.org/grpc"
)

// Client
type Client struct {
	id   string
	cli  FxxkClient
	done chan error

	targetAddr string
}

func NewClient(conn *grpc.ClientConn, targetAddr string) *Client {
	return &Client{
		cli:  NewFxxkClient(conn),
		id:   xid.New().String(),
		done: make(chan error, 1),
	}
}

func (c *Client) Start(initTunnels int) error {
	connectStream, err := c.initConnect()
	if err != nil {
		return err
	}
	go c.handleConnect(connectStream)

	return <-c.done
}

func (c *Client) initConnect() (Fxxk_ConnectClient, error) {
	connectStream, err := c.cli.Connect(context.Background(), &ConnectRequest{
		ClientId: c.id,
	})
	if err != nil {
		return nil, err
	}
	cmd, err := connectStream.Recv()
	if err != nil {
		return nil, err
	}
	if cmd.Type != Command_PING {
		return nil, fmt.Errorf("unknow service command: expecting PING")
	}
	return connectStream, nil
}

func (c *Client) close(err error) {
	select {
	case c.done <- err:
	default:
	}
	return
}

func (c *Client) handleConnect(connectStream Fxxk_ConnectClient) {
	for {
		cmd, err := connectStream.Recv()
		if err != nil {
			c.close(err)
			return
		}
		switch cmd.Type {
		case Command_PING:
		case Command_NEW_TUNEL:
		case Command_CLOSE:
			c.close(connectStream.CloseSend())
			return
		default:
			log.Println("unknow command:", cmd.Type)
		}
	}

}

func (c *Client) startNewTunnel() {
	stream, err := c.cli.Tunel(context.Background())
	if err != nil {
		c.close(err)
		return
	}

	err = stream.Send(&TunnelRequest{
		Req: &TunnelRequest_ClientId{
			ClientId: c.id,
		},
	})
	if err != nil {
		c.close(err)
		return
	}

	v, err := stream.Recv()
	if err != nil {
		c.close(err)
		return
	}
	if len(v.Data) != 0 {
		log.Printf("expect empty data from server")
		stream.CloseSend()
		return
	}

	conn, err := net.Dial("tcp", c.targetAddr)
	if err != nil {
		log.Printf("dial %s error: %v", c.targetAddr, err)
		stream.CloseSend()
		return
	}

	done1 := make(chan struct{})
	done2 := make(chan struct{})
	rw := newWrapStreamFromTunnelClient(stream)
	atob(rw, conn, done1)
	atob(conn, rw, done2)

	select {
	case <-done1:
	case <-done2:
	}
	stream.CloseSend()
	conn.Close()
}
