/*
 Navicat Premium Data Transfer

 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : dbMagicAdmin

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 25/09/2022 21:37:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `menuId` int NOT NULL COMMENT '所属菜单id',
  `menuName` varchar(20) COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求类型',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api名称',
  `route` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority` (
  `createdAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `deletedAt` datetime DEFAULT NULL,
  `authorityId` int NOT NULL COMMENT '角色ID',
  `authorityName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parentId` int DEFAULT NULL COMMENT '父角色ID',
  `defaultRouter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  `menuIds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单ID字符串拼接',
  PRIMARY KEY (`authorityId`) USING BTREE,
  UNIQUE KEY `authority_id` (`authorityId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of authority
-- ----------------------------
BEGIN;
INSERT INTO `authority` (`createdAt`, `updatedAt`, `deletedAt`, `authorityId`, `authorityName`, `parentId`, `defaultRouter`, `menuIds`) VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 1, '超级管理员', 0, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');
INSERT INTO `authority` (`createdAt`, `updatedAt`, `deletedAt`, `authorityId`, `authorityName`, `parentId`, `defaultRouter`, `menuIds`) VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 2, '普通管理员', 1, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');
INSERT INTO `authority` (`createdAt`, `updatedAt`, `deletedAt`, `authorityId`, `authorityName`, `parentId`, `defaultRouter`, `menuIds`) VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, '测试人员', 0, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');
COMMIT;

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典目录表';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for dictionaryItem
-- ----------------------------
DROP TABLE IF EXISTS `dictionaryItem`;
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

-- ----------------------------
-- Records of dictionaryItem
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parentId` int NOT NULL COMMENT '父菜单ID',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单图标',
  `title` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由path',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由name',
  `hidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是普通页面',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `sort` int NOT NULL COMMENT '排序标记',
  `apis` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'API权限',
  `expandAuth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '拓展权限',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, 0, 'setting', '仪表盘', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (2, 0, 'message', '关于我们', 'about', 'about', 0, 'view/about/index.vue', 7, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (3, 0, 'user', '超级管理员', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (4, 3, 'hot-water', '角色管理', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (5, 3, 'box', '菜单管理', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (6, 3, 'filter', 'API管理', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (7, 3, 'coordinate', '用户管理', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (8, 0, 'basketball', '个人信息', 'person', 'person', 1, 'view/person/person.vue', 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (9, 0, 'chat-square', '示例文件', 'example', 'example', 0, 'view/example/index.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (10, 9, 'chat-line-square', 'excel导入导出', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (11, 9, 'upload', '媒体库（上传下载）', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (12, 9, 'upload', '断点续传', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (13, 9, 'coffee-cup', '客户列表（资源示例）', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (14, 0, 'briefcase', '系统工具', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (15, 14, 'cpu', '代码生成器', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (16, 14, 'magic-stick', '表单生成器', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (17, 14, 'film', '系统配置', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (18, 3, 'files', '字典管理', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (19, 3, 'cherry', '字典详情', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (20, 3, 'alarm-clock', '操作历史', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (22, 0, 'add-location', '官方网站', 'https://www.baidu.com', 'https://www.baidu.com', 0, '/', 0, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (23, 0, 'cloudy', '服务器状态', 'state', 'state', 0, 'view/system/state.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (24, 14, 'coin', '自动化代码管理', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (25, 14, 'calendar', '自动化代码（复用）', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (26, 0, 'apple', 'aaabbbbbb', 'aaabbb', 'aaabbb', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (27, 0, 'aim', 'bbbb', 'bbbb', 'bbbb', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `expandAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (28, 27, 'aim', 'ccc', 'ccc', 'ccc', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for menu2
-- ----------------------------
DROP TABLE IF EXISTS `menu2`;
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

-- ----------------------------
-- Records of menu2
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `userName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录密码',
  `nickName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `headImg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'https://www.pubgmobile.com/cp/pubgpic/images/100001.png' COMMENT '用户头像',
  `authorityId` int NOT NULL DEFAULT '2' COMMENT '用户角色ID',
  `sideMode` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `baseColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `activeColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_userName` (`userName`) USING BTREE,
  KEY `idx_deletedAt` (`deletedAt`) USING BTREE,
  KEY `idx_pasword` (`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `createdAt`, `updatedAt`, `deletedAt`, `uuid`, `userName`, `password`, `nickName`, `headImg`, `authorityId`, `sideMode`, `baseColor`, `activeColor`) VALUES (1, '2022-09-17 04:38:45', NULL, NULL, '3e006a79-a28b-4c6f-a084-99ac9c8ea2e9', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '系统管理员', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi.qqkou.com%2Fi%2F1a3077830965x2286750911b26.jpg&refer=http%3A%2F%2Fi.qqkou.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1665981503&t=c8fefca6768618f0b72fe29f770b4225', 1, 'dark', '#fff', '#1890ff');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
