package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	service "github.com/jianyuezhexue/MagicAdmin/service/system"
)

// todo:自定义验证报错信息

// Register 用户注册账号
func Register(c *gin.Context) {
	// 表单验证
	var form system.FormRegister
	err := c.ShouldBind(&form)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := service.Register(form)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, "注册失败", form)
		return
	}

	// 结果返回
	magic.Success(c, http.StatusOK, "注册成功", res)
}

// // FormLogin login structure
// type FormLogin struct {
// 	Username string `json:"userName"` // 用户名
// 	Password string `json:"password"` // 密码
// }

// // Login 用户登录
// func Login(c *gin.Context) {
// 	// 参数校验
// 	var l FormLogin
// 	_ = c.ShouldBindJSON(&l)
// 	if err := request.Verify(l, request.LoginVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	// 逻辑实现
// 	u := &system.SysUser{Username: l.Username, Password: l.Password}
// 	if err, user := userService.Login(u); err != nil {
// 		magic.Logger.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
// 		response.FailWithMessage("用户名不存在或者密码错误", c)
// 	}
// }

// // tokenNext 登录以后签发jwt
// func tokenNext(c *gin.Context, user system.SysUser) {
// 	j := &utils.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)} // 唯一签名
// 	claims := j.CreateClaims(systemReq.BaseClaims{
// 		UUID:        user.UUID,
// 		ID:          user.ID,
// 		NickName:    user.NickName,
// 		Username:    user.Username,
// 		AuthorityId: user.AuthorityID,
// 	})
// 	token, err := j.CreateToken(claims)
// 	if err != nil {
// 		magic.Logger.Error("获取token失败!", zap.Any("err", err))
// 		response.FailWithMessage("获取token失败", c)
// 		return
// 	}
// 	if !magic.Config.System.UseMultipoint {
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
// 		}, "登录成功", c)
// 		return
// 	}

// 	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
// 		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
// 			magic.Logger.Error("设置登录状态失败!", zap.Any("err", err))
// 			response.FailWithMessage("设置登录状态失败", c)
// 			return
// 		}
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
// 		}, "登录成功", c)
// 	} else if err != nil {
// 		magic.Logger.Error("设置登录状态失败!", zap.Any("err", err))
// 		response.FailWithMessage("设置登录状态失败", c)
// 	} else {
// 		var blackJWT system.JwtBlacklist
// 		blackJWT.Jwt = jwtStr
// 		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
// 			response.FailWithMessage("jwt作废失败", c)
// 			return
// 		}
// 		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
// 			response.FailWithMessage("设置登录状态失败", c)
// 			return
// 		}
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
// 		}, "登录成功", c)
// 	}
// }

// // ChangePassword 用户修改密码
// func ChangePassword(c *gin.Context) {
// 	var user systemReq.ChangePasswordStruct
// 	_ = c.ShouldBindJSON(&user)
// 	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	u := &system.SysUser{Username: user.Username, Password: user.Password}
// 	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
// 		magic.Logger.Error("修改失败!", zap.Any("err", err))
// 		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
// 	} else {
// 		response.OkWithMessage("修改成功", c)
// 	}
// }

// // GetUserList 分页获取用户列表 GetUserList
// func GetUserList(c *gin.Context) {
// 	var pageInfo request.PageInfo
// 	_ = c.ShouldBindJSON(&pageInfo)
// 	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
// 		magic.Logger.Error("获取失败!", zap.Any("err", err))
// 		response.FailWithMessage("获取失败", c)
// 	} else {
// 		response.OkWithDetailed(response.PageResult{
// 			List:     list,
// 			Total:    total,
// 			Page:     pageInfo.Page,
// 			PageSize: pageInfo.PageSize,
// 		}, "获取成功", c)
// 	}
// }

// // SetUserAuthority 更改用户权限
// func SetUserAuthority(c *gin.Context) {
// 	var sua systemReq.SetUserAuth
// 	_ = c.ShouldBindJSON(&sua)
// 	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
// 		response.FailWithMessage(UserVerifyErr.Error(), c)
// 		return
// 	}
// 	userID := utils.GetUserID(c)
// 	uuid := utils.GetUserUuid(c)
// 	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
// 		magic.Logger.Error("修改失败!", zap.Any("err", err))
// 		response.FailWithMessage(err.Error(), c)
// 	} else {
// 		claims := utils.GetUserInfo(c)
// 		j := &utils.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)} // 唯一签名
// 		claims.AuthorityId = sua.AuthorityId
// 		if token, err := j.CreateToken(*claims); err != nil {
// 			magic.Logger.Error("修改失败!", zap.Any("err", err))
// 			response.FailWithMessage(err.Error(), c)
// 		} else {
// 			c.Header("new-token", token)
// 			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
// 			response.OkWithMessage("修改成功", c)
// 		}

// 	}
// }

// // SetUserAuthorities 设置用户权限
// func SetUserAuthorities(c *gin.Context) {
// 	var sua systemReq.SetUserAuthorities
// 	_ = c.ShouldBindJSON(&sua)
// 	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
// 		magic.Logger.Error("修改失败!", zap.Any("err", err))
// 		response.FailWithMessage("修改失败", c)
// 	} else {
// 		response.OkWithMessage("修改成功", c)
// 	}
// }

// // DeleteUser 删除用户
// func DeleteUser(c *gin.Context) {
// 	var reqId request.GetById
// 	_ = c.ShouldBindJSON(&reqId)
// 	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	jwtId := utils.GetUserID(c)
// 	if jwtId == uint(reqId.ID) {
// 		response.FailWithMessage("删除失败, 自杀失败", c)
// 		return
// 	}
// 	if err := userService.DeleteUser(reqId.ID); err != nil {
// 		magic.Logger.Error("删除失败!", zap.Any("err", err))
// 		response.FailWithMessage("删除失败", c)
// 	} else {
// 		response.OkWithMessage("删除成功", c)
// 	}
// }

// // SetUserInfo 设置用户信息
// func SetUserInfo(c *gin.Context) {
// 	var user system.SysUser
// 	_ = c.ShouldBindJSON(&user)
// 	if err := utils.Verify(user, utils.IdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err, ReqUser := userService.SetUserInfo(user); err != nil {
// 		magic.Logger.Error("设置失败!", zap.Any("err", err))
// 		response.FailWithMessage("设置失败", c)
// 	} else {
// 		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
// 	}
// }

// // GetUserInfo 获取用户信息
// func GetUserInfo(c *gin.Context) {
// 	uuid := utils.GetUserUuid(c)
// 	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
// 		magic.Logger.Error("获取失败!", zap.Any("err", err))
// 		response.FailWithMessage("获取失败", c)
// 	} else {
// 		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
// 	}
// }

// // UserRouter 系统用户路由
// type UserRouter struct {
// }

// // InitUserRouter 注册用户路由
// func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
// 	baseRouter := Router.Group("user")
// 	baseRouter.POST("login", Login)
// 	return baseRouter
// }
