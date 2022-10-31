package main_test

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
	"waas-service/pb"

	"google.golang.org/grpc"
	grpcMetadata "google.golang.org/grpc/metadata"
)

var accessToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Im5wUGYtZHpaQktfRVJJNTJSNG9pbyJ9.eyJodHRwczovL2JhbXMuZGV2L3JiYWMiOnsiZW1haWxfdmVyaWZpZWQiOnRydWUsInJvbGVzIjpbIkFETUlOIiwiQkFNLVVzZXJzIiwiREFPLU1hbmFnZXIiXSwidGVuYW50X2lkIjoiYmFtcy1kZXYiLCJ1c2VybmFtZSI6ImF1dGgwfDYyYzQ3MTBlYzNlZjU2YTFkOTNhNmJlNiJ9LCJpc3MiOiJodHRwczovL2F1dGguYmFtcy5kZXYvIiwic3ViIjoiYXV0aDB8NjJjNDcxMGVjM2VmNTZhMWQ5M2E2YmU2IiwiYXVkIjpbImh0dHBzOi8vYXV0aC5iYW1zLmRldi9pZCIsImh0dHBzOi8vYmFtLWRldi51cy5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjY1NTYwNjE1LCJleHAiOjE2NjU2NDcwMTUsImF6cCI6IkRQYlVTUVQ2NVVUU0hlVVowaGtPUjE4Z3hLSnViaWh0Iiwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSBlbmNyeXB0OmV0aGVyZXVtIHByb3h5Om5vZGVzIHJlYWQ6ZXRoZXJldW0gc2lnbjpldGhlcmV1bSB3cml0ZTpldGhlcmV1bSBvZmZsaW5lX2FjY2VzcyIsImd0eSI6InBhc3N3b3JkIiwicGVybWlzc2lvbnMiOlsiZW5jcnlwdDpldGhlcmV1bSIsInByb3h5Om5vZGVzIiwicmVhZDpldGhlcmV1bSIsInNpZ246ZXRoZXJldW0iLCJ3cml0ZTpldGhlcmV1bSJdfQ.LXAfveTPbgLV8e4ve-6obUvjmhMoqBVZpWsNzlFQqVB3nvOpVprk03VPQaCmIq-DwUaFR0j3rL3lj-MzS3bZmzZv-_VMNPTtAll_NfWBZ7l0QwZb_1wgG6RdYZUeMXS3IBuZaEYG0Im8spHzX6IPCUlNIKm9I1vnrc1j4-WapDx-CgGVJFwNjAnboPhjySPcZ11IG36-iREpB-x0VNLm2hE1QulRCMx6n36Sfyxazkb0qiJwdqG6Zlo3Ckzc6aSf3BWntdtcc0jClmHKgITesjhxzNMJGvCcWVp5F5FlhvUu_QOyRS-5vZWKng3JgPmgRvNWhOO9qWivx2Z9q1N1jQ"

//var accessToken = "fake jwt"

func TestStreamGRPC(t *testing.T) {
	// dial server
	conn, err := grpc.Dial("waas-grpc-czljur1g.uc.gateway.dev:443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ctx = grpcMetadata.AppendToOutgoingContext(ctx, "access_token", accessToken)

	// create stream
	client := pb.NewPingClient(conn)
	pongReply, err := client.Msg(ctx, &pb.MsgRequest{})
	if err != nil {
		log.Panic(err)
		return
	}
	log.Printf("Pong reply: %s", pongReply.Message)
}
