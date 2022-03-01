package initialize

import (
	"github.com/gin-gonic/gin"
	"sys_api/api"

	"sys_api/middleware"
	. "sys_api/router"
)

func Routers() *gin.Engine {
	routerDefault := gin.Default()
	routerDefault.Use(middleware.Cros())
	routerDefault.GET("/health", api.HealthCheck)

	apiRouter := routerDefault.Group("/sys/v1")

	InitUserRouters(apiRouter)
	InitBaseRouters(apiRouter)

	return routerDefault
}
