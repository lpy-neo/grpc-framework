package main

import (
	"log"

	api "github.com/lpy-neo/grpc-framework/api/proto/echo"
	svr "github.com/lpy-neo/grpc-framework/echo/pkg/echo"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/service/grpc"
)

func main() {
	service := grpc.NewService(
		micro.Name("echo"),
	)

	service.Init()

	api.RegisterEchoServiceServer(service.Server(), svr.NewEchoServiceServer())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
package main

import (
	// Use "golang.org/x/net/context" for Golang version <= 1.6

	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/lpy-neo/grpc-framework/api/proto/echo"
	svr "github.com/lpy-neo/grpc-framework/echo/pkg/echo"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = "localhost:9090"
)

func main() {
	lis, err := net.Listen("tcp", grpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, svr.NewEchoServiceServer())
	grpcServer.Serve(lis)
}
*/
