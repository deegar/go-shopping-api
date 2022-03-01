package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"

	"sys_api/global"
	"sys_api/initialize"
	"sys_api/util/register/consul"
)

func main() {
	//init logger
	initialize.InitZapLogger()
	//init Nacos config
	initialize.InitNacosConfig()
	//init sys config
	initialize.InitServerConfigFromNacos()
	//init multiple languages support
	initialize.InitI18n()
	//init validator
	initialize.InitValidator()
	//init gin server
	routers := initialize.Routers()

	//register health service
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulConfig.Host, int(global.ServerConfig.ConsulConfig.Port))
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, int(global.ServerConfig.Port), global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic(global.GetMessage("FailRegisterConsul"), " err: ", err.Error())
	}

	//init service client
	initialize.InitServerClient()

	err = routers.Run(fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port))
	if err != nil {
		zap.S().Fatal(global.GetMessage("GinStartFail"), "   error:", err.Error())
	}

}
