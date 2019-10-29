package fxxk

import (
	"bytes"
	fmt "fmt"
	"io"
)

type connectClient struct {
	id          string
	done        chan struct{}
	stream      Fxxk_ConnectServer
	idleTunnels []*tunnelStream
}

type tunnelStream struct {
	Fxxk_ProxyTunelServer
	done chan struct{}
}

func (ts *tunnelStream) Close() {
	// TODO
}

func (ts *tunnelStream) ReadWriter() io.ReadWriter {
	return &wrapTunnelStream{
		Fxxk_ProxyTunelServer: ts.Fxxk_ProxyTunelServer,
		buf:                   bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

type wrapTunnelStream struct {
	Fxxk_ProxyTunelServer
	buf *bytes.Buffer
}

func (w *wrapTunnelStream) Read(p []byte) (n int, err error) {
	if w.buf.Len() > 0 {
		return w.buf.Read(p)
	}

	req, err := w.Recv()
	if err != nil {
		return
	}
	var data []byte
	{
		v, ok := req.Req.(*TunnelRequest_Data)
		if !ok {
			return 0, fmt.Errorf("invalid client request: expect data")
		}
		data = v.Data.Data
	}

	n = copy(p, data)
	if len(data) > n {
		w.buf.Reset()
		w.buf.Write(data[n:])
	}
	return n, nil
}

func (w *wrapTunnelStream) Write(p []byte) (n int, err error) {
	err = w.Send(&TunnelPackage{
		Data: p,
	})
	if err != nil {
		return
	}

	return len(p), nil
}
