//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, &gin.H{
		"msg": "service is healthy",
	})
}
