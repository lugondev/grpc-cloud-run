package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
	pb "waas-service/pb"

	"google.golang.org/grpc"
)

var (
	addr = os.Args[1]
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	var opts []grpc.DialOption
	systemRoots, err := x509.SystemCertPool()
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPingClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Msg(ctx, &pb.MsgRequest{})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	log.Printf("%s", r.GetMessage())
}
