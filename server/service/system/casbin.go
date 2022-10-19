package system

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"go.uber.org/zap"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

type CasbinServer struct{}

var CasbinApp = new(CasbinServer)

func (CasbinServer *CasbinServer) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDBUseTableName(magic.Orm, "", "casbin")
		if err != nil {
			magic.Print(err)
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
		`
		model, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(model, adapter)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
