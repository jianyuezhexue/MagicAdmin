package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Register 注册
func Register(data system.FormRegister) (res system.User, err error) {

	// 实例化
	user := &system.User{
		Username:    strings.Trim(data.Username, " "),
		NickName:    strings.Trim(data.NickName, " "),
		Password:    strings.Trim(data.Password, " "),
		AuthorityID: strings.Trim(data.AuthorityID, " "),
	}

	// 查重
	find := magic.Orm.Select("id").Where("userName = ?", user.Username).First(user).Error
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

// Login 用户登录
func Login(data system.FormLogin) (res system.ResLogin, err error) {
	// 实例化
	user := &system.User{}
	res = system.ResLogin{}

	// 验证登录
	data.Password = magic.MD5V(data.Password)
	where := "userName = ? AND password = ?"
	findErr := magic.Orm.Where(where, data.Username, data.Password).First(user).Error
	if findErr != nil {
		return res, errors.New("用户名或密码错误")
	}
	res.User = *user

	// 签发JWT
	token, tErr := createToken(*user)
	if tErr != nil {
		return res, errors.New("生成签名失败")
	}
	res.Token = token

	return res, err
}

// 登录以后签发jwt
func createToken(user system.User) (token string, err error) {
	// 创建Claims
	jwt := &magic.JWT{SigningKey: []byte(magic.Config.JWT.SigningKey)}
	claims := jwt.CreateClaims(magic.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityID: user.AuthorityID,
	})
	if token, err = jwt.CreateToken(claims); err != nil {
		return token, err
	}

	// 单点登录-返回
	if !magic.Config.System.UseMultipoint {
		return token, err
	}

	// 多点登录-先查
	cacheToken, err := magic.Redis.Get(user.Username)
	if err != nil {
		// 多点登录-后存
		err = magic.Redis.Set(user.Username, token, 86400)
		if err != nil {
			return "", err
		}
		return token, err
	}
	fmt.Println("找到了缓存")
	return cacheToken, err
}

// // ChangePassword 修改用户密码
// func ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
// 	var user system.SysUser
// 	u.Password = magic.MD5V([]byte(u.Password))
// 	err = magic.Orm.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", magic.MD5V([]byte(newPassword))).Error
// 	return u, err
// }

// // GetUserInfoList 分页获取数据
// func GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
// 	limit := info.PageSize
// 	offset := info.PageSize * (info.Page - 1)
// 	db := magic.Orm.Model(&system.SysUser{})
// 	var userList []system.SysUser
// 	err = db.Count(&total).Error
// 	if err != nil {
// 		return
// 	}
// 	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
// 	return userList, total, err
// }

// // SetUserAuthority 设置一个用户的权限
// func SetUserAuthority(id uint, uuid uuid.UUID, authorityID string) (err error) {
// 	assignErr := magic.Orm.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityID).First(&system.SysUseAuthority{}).Error
// 	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
// 		return errors.New("该用户无此角色")
// 	}
// 	err = magic.Orm.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityID).Error
// 	return err
// }

// // SetUserAuthorities 设置一个用户的权限
// func SetUserAuthorities(id uint, authorityIds []string) (err error) {
// 	return magic.Orm.Transaction(func(tx *gorm.DB) error {
// 		TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
// 		if TxErr != nil {
// 			return TxErr
// 		}
// 		useAuthority := []system.SysUseAuthority{}
// 		for _, v := range authorityIds {
// 			useAuthority = append(useAuthority, system.SysUseAuthority{
// 				id, v,
// 			})
// 		}
// 		TxErr = tx.Create(&useAuthority).Error
// 		if TxErr != nil {
// 			return TxErr
// 		}
// 		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
// 		if TxErr != nil {
// 			return TxErr
// 		}
// 		// 返回 nil 提交事务
// 		return nil
// 	})
// }

// // DeleteUser 删除用户
// func DeleteUser(id float64) (err error) {
// 	var user system.SysUser
// 	err = magic.Orm.Where("id = ?", id).Delete(&user).Error
// 	if err != nil {
// 		return err
// 	}
// 	err = magic.Orm.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
// 	return err
// }

// // SetUserInfo 设置用户信息
// func SetUserInfo(reqUser system.SysUser) (user system.SysUser, err error) {
// 	err = magic.Orm.Updates(&reqUser).Error
// 	return reqUser, err
// }

// // GetUserInfo 获取用户信息
// func GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
// 	var reqUser system.SysUser
// 	err = magic.Orm.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
// 	return reqUser, err
// }

// // FindUserByID 通过id获取用户信息
// func FindUserByID(id int) (user *system.SysUser, err error) {
// 	var u system.SysUser
// 	err = magic.Orm.Where("`id` = ?", id).First(&u).Error
// 	return &u, err
// }

// // FindUserByUuiD 通过uuid获取用户信息
// func FindUserByUuiD(uuid string) (user *system.SysUser, err error) {
// 	var u system.SysUser
// 	if err = magic.Orm.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
// 		return errors.New("用户不存在"), &u
// 	}
// 	return &u, nil
// }
