package dto

type UserLoginReq struct {
	Username  string `label:"用户名" binding:"required,min=3,max=20"`
	Password  string `label:"密码" binding:"required,min=6,max=20"`
	Captcha   string `label:"验证码" binding:"required,min=4,max=6"`
	CaptchaId string `label:"验证码ID" binding:"required" json:"captcha_id"`
}

type UserLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         struct {
		Username string `json:"username"`
		UserId   uint   `json:"user_id"`
	} `json:"user"`
}

type UserListReq struct{}

type CreateUserReq struct {
	Username string `label:"用户名" binding:"required,min=3,max=20"`
	Password string `label:"密码" binding:"required,min=6,max=20"`
}
