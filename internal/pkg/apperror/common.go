package apperror

var (
	ErrInvalidParam = New(InvalidParam, "参数错误")
	ErrUnauthorized = New(Unauthorized, "未登录")
	ErrForbidden    = New(Forbidden, "无权限")
	ErrInternal     = New(InternalError, "系统内部错误")
)
