package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"sys_api/config"
	"sys_api/proto"
)

var (
	ValidatorErrorTrans ut.Translator            // translate validator error messages
	ServerConfig        = &config.ServerConfig{} // store all server configuration
	ServerClient        proto.SystemClient       // protobuf client for connecting server
	NacosConfig         = &config.NacosConfig{}  // store nacos config data
	Lang                *i18n.Localizer          // use to support multiple languages on messages
)
