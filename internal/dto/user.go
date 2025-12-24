package dto

type UserLoginReq struct {
	Username string
	Password string
}

type UserLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserListReq struct{}

type CreateUserReq struct {
	Username string `bind:"required"`
	Password string `bind:"required"`
}
