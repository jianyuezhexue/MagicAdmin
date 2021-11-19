package system

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"github.com/jianyuezhexue/MagicAdmin/request"
	"go.uber.org/zap"
)

// User struct
type User struct {
}

// LoginBind login structure
type LoginBind struct {
	Username string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

// Test 基础测试
func Test(c *gin.Context) {
	c.JSON(200, "ok")
}

// // Login 用户登录
// func (u *User) Login(c *gin.Context) {
// 	var l LoginBind
// 	_ = c.ShouldBindJSON(&l)
// 	if err := utils.Verify(l, utils.LoginVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	u := &system.SysUser{Username: l.Username, Password: l.Password}
// 	if err, user := userService.Login(u); err != nil {
// 		magic.Logger.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
// 		response.FailWithMessage("用户名不存在或者密码错误", c)
// 	} else {
// 		u.tokenNext(c, *user)
// 	}
// }

// Login 用户登录
func Login(c *gin.Context) {
	var l LoginBind
	_ = c.ShouldBindJSON(&l)
	if err := request.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.SysUser{Username: l.Username, Password: l.Password}
	if err, user := userService.Login(u); err != nil {
		magic.Logger.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	}
}

// tokenNext 登录以后签发jwt
func (u *User) tokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		magic.Logger.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !magic.Config.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			magic.Logger.Error("设置登录状态失败!", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		magic.Logger.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register 用户注册账号
func (u *User) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities}
	err, userReturn := userService.Register(*user)
	if err != nil {
		magic.Logger.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// ChangePassword 用户修改密码
func (u *User) ChangePassword(c *gin.Context) {
	var user systemReq.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.SysUser{Username: user.Username, Password: user.Password}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		magic.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// 分页获取用户列表 GetUserList
func (u *User) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userService.GetUserInfoList(pageInfo); err != nil {
		magic.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// SetUserAuthority 更改用户权限
func (u *User) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	uuid := utils.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, sua.AuthorityId); err != nil {
		magic.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		claims := utils.GetUserInfo(c)
		j := &utils.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)} // 唯一签名
		claims.AuthorityId = sua.AuthorityId
		if token, err := j.CreateToken(*claims); err != nil {
			magic.Logger.Error("修改失败!", zap.Any("err", err))
			response.FailWithMessage(err.Error(), c)
		} else {
			c.Header("new-token", token)
			c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt, 10))
			response.OkWithMessage("修改成功", c)
		}

	}
}

// SetUserAuthorities 设置用户权限
func (u *User) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	_ = c.ShouldBindJSON(&sua)
	if err := userService.SetUserAuthorities(sua.ID, sua.AuthorityIds); err != nil {
		magic.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// DeleteUser 删除用户
func (u *User) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		magic.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// SetUserInfo 设置用户信息
func (u *User) SetUserInfo(c *gin.Context) {
	var user system.SysUser
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, ReqUser := userService.SetUserInfo(user); err != nil {
		magic.Logger.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "设置成功", c)
	}
}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	if err, ReqUser := userService.GetUserInfo(uuid); err != nil {
		magic.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
	}
}

// UserRouter 系统用户路由
type UserRouter struct {
}

// InitUserRouter 注册用户路由
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("user")
	baseRouter.POST("login", Login)
	return baseRouter
}
