package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sys_api/form"
	"sys_api/global"
	"sys_api/global/response"
	"sys_api/middleware"
	"sys_api/model"
	"sys_api/proto"
)

func ConvertErrorToHttp(err error, ctx *gin.Context) {
	if state, ok := status.FromError(err); ok {
		switch state.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": global.GetMessage("InfoNotFound"),
			})
		case codes.Internal:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": global.GetMessage("InternalError"),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": global.GetMessage("InvalidArgument"),
			})
		case codes.Unavailable:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": global.GetMessage("Unavailable"),
			})
		case codes.AlreadyExists:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": global.GetMessage("AlreadyExist"),
			})
		}
		return
	}
}

func HandleValidationError(c *gin.Context, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errors.Translate(global.ValidatorErrorTrans),
		})
	} else {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
	}
}

func GetUserList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("page_num", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	list, err := global.ServerClient.GetUserList(context.Background(), &proto.PageInfo{
		PageNum:  uint32(pageNum),
		PageSize: uint32(pageSize),
	})

	if err != nil {
		zap.S().Error(err)
		ConvertErrorToHttp(err, ctx)
		return
	}

	result := make([]response.UserResponse, 0)

	for _, value := range list.Data {
		result = append(result, response.UserResponse{
			Id:       int32(value.Id),
			NickName: value.NickName,
			Mobile:   value.Mobile,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func Register(c *gin.Context) {
	registerForm := form.RegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidationError(c, err)
		return
	}

	//if !store.Verify(registerForm.CaptchaId, registerForm.Captcha, true) {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"msg": global.GetMessage("WrongCaptcha"),
	//	})
	//	return
	//}

	_, err := global.ServerClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName:  &registerForm.NickName,
		LoginName: registerForm.LoginName,
		Mobile:    &registerForm.Mobile,
		Email:     &registerForm.Email,
		Password:  registerForm.Password,
	})

	if err != nil {
		zap.S().Errorf("%s, data:%v", global.GetMessage("FailRegister"), c)
		ConvertErrorToHttp(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func PasswordLogin(c *gin.Context) {
	loginForm := form.PasswordLoginForm{}
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		HandleValidationError(c, err)
		return
	}

	if !store.Verify(loginForm.CaptchaId, loginForm.Captcha, true) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": global.GetMessage("WrongCaptcha"),
		})
		return
	}

	resp, err := global.ServerClient.GetUserByLoginName(context.Background(),
		&proto.BaseRequest{Msg: loginForm.LoginName})
	if err != nil {
		zap.S().Error(err)
		ConvertErrorToHttp(err, c)
		return
	}

	result, err := global.ServerClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:          loginForm.Password,
		EncryptedPassword: resp.Password,
	})

	if err != nil {
		zap.S().Error(err)
		ConvertErrorToHttp(err, c)
		return
	}

	if result.Success == false {
		c.JSON(http.StatusOK, gin.H{
			"msg": global.GetMessage("WrongPassword"),
		})
		return
	} else {
		j := middleware.NewJWT()
		token := j.NewToken(model.CustomClaims{
			ID:          uint32(resp.Id),
			NickName:    resp.NickName,
			AuthorityId: uint(resp.RoleId),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				Issuer:    "Ryan",
			},
		})
		c.JSON(http.StatusOK, gin.H{
			"msg":     global.GetMessage("LoginSuccess"),
			"x-token": token,
		})
		return
	}
}
