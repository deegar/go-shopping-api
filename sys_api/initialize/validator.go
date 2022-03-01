//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import (
	"go.uber.org/zap"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	"sys_api/global"
	myvalidator "sys_api/validator"
)

func InitValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New()
		enT := en.New()

		uni := ut.New(zhT, zhT, enT)
		global.ValidatorErrorTrans, ok = uni.GetTranslator(global.ServerConfig.Lang)
		if !ok {
			zap.S().Errorf("uni.GetTranslator(%s)", global.ServerConfig.Lang)
		}
		var err error
		switch global.ServerConfig.Lang {
		case "en":
			err = en_translations.RegisterDefaultTranslations(v, global.ValidatorErrorTrans)
		case "zh":
			err = zh_translations.RegisterDefaultTranslations(v, global.ValidatorErrorTrans)
		default:
			err = zh_translations.RegisterDefaultTranslations(v, global.ValidatorErrorTrans)
		}
		if err != nil {
			zap.S().Errorf(global.GetMessage("FailRegisterTrans"))
		}

	} else {
		zap.S().Error(global.GetMessage("FailRegisterTrans"))
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.ValidatorErrorTrans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
}
