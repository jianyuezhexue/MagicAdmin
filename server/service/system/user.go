package system

import (
	"errors"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"github.com/jianyuezhexue/MagicAdmin/request"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// UserService 用户注册
type UserService struct {
}

// Register 注册
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(magic.Orm.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = magic.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = magic.Orm.Create(&u).Error
	return u, err
}

// Login 用户登录
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	var user system.SysUser
	u.Password = magic.MD5V([]byte(u.Password))
	err = magic.Orm.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return &user, err
}

// ChangePassword 修改用户密码
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	u.Password = magic.MD5V([]byte(u.Password))
	err = magic.Orm.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", magic.MD5V([]byte(newPassword))).Error
	return u, err
}

// GetUserInfoList 分页获取数据
func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := magic.Orm.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority 设置一个用户的权限
func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityID string) (err error) {
	assignErr := magic.Orm.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityID).First(&system.SysUseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = magic.Orm.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityID).Error
	return err
}

// SetUserAuthorities 设置一个用户的权限
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return magic.Orm.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []system.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

// DeleteUser 删除用户
func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.SysUser
	err = magic.Orm.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = magic.Orm.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

// SetUserInfo 设置用户信息
func (userService *UserService) SetUserInfo(reqUser system.SysUser) (user system.SysUser, err error) {
	err = magic.Orm.Updates(&reqUser).Error
	return reqUser, err
}

// GetUserInfo 获取用户信息
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = magic.Orm.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return reqUser, err
}

// FindUserByID 通过id获取用户信息
func (userService *UserService) FindUserByID(id int) (user *system.SysUser, err error) {
	var u system.SysUser
	err = magic.Orm.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

// FindUserByUuiD 通过uuid获取用户信息
func (userService *UserService) FindUserByUuiD(uuid string) (user *system.SysUser, err error) {
	var u system.SysUser
	if err = magic.Orm.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return &u, nil
}
