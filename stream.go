package fxxk

import (
	"bytes"
	fmt "fmt"
)

type streamRecvSend interface {
	Recv() ([]byte, error)
	Send([]byte) error
}

type tunnelServer struct {
	stream Fxxk_TunelServer
}

func (ts *tunnelServer) Recv() ([]byte, error) {
	req, err := ts.stream.Recv()
	if err != nil {
		return nil, err
	}
	v, ok := req.Req.(*TunnelRequest_Data)
	if !ok {
		return nil, fmt.Errorf("invalid client request: expect data")
	}
	return v.Data.Data, nil

}

func (ts *tunnelServer) Send(p []byte) error {
	return ts.stream.Send(&TunnelPackage{
		Data: p,
	})
}

type tunnelClient struct {
	stream Fxxk_TunelClient
}

func (tc *tunnelClient) Recv() ([]byte, error) {
	req, err := tc.stream.Recv()
	if err != nil {
		return nil, err
	}
	return req.GetData(), nil
}

func (tc *tunnelClient) Send(p []byte) error {
	return tc.stream.Send(&TunnelRequest{
		Req: &TunnelRequest_Data{
			Data: &TunnelPackage{
				Data: p,
			},
		},
	})
}

func newWrapStreamFromTunnelClient(stream Fxxk_TunelClient) *wrapStreamIOReadWriter {
	return &wrapStreamIOReadWriter{
		stream: &tunnelClient{stream: stream},
		buf:    bytes.NewBuffer(make([]byte, 0, 1024)),
	}
}

type wrapStreamIOReadWriter struct {
	stream streamRecvSend
	buf    *bytes.Buffer
}

func (w *wrapStreamIOReadWriter) Read(p []byte) (n int, err error) {
	if w.buf.Len() > 0 {
		return w.buf.Read(p)
	}

	data, err := w.stream.Recv()
	if err != nil {
		return
	}

	n = copy(p, data)
	if len(data) > n {
		w.buf.Reset()
		w.buf.Write(data[n:])
	}
	return n, nil
}

func (w *wrapStreamIOReadWriter) Write(p []byte) (n int, err error) {
	err = w.stream.Send(p)
	if err != nil {
		return 0, nil
	}
	return len(p), nil
}
