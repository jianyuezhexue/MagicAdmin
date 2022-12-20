package system

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

type DictionaryServer struct{}

var DictionaryApp = new(DictionaryServer)

// 新建目录
func (d *DictionaryServer) Create(data system.Dictionary) (system.Dictionary, error) {
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
func (d *DictionaryServer) List(data system.SearchDictionary) (res model.ResPageData, err error) {
	// 初始化DB
	db := magic.Orm.Model(&system.Dictionary{})

	// 组合搜索条件 | 可以用循环简化代码
	if data.Pid != 0 {
		db = db.Where("`pid` = ?", strconv.Itoa(data.Pid))
	}
	if data.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+data.Name+"%")
	}
	if data.Value != "" {
		db = db.Where("`value` LIKE ?", "%"+data.Value+"%")
	}
	if data.Super != 0 {
		db = db.Where("`super` = ?", data.Super)
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
	res = model.ResPageData{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	// 返回数据
	return res, err
}

// ID查询目录
func (d *DictionaryServer) Item(id model.GetById) (res system.Dictionary, err error) {
	// 查询数据
	err = magic.Orm.Where("id = ?", id.Id).Find(&res).Error
	return res, err
}

// 更新目录
func (d *DictionaryServer) Update(data system.Dictionary) (res system.Dictionary, err error) {
	// 查询数据
	var find system.Dictionary
	err = magic.Orm.Where("id = ?", data.Id).Find(&find).Error
	if err != nil {
		magic.Logger.Info(err.Error())
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 查询条目不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		magic.Logger.Info(err.Error())
		return res, errors.New("编辑的目标不存在")
	}

	// 更新数据
	err = magic.Orm.Updates(data).Error
	return data, err
}

// 删除目录
func (d *DictionaryServer) Delete(id model.GetById) (res system.Dictionary, err error) {
	// 查询数据
	var find system.Dictionary
	err = magic.Orm.Where("id = ?", id.Id).Find(&find).Error
	if err != nil {
		magic.Logger.Info(err.Error())
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 查询条目不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		magic.Logger.Info(err.Error())
		return res, errors.New("删除的目标不存在")
	}

	// 删除数据
	err = magic.Orm.Delete(&find).Error
	return find, err
}

// 根据key查找子目录
func (d *DictionaryServer) DictionaryByKey(key string) magic.BackData {
	// 校验key是否存在
	var find system.Dictionary
	err := magic.Orm.Where("value = ?", key).First(&find).Error

	// 如果不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return magic.Back(21520, "当前key查询不存在，请检查", err.Error())
	}

	// 异常报错
	if err != nil {
		return magic.Back(21522, "系统繁忙，请稍后再试！", err.Error())
	}

	// 查询字典详情
	var dictionaryDetail []system.DictionaryDetail
	err = magic.Orm.Select([]string{"id", "name"}).Where("pid = ?", find.Id).Find(&dictionaryDetail).Error
	if err != nil {
		return magic.Back(21524, "系统繁忙，请稍后再试！", err.Error())
	}

	// 返回结果
	return magic.Back(0, "根据key查找子目录成功", dictionaryDetail)
}

// 根据key查找子目录
func (d *DictionaryServer) DictionaryByKeys(keys string) magic.BackData {

	// 查询keys是否存在
	var Dictionarys []system.Dictionary
	slice := strings.Split(keys, ",")

	err := magic.Orm.Where("value IN ?", slice).Find(&Dictionarys).Error
	if err != nil {
		return magic.Back(21526, "系统繁忙，请稍后再试！", err.Error())
	}

	// 校验key是否正确
	if len(Dictionarys) != len(slice) {
		return magic.Back(21528, "系统繁忙，请稍后再试！", nil)
	}

	// 组合pid
	pids := make([]int, 0)
	for _, item := range Dictionarys {
		pids = append(pids, int(item.Id))
	}

	// 查询字典详情
	var dictionaryDetails []system.DictionaryDetail
	err = magic.Orm.Where("pid IN ?", pids).Find(&dictionaryDetails).Error
	if err != nil {
		return magic.Back(21530, "系统繁忙，请稍后再试！", err.Error())
	}

	// 组合结果
	res := make(map[string][]system.DictionaryDetail, 0)
	for _, value := range dictionaryDetails {
		for _, item := range Dictionarys {
			if value.Pid == int(item.Id) {
				res[item.Value] = append(res[item.Value], value)
			}
		}
	}

	// 返回结果
	return magic.Back(0, "根据key查找子目录成功", res)
}
