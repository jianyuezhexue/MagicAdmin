package request

var (
	// LoginVerify 登录验证规则
	LoginVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
