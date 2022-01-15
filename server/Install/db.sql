-- 用户表
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `createdAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  `uuid` varchar(64) NOT NULL COMMENT '用户UUID',
  `userName` varchar(40) NOT NULL COMMENT '用户登录名',
  `password` varchar(255) NOT NULL COMMENT '用户登录密码',
  `nickName` varchar(40) NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `headImg` varchar(255) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authorityId` int(20) NOT NULL DEFAULT '2' COMMENT '用户角色ID',
  `sideMode` varchar(5) DEFAULT 'dark' COMMENT '用户侧边主题',
  `baseColor` varchar(10) DEFAULT '#fff' COMMENT '基础颜色',
  `activeColor` varchar(10) DEFAULT '#1890ff' COMMENT '活跃颜色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_userName` (`userName`) USING BTREE,
  KEY `idx_deletedAt` (`deletedAt`),
  KEY `idx_pasword` (`password`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- 菜单表
CREATE TABLE `tbMenus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `createdAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `menuLevel` bigint(20) unsigned DEFAULT NULL,
  `parentId` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `meta` varchar(255) NOT NULL COMMENT 'meta',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  PRIMARY KEY (`id`),
  KEY `idx_deletedAt` (`deletedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;