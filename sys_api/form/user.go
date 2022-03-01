package form

type PasswordLoginForm struct {
	LoginName string `form:"login_name" json:"login_name" binding:"required, max=10"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	NickName  string `form:"nick_name" json:"nick_name" binding:"required,max=10"`
	LoginName string `form:"login_name" json:"login_name" binding:"required,max=10"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
