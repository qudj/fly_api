package config

import (
	"github.com/qudj/fly_lib/models/proto/fcc_serv"
	"github.com/qudj/fly_lib/models/proto/fly_starling_serv"
	"google.golang.org/grpc"
)

var (
	FccRpcClient      fcc_serv.FccServiceClient
	StarlingRpcClient fly_starling_serv.StarlingServiceClient
)

func InitGRPClient() {
	conn1, err := grpc.Dial(Global.Host.FccRpcHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	FccRpcClient = fcc_serv.NewFccServiceClient(conn1)

	conn2, err := grpc.Dial(Global.Host.StarlingRpcHost, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	StarlingRpcClient = fly_starling_serv.NewStarlingServiceClient(conn2)
}
