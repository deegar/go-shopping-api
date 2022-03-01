//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"encoding/json"
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"sys_api/global"
)

// load the config data from local config.yaml file.
// Make sure that the environment variable GO_SHOPPING_SRV is set to the project directory.

func InitNacosConfig() {
	v := viper.New()
	viper.AutomaticEnv() //get the value from system env, we should set the project path to the GO_SHOPPING key
	v.SetConfigFile(fmt.Sprintf("%ssys_api/config.yaml",
		viper.GetString("GO_SHOPPING_API")))
	err := v.ReadInConfig()
	if err != nil {
		zap.S().Fatal("viper cannot load Nacos config")
	}

	err = v.Unmarshal(&global.NacosConfig)
	if err != nil {
		zap.S().Fatal("fail to unmarshal Nacos config")
	}
}

// initialize the system configuration from nacos and listen the config changes

func InitServerConfigFromNacos() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              global.NacosConfig.NacosClientConfig.LogDir,
		CacheDir:            global.NacosConfig.NacosClientConfig.CacheDir,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		zap.S().Fatal("fail to get nacos config. error:", err.Error())
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("fail to load sys config from Nacos： %s", err.Error())
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			zap.S().Infof("nacos config file: %s has been changed", global.NacosConfig.DataId)
		},
	})

	if err != nil {
		zap.S().Error("fail to monitor the changes on nacos")
	}

	zap.S().Info("initialize system configuration successfully")
}
