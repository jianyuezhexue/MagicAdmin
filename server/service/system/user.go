package system

import (
	"errors"
	"strconv"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserServer struct{}

var UserApp = new(UserServer)

// Register 注册
func (u *UserServer) Register(data system.User) (res system.User, err error) {
	// 查重
	var user system.User
	err = magic.Orm.Select("id").Where("userName = ?", user.UserName).First(user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return system.User{}, errors.New("用户名已注册")
	}

	// 加密
	data.Password = magic.MD5V(user.Password)
	data.UUID = uuid.NewV4()

	// 创建
	err = magic.Orm.Create(&data).Error
	return user, err
}

// 登录以后签发jwt
func createToken(user system.User, changeToken bool) (token string, err error) {
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
	if !magic.Config.System.UseMultiPoint {
		return token, err
	}

	// 多点登录
	var redisKey string = "token_" + user.UserName

	// 切换角色
	if changeToken {
		_, err = magic.Redis.Del(redisKey)
		return token, err
	}

	// 登录先查
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
	findErr := magic.Orm.Select("*").Where(where, data.UserName).First(&user).Error
	if findErr != nil || user.Password != data.Password {
		return user, errors.New("用户名或密码错误")
	}

	// 查询用户权限
	var auth system.Authority
	err = magic.Orm.Where("id = ?", user.AuthorityId).Find(&auth).Error
	user.Authority = auth

	// 签发JWT
	token, tErr := createToken(user, false)
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
	authIds := make([]string, 0)
	authIds = user.AuthorityIds
	err = magic.Orm.Where("id IN ?", authIds).Order("id").Find(&auths).Error
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
func (u *UserServer) List(data system.SearchUser) (res model.ResPageData, err error) {
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
	err = db.Limit(limit).Offset(offset).Find(&list).Error

	// 组合返回数据
	res = model.ResPageData{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	return res, err
}

// 设置用户角色
func (u *UserServer) SetUserAuth(data system.SetUserAuth) (user system.User, err error) {
	// 验证用户是否存在
	var find system.User
	err = magic.Orm.Where("id = ?", data.Id).Find(&find).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您设置的用户不存在")
	}

	// 设置用户角色IDs
	err = magic.Orm.Model(&user).Where("id = ?", find.Id).Update("authorityIds", data.Ids).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	return user, err
}

// 设置用户状态
func (u *UserServer) SetUserStatus(id model.GetById) (user system.User, err error) {
	// 验证用户是否存在
	var find system.User
	err = magic.Orm.Where("id = ?", id.Id).Find(&find).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您设置的用户不存在")
	}

	// 设置用户角色IDs
	newStatus := 1
	if find.Enable == newStatus {
		newStatus = 2
	}
	err = magic.Orm.Model(&user).Where("id = ?", find.Id).Update("enable", newStatus).Error
	if err != nil {
		return user, errors.New("系统繁忙，请稍后再试")
	}

	return find, err
}

// 删除用户
func (u *UserServer) Delete(id model.GetById) (user system.User, err error) {
	// 查询角色是否存在
	err = magic.Orm.Where("id = ?", id.Id).Find(&user).Error
	if err != nil {
		return user, err
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您删除的用户不存在")
	}

	// 删除处理
	err = magic.Orm.Delete(&user).Error
	return user, err
}

// 更新用户信息
func (u *UserServer) Update(data system.User) (user system.User, err error) {
	// 查询角色是否存在
	err = magic.Orm.Where("id = ?", data.Id).Find(&user).Error
	if err != nil {
		return user, err
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您编辑的用户不存在")
	}

	// 更新处理
	err = magic.Orm.Updates(&data).Error
	return user, err
}

// 重置用户密码|123456
func (u *UserServer) ReSetPwd(id model.GetById) (user system.User, err error) {
	// 查询角色是否存在
	err = magic.Orm.Where("id = ?", id.Id).Find(&user).Error
	if err != nil {
		return user, err
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您编辑的用户不存在")
	}

	// 修改字段
	newPsw := magic.MD5V("123456")
	err = magic.Orm.Model(&user).UpdateColumn("password", newPsw).Error
	return user, err
}

// 切换角色
func (u *UserServer) SwitchAuth(data system.SwitchAuth) (user system.User, err error) {
	// 查询角色是否存在
	err = magic.Orm.Where("uuid = ?", data.UUID).Find(&user).Error
	if err != nil {
		return user, err
	}

	// 用户不存在
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("您编辑的用户不存在")
	}

	// 目标角色和当前角色一致
	if user.AuthorityId == data.AuthorityId {
		return user, errors.New("目标角色当前角色一样")
	}

	// 生成新的token
	user.AuthorityId = data.AuthorityId
	token, err := createToken(user, true)
	if err != nil {
		return user, errors.New("生成签名失败")
	}

	// 更新用户当前角色
	err = magic.Orm.Model(&user).UpdateColumn("authorityId", data.AuthorityId).Error
	if err != nil {
		return user, errors.New("更新当前角色失败")
	}

	// 返回最新的token
	user.Token = token
	return user, err
}
