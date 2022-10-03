## :airplane: 渠道数据中台-管理后台设计

### :tada: 基础建设-数据字典(dictionary)
#### :cake: 需求说明
 - 目录页，数状结构，有父子关系
 - 详情页，中文名，值，数字类型拓展，字符串类型拓展
#### :cake: 数据库设计

```sql
-- 目录表
CREATE TABLE `dictionary` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int DEFAULT NULL COMMENT '父级ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文名',
  `value` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'val值',
  `super` tinyint NOT NULL DEFAULT '0' COMMENT '是否只有超管才有权限',
  `sort` int NOT NULL DEFAULT '50' COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_val` (`value`) USING BTREE COMMENT '字典目录val唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='目录表';

-- 字典详情表
CREATE TABLE `dictionaryItem` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL COMMENT '所属目录ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文名',
  `value` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'val值',
  `super` tinyint NOT NULL DEFAULT '0' COMMENT '是否只有超管才有权限',
  `sort` int NOT NULL DEFAULT '50' COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_pid` (`pid`) USING BTREE COMMENT '查询索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典详情表';
```

#### :cake: 接口列表
 1. 两个表的增删改查
 1. 根据目录sValue查询字典信息

### :tada: 基础建设-权限管理
#### :cake: 需求说明
 - 一级权限类型有:菜单页面权限,业务权限
 - 菜单权限下包含有:页面里的按钮权限，页面里的数据权限
 - 一个人有多个角色

#### :cake: 设计方案
1. 菜单管理中包含了API配置，数据权限配置，表示当前菜单下有哪些接口[页面查询+按钮]，哪些数据字段可以查看；
2. 角色设置权限，每个角色保存:拥有的菜单，拥有的权限：casbin

#### :cake: 方案推导
1. 

#### :cake: 数据库设计

```sql
-- 菜单表
CREATE TABLE `menu2` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parentId` int NOT NULL COMMENT '父菜单ID',
  `path` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由path',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由name',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `hidden` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '普通页面而非菜单',
  `sort` int NOT NULL COMMENT '排序标记',
  `title` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单图标',
  `apis` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '当前菜单下的所有API',
  `fileds` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '当前菜单下的字段权限',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 路由权限表

-- 字段配置权限

-- 角色表

```

#### :cake: 随便写写

- 菜单里有API
- 设置权限会设置菜单和API
- 编辑菜单需要删除API
- 那么角色中的菜单权限，API权限 就不存在了，
- 设置权限最终落库到 casein表
- 删除菜单，角色中的菜单不影响，
- 删除菜单，角色中的API会错位，目前这种设计

- --------------
- 菜单要编辑 ｜ 删除API
- 菜单要删除
- 角色要编辑
- 角色要删除
- 最后如何保证casbin表的数据的正确性；

- --------------

- 

#### :cake: 先确定一个方案出来
 - 菜单表，API表，权限表，casbin表
 - 编辑菜单 -> 先删API表，再重新存
 - 删除菜单或API -> 对应的删除API表 -> 删除casbin中对应的API
 - 设置权限 -> 存菜单ids,存菜单id-权限ID 拼接ID，
 - 设置权限 -> 拿到所有的API信息，存储到casein表
 - 回显用户权限，查菜单ID，权限ID【有被删除的，过滤掉不存在的】
 - 搞定！



https://api.github.com/repos/jianyuezhexue/MagicAdmin/commits?page=0