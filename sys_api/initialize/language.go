//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"sys_api/global"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitI18n() {
	var langTag language.Tag
	var filePath string
	if global.ServerConfig.Lang == "zh" {
		langTag = language.Chinese
		filePath = fmt.Sprintf("%ssys_api/language/active.zh.toml", viper.GetString("GO_SHOPPING_API"))
	} else if global.ServerConfig.Lang == "en" {
		langTag = language.English
		filePath = fmt.Sprintf("%ssys_api/language/active.en.toml", viper.GetString("GO_SHOPPING_API"))
	}
	bundle := i18n.NewBundle(langTag)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(filePath)

	global.Lang = i18n.NewLocalizer(bundle, global.ServerConfig.Lang)
}
