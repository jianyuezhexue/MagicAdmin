package system

import (
	"errors"
	"strconv"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type DictionaryServer struct{}

var DictionaryApp = new(DictionaryServer)

// 新建目录
func (d DictionaryServer) Create(data system.Dictionary) (system.Dictionary, error) {
	// 查重Value值
	find := system.Dictionary{}
	err := magic.Orm.Where("value = ?", data.Value).Find(&find).Error
	if err != nil {
		return find, errors.New("数据库跪了")
	}

	// val值重复提醒
	if data.Value == find.Value {
		return find, errors.New("英文字典名不能重复，请修改")
	}

	// 保存数据
	err = magic.Orm.Create(&data).Error
	return data, err
}

// 分页查询
func (d DictionaryServer) List(data system.SearchInfo) (res magic.PageResult, err error) {
	// 初始化DB
	db := magic.Orm.Model(&system.Dictionary{})

	// 组合搜索条件 | 可以用循环简化代码
	if data.Pid != 0 {
		db = db.Where("`name` LIKE ?", "%"+strconv.Itoa(data.Pid)+"%")
	}
	if data.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+data.Name+"%")
	}
	if data.Value != "" {
		db = db.Where("`name` LIKE ?", "%"+data.Value+"%")
	}
	if data.Super != 0 {
		db = db.Where("`name` LIKE ?", "%"+strconv.Itoa(data.Super)+"%")
	}
	if data.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+data.Desc+"%")
	}

	// 查询总数
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return res, err
	}

	// 查询列表数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.Dictionary
	err = db.Limit(limit).Offset(offset).Find(&list).Error

	// 组合返回数据
	res = magic.PageResult{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	// 返回数据
	return res, err
}
