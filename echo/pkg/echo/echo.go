package echo

import (
	"context"

	api "github.com/lpy-neo/grpc-framework/api/proto/echo"
	"github.com/lpy-neo/grpc-framework/api/proto/shared/types"
)

func NewEchoServiceServer() api.EchoServiceServer {
	return &echoServiceServer{}
}

type echoServiceServer struct{}

func (e echoServiceServer) Echo(_ context.Context, in *types.StringMessage) (rsp *types.StringMessage, err error) {
	rsp.Value = "Echo " + in.Value
	return
}
