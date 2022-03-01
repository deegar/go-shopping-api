//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"sys_api/global"
	"sys_api/model"
)

type JWT struct {
	SigningKey []byte
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("g-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登录",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.TokenParseToClaims(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{[]byte(global.ServerConfig.JWTConfig.SigningKey)}
}

func (j *JWT) NewToken(claims model.CustomClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.SigningKey)
	if err != nil {
		zap.S().Errorf("Token签名失败")
	}
	fmt.Println(signedToken)
	return signedToken
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func (j *JWT) TokenParseToClaims(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		fmt.Println(err)
		//zap.S().Panicw("转换claims失败", "err", err.Error())
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenExpired
			} else {
				return nil, TokenNotValidYet
			}
		} else {
			return nil, TokenInvalid
		}
	}
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
		return j.NewToken(*claims), nil
	}
	return "", TokenInvalid
}
