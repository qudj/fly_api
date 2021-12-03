package config

import (
	"github.com/qudj/fly_lib/models/fly_conf"
	"google.golang.org/grpc"
)


var (
	FccRpcClient  fly_conf.FccServiceClient
)

func InitGRPClient() {
	conn, err := grpc.Dial(Global.Host.FccRpcHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	FccRpcClient = fly_conf.NewFccServiceClient(conn)
}