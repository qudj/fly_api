package config

import (
	"github.com/qudj/fly_lib/models/proto/fcc_serv"
	"google.golang.org/grpc"
)


var (
	FccRpcClient  fcc_serv.FccServiceClient
)

func InitGRPClient() {
	conn, err := grpc.Dial(Global.Host.FccRpcHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	FccRpcClient = fcc_serv.NewFccServiceClient(conn)
}