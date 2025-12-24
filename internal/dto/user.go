package dto

type UserLoginReq struct {
	Username string
	Password string
}

type UserListReq struct{}

type CreateUserReq struct {
	Username string `bind:"required"`
	Password string `bind:"required"`
}
