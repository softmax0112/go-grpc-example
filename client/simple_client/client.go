package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/EDDYCJY/go-grpc-example/pkg/gtls"
	pb "github.com/EDDYCJY/go-grpc-example/proto"
)

const PORT = "9001"

func main() {
	tlsClient := gtls.Client{
		ServerName: "go-grpc-example",
		CaFile:     "../../conf/ca.pem",
		CertFile:   "../../conf/client/client.pem",
		KeyFile:    "../../conf/client/client.key",
	}

	c, err := tlsClient.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
