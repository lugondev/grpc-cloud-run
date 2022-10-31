package main

import (
	"context"
	"fmt"
	ggrpc "google.golang.org/grpc"
	"log"
	"net"
	"time"
	"waas-service/auth/jwt/jose"
	"waas-service/config"
	"waas-service/grpc"
	"waas-service/pb"
)

// server is used to implement proto.GreeterServer.
type server struct {
	pb.PingServer
	//store *db.SQLStore
}

// Ping implements pingServer.Ping
func (s *server) Ping(ctx context.Context, _ *pb.PingRequest) (*pb.PongReply, error) {
	//group, err := s.store.CreateGroup(ctx, db.CreateGroupParams{
	//	Owner:     "lugon1",
	//	GroupName: "group1",
	//})
	//if err != nil {
	//	return nil, err
	//}
	//group, err := s.store.GetGroup(ctx, 1)
	//if err != nil {
	//	return nil, err
	//}
	//log.Printf("group: %v", group)
	return &pb.PongReply{Message: fmt.Sprintf("Pong at %s", time.Now())}, nil
}

func (s *server) Msg(_ context.Context, _ *pb.MsgRequest) (*pb.MsgReply, error) {
	return &pb.MsgReply{Message: fmt.Sprintf("Reply at: %s", time.Now())}, nil
}

func init() {
	config.LoadConfiguration()
}

func main() {
	//auth0Management := auth0.NewAuth0Management()
	appConfig := jose.GetAppConfig()
	//dbStore, err := config.NewDB()
	//if err != nil {
	//	log.Fatalf("failed to create DB store: %v", err)
	//}

	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", appConfig.GetServerPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := ggrpc.NewServer(
		ggrpc.UnaryInterceptor(grpc.UnaryInterceptor))

	pb.RegisterPingServer(s, &server{
		//store: dbStore,
	})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
