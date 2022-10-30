package system

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type RecodeServer struct{}

var RecodeApp = new(RecodeServer)

// 保存记录
func (r *RecodeServer) Create(data system.Record) (err error) {
	err = magic.Orm.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}
