## :tada: 项目简介
&emsp;&emsp; Golang+vue 实现的后台管理系统，包含了 JWT登录态，Casbin鉴权，动态菜单管理，人员管理，多角色管理，权限管理[菜单权限，API权限，按钮/数据权限管理]，历史操作记录，定时任务...

## :airplane: 环境安装
 - golang 1.8
 - npm
 - mysql
 - redis

## :cake: 运行调试
```
// 配置文件修改
server/config/config.yaml

// 启动后端
cd server
go run main.go

// 启动前段
ce web
npm run serve

```