package router

import (
	"github.com/gin-gonic/gin"

	"sys_api/api"
)

func InitUserRouters(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user")
	{
		//userRouter.GET("list", middleware.JWTAuth(), api.GetUserList)
		userRouter.GET("list", api.GetUserList)
		userRouter.POST("register", api.Register)
		userRouter.POST("pwd_login", api.PasswordLogin)
	}
}
