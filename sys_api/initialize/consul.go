//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"sys_api/global"
	"sys_api/proto"
)

func InitServerClient() {
	consulInfo := global.ServerConfig.ConsulConfig
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port,
			global.ServerConfig.SysSrvConfig.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal(global.GetMessage("FailGetService"))
	}

	client := proto.NewSystemClient(conn)
	global.ServerClient = client
}
