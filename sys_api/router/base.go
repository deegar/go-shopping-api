package router

import (
	"github.com/gin-gonic/gin"
	"sys_api/api"
)

func InitBaseRouters(Router *gin.RouterGroup) {
	userRouter := Router.Group("/base")
	{
		userRouter.GET("captcha", api.GetCaptcha)
	}
}
