package fxxk

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	grpc "google.golang.org/grpc"
)

func testCommonServer(t *testing.T, addr string, started chan struct{}) {
	ls, err := net.Listen("tcp", addr)
	require.NoError(t, err)
	defer ls.Close()
	started <- struct{}{}

	require.NoError(t, err)
	conn, err := ls.Accept()
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		require.NoError(t, err)
		require.Equal(t, "PING-SERVER", string(buf[:n]))

		n, err = conn.Write([]byte("PING-CLIENT"))
		require.Equal(t, n, 11)
		require.NoError(t, err)
	}
}

func testCommonClient(t *testing.T, addr string) {
	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer conn.Close()

	for i := 0; i < 10; i++ {
		n, err := conn.Write([]byte("PING-SERVER"))
		require.Equal(t, n, 11)
		require.NoError(t, err)

		var buf [1024]byte
		n, err = conn.Read(buf[:])
		require.NoError(t, err)
		require.Equal(t, "PING-CLIENT", string(buf[:n]))
	}
}

func TestCommon(t *testing.T) {
	var (
		commonServerAddr = "127.0.0.1:7071"
	)

	started := make(chan struct{}, 16)
	go testCommonServer(t, commonServerAddr, started)
	<-started
	testCommonClient(t, commonServerAddr)
}

func testProxyClient(t *testing.T, server, target string, started chan struct{}) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := NewClient(conn, target)
	started <- struct{}{}
	err = client.Start(5)
	require.NoError(t, err)
}

func testProxyServer(t *testing.T, server, proxyServe string, started chan struct{}) {
	ln, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	svc := NewService()

	svc.Register(grpcServer)
	go svc.ProxyAndListen(proxyServe)

	started <- struct{}{}

	grpcServer.Serve(ln)
}

func TestServiceAll(t *testing.T) {
	var (
		commonServerAddr = "127.0.0.1:7071"
		proxyServer      = "127.0.0.1:8081"
		connectServer    = "127.0.0.1:8082"
	)
	started := make(chan struct{}, 16)

	go testProxyServer(t, connectServer, proxyServer, started)
	<-started
	go testProxyClient(t, connectServer, commonServerAddr, started)
	<-started
	<-time.After(time.Second)

	go testCommonServer(t, commonServerAddr, started)
	<-started
	testCommonClient(t, proxyServer)
}
