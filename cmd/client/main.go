package main

import (
	"flag"
	"log"
	"time"

	"github.com/heimonsy/fxxk"
	"google.golang.org/grpc"
)

var (
	server     = flag.String("server", "127.0.0.1:8081", "")
	targetAddr = flag.String("targetAddr", "127.0.0.1:7071", "")
)

func main() {
	flag.Parse()
	newConn := func() {
		for {
			conn, err := grpc.Dial(*server, grpc.WithInsecure())
			if err != nil {
				log.Println("dial error", err)
				<-time.After(time.Second * 3)
			}
			client := fxxk.NewClient(conn, *targetAddr)
			err = client.Start(5)
			if err != nil {
				log.Println("conn closed:", err)
				<-time.After(time.Second * 3)
			}
			conn.Close()
		}
	}

	for i := 0; i < 5; i++ {
		go newConn()
	}
	ch := make(chan struct{})
	<-ch
}
