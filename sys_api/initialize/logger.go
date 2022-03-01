//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package initialize

import "go.uber.org/zap"

func InitZapLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("create zap logger fail")
	}
	zap.ReplaceGlobals(logger)
}
