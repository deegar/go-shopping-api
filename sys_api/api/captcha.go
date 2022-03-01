package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	cp := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	captcha, b64s, err := cp.Generate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":        "验证成功",
		"captcha_id": captcha,
		"image":      b64s,
	})
	fmt.Println(b64s)
	return
}
