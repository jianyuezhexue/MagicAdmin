## :tada: 项目简介
&emsp;&emsp; Golang+vue 实现的后台管理系统，包含了 JWT登录态，Casbin鉴权，动态菜单管理，人员管理，多角色管理，权限管理[菜单权限，API权限，按钮/数据权限管理]，历史操作记录，定时任务等

&emsp;&emsp; 正在新增功能：社群管理系统，文章，评论，点赞，收藏...

页面展示

![登录页](./doc/m1.png)
![仪表盘](./doc/m2.png)
![菜单管理](./doc/m3.png)
![权限管理](./doc/m4.png)
![角色管理](./doc/m5.png)
![字典管理](./doc/m6.png)
![操作历史](./doc/m7.png)
![文章管理](./doc/m8.png)

## :airplane: 环境安装
 - golang 1.8
 - npm
 - mysql
 - redis
 - eleastric search

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

## :cake: 运行调试
