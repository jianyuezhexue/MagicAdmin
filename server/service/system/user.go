package system

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserServer struct{}

var UserApp = new(UserServer)

// Register 注册
func (u *UserServer) Register(data system.FormRegister) (res system.User, err error) {

	// 实例化
	user := &system.User{
		UserName:    strings.Trim(data.UserName, " "),
		NickName:    strings.Trim(data.NickName, " "),
		Password:    strings.Trim(data.Password, " "),
		AuthorityId: strings.Trim(data.AuthorityId, " "),
	}

	// 查重
	find := magic.Orm.Select("id").Where("userName = ?", user.UserName).First(user).Error
	if !errors.Is(find, gorm.ErrRecordNotFound) {
		return *user, errors.New("用户名已注册")
	}

	// 加密
	user.Password = magic.MD5V(user.Password)
	user.UUID = uuid.NewV4()

	// 创建
	err = magic.Orm.Create(user).Error
	return *user, err
}

// 登录以后签发jwt
func createToken(user system.User) (token string, err error) {
	// 创建Claims
	jwt := &magic.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)}
	claims := jwt.CreateClaims(magic.BaseClaims{
		UUID:        user.UUID,
		ID:          user.Id,
		NickName:    user.NickName,
		UserName:    user.UserName,
		AuthorityId: user.AuthorityId,
	})
	if token, err = jwt.CreateToken(claims); err != nil {
		return token, err
	}

	// 单点登录-返回
	if !magic.Config.System.UseMultipoint {
		return token, err
	}

	// 多点登录-先查
	var redisKey string = "token_" + user.UserName
	cacheToken, err := magic.Redis.Get(redisKey)
	if err != nil {
		// 多点登录-后存
		err = magic.Redis.Set(redisKey, token, 86400)
		if err != nil {
			return "", err
		}
		return token, err
	}
	return cacheToken, err
}

// Login 用户登录
func (u *UserServer) Login(data system.FormLogin) (user system.User, err error) {
	// 验证登录
	data.Password = magic.MD5V(data.Password)
	where := "userName = ?"
	findErr := magic.Orm.Select("*").Where(where, data.UserName).Preload("Authority").First(&user).Error
	if findErr != nil || user.Password != data.Password {
		return user, errors.New("用户名或密码错误")
	}

	// 签发JWT
	token, tErr := createToken(user)
	if tErr != nil {
		return user, errors.New("生成签名失败")
	}
	user.Token = token

	return user, err
}

// 查询用户信息
func (u *UserServer) UserInfo(uuid uuid.UUID) (user system.User, err error) {
	// 查询用户数据
	err = magic.Orm.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	// 查询角色数据
	var auths []system.Authority
	authIds := strings.Split(user.AuthorityIds, ".")
	err = magic.Orm.Where("id", authIds).Find(&auths).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	// 组合数据
	user.Authorities = auths
	authId, _ := strconv.Atoi(user.AuthorityId)
	for _, item := range auths {
		if item.Id == authId {
			user.Authority = item
			break
		}
	}
	// 返回数据
	return user, err
}

// 用户列表
func (u *UserServer) List(data system.SearchUser) (res magic.PageResult, err error) {

	db := magic.Orm.Model(&system.User{})

	// 查询总数
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return res, err
	}

	// 查询列表数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.User
	err = db.Preload("Authority").Limit(limit).Offset(offset).Find(&list).Error

	// 组合返回数据
	res = magic.PageResult{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	return res, err
}
