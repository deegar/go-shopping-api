//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package global

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
)

func GetMessage(id string) string {
	message, err := Lang.Localize(&i18n.LocalizeConfig{
		MessageID: id,
	})

	if err != nil {
		zap.S().Errorf("i18n cannot find message with ID: %s", id)
	}

	return message
}
