-- 用户表
CREATE TABLE `tbUsers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `createdAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `userName` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nickName` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `headerImg` varchar(191) DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '用户头像',
  `authorityId` varchar(90) DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_userName` (`userName`) USING BTREE,
  KEY `idx_deletedAt` (`deletedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;

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