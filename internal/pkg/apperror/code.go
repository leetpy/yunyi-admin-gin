package apperror

const (
	// 通用
	Success       int = 0
	InvalidParam  int = 10001
	Unauthorized  int = 10002
	Forbidden     int = 10003
	NotFound      int = 10004
	InternalError int = 10005

	// 用户模块 2xxxx
	UserNotFound  int = 20001
	UserExists    int = 20002
	PasswordWrong int = 20003
)
