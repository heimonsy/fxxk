package main

import (
	"flag"
	"log"
	"net"

	"github.com/heimonsy/fxxk"
	"google.golang.org/grpc"
)

var (
	listen      = flag.String("listen", ":8081", "")
	proxyListen = flag.String("proxy-listen", ":8082", "")
)

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	svc := fxxk.NewService()

	svc.Register(grpcServer)
	go svc.ProxyAndListen(*proxyListen)

	err = grpcServer.Serve(ln)
	if err != nil {
		log.Fatalln("grpc serve error:", err)
	}
}
