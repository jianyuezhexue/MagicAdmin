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

 Date: 05/11/2022 12:53:46
*/

create database `dbMagicAdmin` default character set utf8mb4 collate utf8mb4_general_ci;

use dbMagicAdmin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `menuId` int NOT NULL COMMENT '所属菜单id',
  `menuName` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求类型',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api名称',
  `route` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由路径',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_type_route` (`method`,`route`) USING BTREE COMMENT '路由不能重复'
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (62, 19, '', 'GET', '获取字典列表', '/dictionaryDetail', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (63, 19, '', 'POST', '新增字典详情', '/dictionaryDetail', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (65, 19, '', 'PUT', '编辑字典详情', '/dictionaryDetail', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (68, 19, '', 'DELETE', '删除字典详情', '/dictionaryDetail/', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (69, 19, '', 'GET', '查询字典详情', '/dictionaryDetail/', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (73, 18, '', 'GET', '字典目录列表', '/dictionary', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (74, 18, '', 'POST', '新增字典目录', '/dictionary', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (75, 18, '', 'PUT', '编辑字典目录', '/dictionary', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (76, 18, '', 'GET', '查询字典目录', '/dictionary/', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (77, 18, '', 'DELETE', '删除字典目录', '/dictionary/', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (78, 5, '', 'GET', '菜单列表', '/menus', '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (79, 5, '', 'POST', '新建菜单', '/menu', '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (80, 5, '', 'PUT', '编辑菜单', '/menu', '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (81, 5, '', 'GET', '查询菜单', '/menu/', '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
INSERT INTO `api` (`id`, `menuId`, `menuName`, `method`, `name`, `route`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (82, 5, '', 'DELETE', '删除菜单', '/menu/', '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
COMMIT;

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名',
  `pid` int DEFAULT NULL COMMENT '父角色ID',
  `defaultRouter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  `menuIds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单ID字符串拼接',
  `apiIds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'API权限ID字符串',
  `extAuth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '其他权限',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色描述',
  `createdAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `deletedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of authority
-- ----------------------------
BEGIN;
INSERT INTO `authority` (`id`, `name`, `pid`, `defaultRouter`, `menuIds`, `apiIds`, `extAuth`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, '超级管理员', 0, 'dashboard', '1,2,3,4,5,6,7,11,16,18,19,20,8', '78,79,80,81,82,73,74,75,76,77,62,63,65,68,69', NULL, '拥有全部权限', '2021-11-22 17:07:02', '2022-10-30 09:29:07', NULL);
INSERT INTO `authority` (`id`, `name`, `pid`, `defaultRouter`, `menuIds`, `apiIds`, `extAuth`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (2, '开发人员', 1, 'dashboard', '1,2,3,4,5,6,7,11,16,18,19,20,8', '78,79,80,81,82,73,74,75,76,77,62,63,65,68,69', NULL, '开发人员权限', '2021-11-22 17:07:02', '2022-10-23 05:12:32', NULL);
INSERT INTO `authority` (`id`, `name`, `pid`, `defaultRouter`, `menuIds`, `apiIds`, `extAuth`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (3, '测试人员', 1, 'dashboard', '1,2,3,4,7,20,8', NULL, NULL, '测试人员权限', '2021-11-22 17:07:02', '2022-10-15 05:19:38', NULL);
INSERT INTO `authority` (`id`, `name`, `pid`, `defaultRouter`, `menuIds`, `apiIds`, `extAuth`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (6, '产品同学', 0, '', '1,2,3,4,5,6,7,11,16,18,19,20,8', NULL, NULL, '产品同学权限', '2022-10-01 13:35:31', '2022-10-15 05:19:48', NULL);
COMMIT;

-- ----------------------------
-- Table structure for casbin
-- ----------------------------
DROP TABLE IF EXISTS `casbin`;
CREATE TABLE `casbin` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin
-- ----------------------------
BEGIN;
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (85, 'p', '1', 'DELETE', '/dictionary/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (79, 'p', '1', 'DELETE', '/dictionaryDetail/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (90, 'p', '1', 'DELETE', '/menu/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (81, 'p', '1', 'GET', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (84, 'p', '1', 'GET', '/dictionary/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (76, 'p', '1', 'GET', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (80, 'p', '1', 'GET', '/dictionaryDetail/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (89, 'p', '1', 'GET', '/menu/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (86, 'p', '1', 'GET', '/menus', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (82, 'p', '1', 'POST', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (77, 'p', '1', 'POST', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (87, 'p', '1', 'POST', '/menu', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (83, 'p', '1', 'PUT', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (78, 'p', '1', 'PUT', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (88, 'p', '1', 'PUT', '/menu', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (25, 'p', '2', 'DELETE', '/dictionary/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (19, 'p', '2', 'DELETE', '/dictionaryDetail/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (30, 'p', '2', 'DELETE', '/menu/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (21, 'p', '2', 'GET', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (24, 'p', '2', 'GET', '/dictionary/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (16, 'p', '2', 'GET', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (20, 'p', '2', 'GET', '/dictionaryDetail/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (29, 'p', '2', 'GET', '/menu/', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (26, 'p', '2', 'GET', '/menus', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (22, 'p', '2', 'POST', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (17, 'p', '2', 'POST', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (27, 'p', '2', 'POST', '/menu', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (23, 'p', '2', 'PUT', '/dictionary', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (18, 'p', '2', 'PUT', '/dictionaryDetail', '', '', '');
INSERT INTO `casbin` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (28, 'p', '2', 'PUT', '/menu', '', '', '');
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
  `super` tinyint NOT NULL DEFAULT '0' COMMENT '超管才有权限 1-有 2-无',
  `sort` int NOT NULL DEFAULT '50' COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_val` (`value`) USING BTREE COMMENT '字典目录val唯一'
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='字典目录表';

-- ----------------------------
-- Records of dictionary
-- ----------------------------
BEGIN;
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (9, 0, '菜单权限类型', 'powerType', 1, 50, '菜单权限类型', '2022-10-01 05:40:34', '2022-10-01 05:40:35', NULL);
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (16, 0, '人员身份类型', 'idType', 1, 50, '人员身份类型', '2022-10-01 05:40:15', '2022-10-01 05:40:16', NULL);
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (19, 0, 'bbb', 'bbb', 2, 50, 'bbbb', '2022-10-01 05:17:31', '2022-09-30 08:05:26', '2022-10-01 05:17:31');
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (20, 0, 'asd', 'asdf', 2, 50, 'adsf', '2022-10-01 04:11:29', '2022-09-30 08:22:05', '2022-10-01 04:11:29');
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (21, 0, '111112', '111112', 1, 50, '111112', '2022-09-30 16:40:29', '2022-09-30 08:12:52', '2022-09-30 08:40:29');
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (22, 0, 'ddd', 'ddd', 1, 50, 'ddd', '2022-10-01 05:17:36', '2022-10-01 04:11:39', '2022-10-01 05:17:37');
INSERT INTO `dictionary` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (23, 0, 'a', 'a', 1, 50, 'a', '2022-10-29 10:54:55', '2022-10-29 10:54:06', '2022-10-29 10:54:56');
COMMIT;

-- ----------------------------
-- Table structure for dictionaryDetail
-- ----------------------------
DROP TABLE IF EXISTS `dictionaryDetail`;
CREATE TABLE `dictionaryDetail` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `pid` int DEFAULT NULL COMMENT '所属目录ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文名',
  `value` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'val值',
  `super` tinyint NOT NULL DEFAULT '0' COMMENT '是否只有超管才有权限',
  `sort` int NOT NULL DEFAULT '50' COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_pid` (`pid`) USING BTREE COMMENT '查询索引'
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='字典详情表';

-- ----------------------------
-- Records of dictionaryDetail
-- ----------------------------
BEGIN;
INSERT INTO `dictionaryDetail` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (7, 9, '数据权限', 'dataPower', 1, 50, '用户是否拥有某一个字段的权限', '2022-10-01 05:26:33', '2022-10-01 05:26:33', NULL);
INSERT INTO `dictionaryDetail` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (8, 9, '按钮权限', 'btnPower', 1, 50, '用户在这个页面上按钮是否显示', '2022-10-01 09:51:49', '2022-10-01 09:51:49', NULL);
INSERT INTO `dictionaryDetail` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (9, 16, '外发PM', 'PM', 1, 50, '外发PM', '2022-10-01 05:18:53', '2022-10-01 05:18:53', NULL);
INSERT INTO `dictionaryDetail` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (10, 16, '跟进产品', 'pm', 1, 50, '跟进产品', '2022-10-01 09:53:21', '2022-10-01 09:53:21', NULL);
INSERT INTO `dictionaryDetail` (`id`, `pid`, `name`, `value`, `super`, `sort`, `desc`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (11, 9, 'a', 'a', 2, 50, 'a', '2022-10-29 10:25:23', '2022-10-29 10:25:02', '2022-10-29 10:25:23');
COMMIT;

-- ----------------------------
-- Table structure for extAuth
-- ----------------------------
DROP TABLE IF EXISTS `extAuth`;
CREATE TABLE `extAuth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `menuId` int NOT NULL,
  `menuName` varchar(40) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `type` tinyint(1) NOT NULL COMMENT '权限类型:按钮权限，数据权限',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限中文名',
  `val` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '字段值',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of extAuth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mediaFiles
-- ----------------------------
DROP TABLE IF EXISTS `mediaFiles`;
CREATE TABLE `mediaFiles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件名',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件地址',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件标签',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件编号',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of mediaFiles
-- ----------------------------
BEGIN;
INSERT INTO `mediaFiles` (`id`, `name`, `url`, `tag`, `key`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, '1', 'https://i.postimg.cc/Wb026mfD/logo.jpg', '1', '1', '2022-10-23 14:02:24', NULL, NULL);
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
  `hidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT '菜单不显示，且tab栏失焦自动关闭',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `sort` int NOT NULL COMMENT '排序标记',
  `apis` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'API权限',
  `extAuth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '拓展权限',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, 0, 'home-filled', '仪表盘', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '[]', '[]', '2022-10-01 05:45:00', '2022-10-01 05:45:00', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (2, 0, 'message', '关于我们', 'about', 'about', 0, 'view/about/index.vue', 4, NULL, NULL, '2022-10-15 13:37:41', '2022-10-15 13:37:42', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (3, 0, 'user', '系统管理', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 2, '', NULL, '2022-10-15 05:13:17', '2022-10-15 05:13:18', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (4, 3, 'hot-water', '角色管理', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 4, NULL, '[{\"type\":\"type\",\"name\":\"设置权限\",\"val\":\"setAuth\"}]', '2022-10-15 05:04:25', '2022-10-15 05:04:26', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (5, 3, 'menu', '菜单管理', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 1, NULL, NULL, '2022-10-15 05:03:48', '2022-10-15 05:03:48', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (6, 3, 'filter', 'API管理', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 2, NULL, NULL, '2022-10-15 05:05:07', '2022-10-15 05:05:07', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (7, 3, 'coordinate', '用户管理', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 5, NULL, NULL, '2022-10-15 05:04:40', '2022-10-15 05:04:41', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (8, 0, 'basketball', '个人信息', 'person', 'person', 1, 'view/person/person.vue', 5, NULL, NULL, '2022-10-15 13:37:47', '2022-10-15 13:37:47', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (9, 0, 'chat-square', '示例文件', 'example', 'example', 0, 'view/example/index.vue', 6, NULL, NULL, '2022-10-15 04:57:06', NULL, '2022-10-15 04:57:06');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (10, 9, 'chat-line-square', 'excel导入导出', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, NULL, NULL, '2022-10-15 04:56:56', NULL, '2022-10-15 04:56:56');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (11, 3, 'upload', '媒体库', 'upload', 'upload', 0, 'view/superAdmin/upload/upload.vue', 7, NULL, NULL, '2022-10-15 05:05:49', '2022-10-15 05:05:50', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (12, 9, 'upload', '断点续传', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, NULL, NULL, '2022-10-15 04:56:59', NULL, '2022-10-15 04:56:59');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (13, 9, 'coffee-cup', '客户列表（资源示例）', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, NULL, NULL, '2022-10-15 04:57:02', NULL, '2022-10-15 04:57:02');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (14, 0, 'briefcase', '系统工具', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, NULL, NULL, '2022-10-15 04:58:16', NULL, '2022-10-15 04:58:17');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (15, 14, 'cpu', '代码生成器', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, NULL, NULL, '2022-10-15 04:58:08', NULL, '2022-10-15 04:58:09');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (16, 3, 'magic-stick', '表单生成器', 'formCreate', 'formCreate', 0, 'view/superAdmin/formCreate/index.vue', 6, NULL, NULL, '2022-10-15 05:05:40', '2022-10-15 05:05:41', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (17, 14, 'film', '系统配置', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, NULL, NULL, '2022-10-15 04:58:05', NULL, '2022-10-15 04:58:06');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (18, 3, 'files', '字典管理', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/dictionary.vue', 3, '[{\"type\":\"GET\",\"name\":\"查询字典目录列表\",\"route\":\"aaa\"},{\"type\":\"POST\",\"name\":\"新增字典目录\",\"route\":\"aaa\"},{\"type\":\"PUT\",\"name\":\"修改字典目录\",\"route\":\"aaa\"},{\"type\":\"DELETE\",\"name\":\"删除字典目录\",\"route\":\"aaa\"}]', '[]', '2022-10-15 05:05:19', '2022-10-15 05:05:20', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (19, 3, 'cherry', '字典详情', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/dictionaryDetail.vue', 3, '[{\"type\":\"GET\",\"name\":\"查询字典详情列表\",\"route\":\"aaa\"},{\"type\":\"POST\",\"name\":\"新增字典详情\",\"route\":\"aaa\"},{\"type\":\"PUT\",\"name\":\"修改字典详情\",\"route\":\"aaa\"},{\"type\":\"DELETE\",\"name\":\"删除字典详情\",\"route\":\"aaa\"}]', '[]', '2022-10-15 05:04:17', '2022-10-15 05:04:18', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (20, 3, 'alarm-clock', '操作历史', 'operation', 'operation', 0, 'view/superAdmin/operation/record.vue', 8, NULL, NULL, '2022-10-30 02:40:38', '2022-10-30 02:40:38', NULL);
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (22, 0, 'compass', '官方网站', 'https://www.baidu.com', 'https://www.baidu.com', 0, '/', 0, '[]', '[]', '2022-10-15 04:59:24', '2022-10-01 05:46:02', '2022-10-15 04:59:25');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (23, 0, 'cloudy', '服务器状态', 'state', 'state', 0, 'view/system/state.vue', 6, NULL, NULL, '2022-10-15 04:58:36', NULL, '2022-10-15 04:58:37');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (24, 14, 'coin', '自动化代码管理', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, NULL, NULL, '2022-10-15 04:58:11', NULL, '2022-10-15 04:58:11');
INSERT INTO `menu` (`id`, `parentId`, `icon`, `title`, `path`, `name`, `hidden`, `component`, `sort`, `apis`, `extAuth`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (25, 14, 'calendar', '自动化代码（复用）', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, NULL, NULL, '2022-10-15 04:58:13', NULL, '2022-10-15 04:58:14');
COMMIT;

-- ----------------------------
-- Table structure for record
-- ----------------------------
DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `userId` int DEFAULT NULL COMMENT '用户ID',
  `userName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名',
  `ip` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'IP地址',
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` int DEFAULT NULL COMMENT '状态码',
  `costTime` float(10,5) DEFAULT NULL COMMENT '请求耗时',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '返回消息',
  `params` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求传参',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '接口返回',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=244 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of record
-- ----------------------------
BEGIN;
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (230, 1, 'admin', '127.0.0.1', 'GET', '/users', 200, 25.60691, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询用户列表成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '{\"code\":0,\"msg\":\"分页查询用户列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-29 10:53:41\",\"uuid\":\"3e006a79-a28b-4c6f-a084-99ac9c8ea2e9\",\"userName\":\"admin\",\"nickName\":\"超级管理员\",\"phone\":\"15848606970\",\"email\":\"490513975@qq.com\",\"headImg\":\"https://i.postimg.cc/Wb026mfD/logo.jpg\",\"authorityId\":\"1\",\"authorityIds\":[\"1\",\"3\"],\"enable\":1,\"sideMode\":\"dark\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"token\":\"\",\"authority\":{\"createdAt\":\"0001-01-01 00:00:00\",\"id\":0,\"name\":\"\",\"desc\":\"\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":null},\"authorities\":null}],\"total\":1,\"page\":1,\"pageSize\":10},\"costTime\":25.56493}', '2022-10-30 09:57:27', '2022-10-30 09:57:27', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (231, 1, 'admin', '127.0.0.1', 'GET', '/authorityTree', 200, 44.69606, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取角色列表成功', '{}', '{\"code\":0,\"msg\":\"获取角色列表成功\",\"data\":{\"list\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":1,\"name\":\"超级管理员\",\"desc\":\"拥有全部权限\",\"pid\":0,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":2,\"name\":\"开发人员\",\"desc\":\"开发人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[]},{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":3,\"name\":\"测试人员\",\"desc\":\"测试人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,7,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}]},{\"createdAt\":\"2022-10-01 13:35:31\",\"id\":6,\"name\":\"产品同学\",\"desc\":\"产品同学权限\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}],\"total\":999,\"page\":1,\"pageSize\":10},\"costTime\":44.65101}', '2022-10-30 09:57:30', '2022-10-30 09:57:30', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (232, 1, 'admin', '127.0.0.1', 'GET', '/authorityTree', 200, 6.78298, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取角色列表成功', '{}', '{\"code\":0,\"msg\":\"获取角色列表成功\",\"data\":{\"list\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":1,\"name\":\"超级管理员\",\"desc\":\"拥有全部权限\",\"pid\":0,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":2,\"name\":\"开发人员\",\"desc\":\"开发人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[]},{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":3,\"name\":\"测试人员\",\"desc\":\"测试人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,7,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}]},{\"createdAt\":\"2022-10-01 13:35:31\",\"id\":6,\"name\":\"产品同学\",\"desc\":\"产品同学权限\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}],\"total\":999,\"page\":1,\"pageSize\":10},\"costTime\":6.29402}', '2022-10-30 09:57:33', '2022-10-30 09:57:33', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (233, 1, 'admin', '127.0.0.1', 'GET', '/recode', 200, 8.82586, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询操作记录成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '不记录', '2022-10-30 09:57:36', '2022-10-30 09:57:36', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (234, 1, 'admin', '127.0.0.1', 'GET', '/authorityTree', 200, 5.29485, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取角色列表成功', '{}', '{\"code\":0,\"msg\":\"获取角色列表成功\",\"data\":{\"list\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":1,\"name\":\"超级管理员\",\"desc\":\"拥有全部权限\",\"pid\":0,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":2,\"name\":\"开发人员\",\"desc\":\"开发人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[]},{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":3,\"name\":\"测试人员\",\"desc\":\"测试人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,7,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}]},{\"createdAt\":\"2022-10-01 13:35:31\",\"id\":6,\"name\":\"产品同学\",\"desc\":\"产品同学权限\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}],\"total\":999,\"page\":1,\"pageSize\":10},\"costTime\":5.248}', '2022-10-30 09:57:39', '2022-10-30 09:57:39', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (235, 1, 'admin', '127.0.0.1', 'GET', '/users', 200, 9.56006, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询用户列表成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '{\"code\":0,\"msg\":\"分页查询用户列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-29 10:53:41\",\"uuid\":\"3e006a79-a28b-4c6f-a084-99ac9c8ea2e9\",\"userName\":\"admin\",\"nickName\":\"超级管理员\",\"phone\":\"15848606970\",\"email\":\"490513975@qq.com\",\"headImg\":\"https://i.postimg.cc/Wb026mfD/logo.jpg\",\"authorityId\":\"1\",\"authorityIds\":[\"1\",\"3\"],\"enable\":1,\"sideMode\":\"dark\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"token\":\"\",\"authority\":{\"createdAt\":\"0001-01-01 00:00:00\",\"id\":0,\"name\":\"\",\"desc\":\"\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":null},\"authorities\":null}],\"total\":1,\"page\":1,\"pageSize\":10},\"costTime\":9.5191}', '2022-10-30 09:57:42', '2022-10-30 09:57:42', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (236, 1, 'admin', '127.0.0.1', 'GET', '/menus', 200, 3.37997, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取菜单列表成功', '{}', '{\"code\":0,\"msg\":\"获取菜单列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-01 05:45:00\",\"meta\":{\"title\":\"仪表盘\",\"icon\":\"home-filled\"},\"parentId\":0,\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":3,\"createdAt\":\"2022-10-15 05:13:17\",\"meta\":{\"title\":\"系统管理\",\"icon\":\"user\"},\"parentId\":0,\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":2,\"api\":null,\"extAuth\":null,\"children\":[{\"id\":5,\"createdAt\":\"2022-10-15 05:03:48\",\"meta\":{\"title\":\"菜单管理\",\"icon\":\"menu\"},\"parentId\":3,\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":1,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":6,\"createdAt\":\"2022-10-15 05:05:07\",\"meta\":{\"title\":\"API管理\",\"icon\":\"filter\"},\"parentId\":3,\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":2,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":18,\"createdAt\":\"2022-10-15 05:05:19\",\"meta\":{\"title\":\"字典管理\",\"icon\":\"files\"},\"parentId\":3,\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/dictionary.vue\",\"sort\":3,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":19,\"createdAt\":\"2022-10-15 05:04:17\",\"meta\":{\"title\":\"字典详情\",\"icon\":\"cherry\"},\"parentId\":3,\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/dictionaryDetail.vue\",\"sort\":3,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":4,\"createdAt\":\"2022-10-15 05:04:25\",\"meta\":{\"title\":\"角色管理\",\"icon\":\"hot-water\"},\"parentId\":3,\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":4,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":7,\"createdAt\":\"2022-10-15 05:04:40\",\"meta\":{\"title\":\"用户管理\",\"icon\":\"coordinate\"},\"parentId\":3,\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":5,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":16,\"createdAt\":\"2022-10-15 05:05:40\",\"meta\":{\"title\":\"表单生成器\",\"icon\":\"magic-stick\"},\"parentId\":3,\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/superAdmin/formCreate/index.vue\",\"sort\":6,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":11,\"createdAt\":\"2022-10-15 05:05:49\",\"meta\":{\"title\":\"媒体库\",\"icon\":\"upload\"},\"parentId\":3,\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/superAdmin/upload/upload.vue\",\"sort\":7,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":20,\"createdAt\":\"2022-10-30 02:40:38\",\"meta\":{\"title\":\"操作历史\",\"icon\":\"alarm-clock\"},\"parentId\":3,\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/record.vue\",\"sort\":8,\"api\":null,\"extAuth\":null,\"children\":[]}]},{\"id\":2,\"createdAt\":\"2022-10-15 13:37:41\",\"meta\":{\"title\":\"关于我们\",\"icon\":\"message\"},\"parentId\":0,\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":4,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":8,\"createdAt\":\"2022-10-15 13:37:47\",\"meta\":{\"title\":\"个人信息\",\"icon\":\"basketball\"},\"parentId\":0,\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":5,\"api\":null,\"extAuth\":null,\"children\":[]}],\"total\":13,\"page\":1,\"pageSize\":999},\"costTime\":3.28704}', '2022-10-30 09:57:45', '2022-10-30 09:57:45', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (237, 1, 'admin', '127.0.0.1', 'GET', '/menus', 200, 8.90598, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取菜单列表成功', '{}', '{\"code\":0,\"msg\":\"获取菜单列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-01 05:45:00\",\"meta\":{\"title\":\"仪表盘\",\"icon\":\"home-filled\"},\"parentId\":0,\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":3,\"createdAt\":\"2022-10-15 05:13:17\",\"meta\":{\"title\":\"系统管理\",\"icon\":\"user\"},\"parentId\":0,\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":2,\"api\":null,\"extAuth\":null,\"children\":[{\"id\":5,\"createdAt\":\"2022-10-15 05:03:48\",\"meta\":{\"title\":\"菜单管理\",\"icon\":\"menu\"},\"parentId\":3,\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":1,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":6,\"createdAt\":\"2022-10-15 05:05:07\",\"meta\":{\"title\":\"API管理\",\"icon\":\"filter\"},\"parentId\":3,\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":2,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":18,\"createdAt\":\"2022-10-15 05:05:19\",\"meta\":{\"title\":\"字典管理\",\"icon\":\"files\"},\"parentId\":3,\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/dictionary.vue\",\"sort\":3,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":19,\"createdAt\":\"2022-10-15 05:04:17\",\"meta\":{\"title\":\"字典详情\",\"icon\":\"cherry\"},\"parentId\":3,\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/dictionaryDetail.vue\",\"sort\":3,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":4,\"createdAt\":\"2022-10-15 05:04:25\",\"meta\":{\"title\":\"角色管理\",\"icon\":\"hot-water\"},\"parentId\":3,\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":4,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":7,\"createdAt\":\"2022-10-15 05:04:40\",\"meta\":{\"title\":\"用户管理\",\"icon\":\"coordinate\"},\"parentId\":3,\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":5,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":16,\"createdAt\":\"2022-10-15 05:05:40\",\"meta\":{\"title\":\"表单生成器\",\"icon\":\"magic-stick\"},\"parentId\":3,\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/superAdmin/formCreate/index.vue\",\"sort\":6,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":11,\"createdAt\":\"2022-10-15 05:05:49\",\"meta\":{\"title\":\"媒体库\",\"icon\":\"upload\"},\"parentId\":3,\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/superAdmin/upload/upload.vue\",\"sort\":7,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":20,\"createdAt\":\"2022-10-30 02:40:38\",\"meta\":{\"title\":\"操作历史\",\"icon\":\"alarm-clock\"},\"parentId\":3,\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/record.vue\",\"sort\":8,\"api\":null,\"extAuth\":null,\"children\":[]}]},{\"id\":2,\"createdAt\":\"2022-10-15 13:37:41\",\"meta\":{\"title\":\"关于我们\",\"icon\":\"message\"},\"parentId\":0,\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":4,\"api\":null,\"extAuth\":null,\"children\":[]},{\"id\":8,\"createdAt\":\"2022-10-15 13:37:47\",\"meta\":{\"title\":\"个人信息\",\"icon\":\"basketball\"},\"parentId\":0,\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":5,\"api\":null,\"extAuth\":null,\"children\":[]}],\"total\":13,\"page\":1,\"pageSize\":999},\"costTime\":8.81101}', '2022-10-30 09:57:48', '2022-10-30 09:57:48', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (238, 1, 'admin', '127.0.0.1', 'GET', '/dictionary', 200, 9.90003, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询字典目录成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '{\"code\":0,\"msg\":\"分页查询字典目录成功\",\"data\":{\"list\":[{\"id\":9,\"createdAt\":\"2022-10-01 05:40:34\",\"pid\":0,\"name\":\"菜单权限类型\",\"value\":\"powerType\",\"super\":1,\"sort\":50,\"desc\":\"菜单权限类型\"},{\"id\":16,\"createdAt\":\"2022-10-01 05:40:15\",\"pid\":0,\"name\":\"人员身份类型\",\"value\":\"idType\",\"super\":1,\"sort\":50,\"desc\":\"人员身份类型\"}],\"total\":2,\"page\":1,\"pageSize\":10},\"costTime\":9.80787}', '2022-10-30 09:57:51', '2022-10-30 09:57:51', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (239, 1, 'admin', '127.0.0.1', 'GET', '/authorityTree', 200, 7.69792, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取角色列表成功', '{}', '{\"code\":0,\"msg\":\"获取角色列表成功\",\"data\":{\"list\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":1,\"name\":\"超级管理员\",\"desc\":\"拥有全部权限\",\"pid\":0,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":2,\"name\":\"开发人员\",\"desc\":\"开发人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[]},{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":3,\"name\":\"测试人员\",\"desc\":\"测试人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,7,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}]},{\"createdAt\":\"2022-10-01 13:35:31\",\"id\":6,\"name\":\"产品同学\",\"desc\":\"产品同学权限\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}],\"total\":999,\"page\":1,\"pageSize\":10},\"costTime\":7.648}', '2022-10-30 09:57:54', '2022-10-30 09:57:54', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (240, 1, 'admin', '127.0.0.1', 'GET', '/authorityTree', 200, 8.73114, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '获取角色列表成功', '{}', '{\"code\":0,\"msg\":\"获取角色列表成功\",\"data\":{\"list\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":1,\"name\":\"超级管理员\",\"desc\":\"拥有全部权限\",\"pid\":0,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":2,\"name\":\"开发人员\",\"desc\":\"开发人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"78,79,80,81,82,73,74,75,76,77,62,63,65,68,69\",\"extAuthIds\":\"\",\"children\":[]},{\"createdAt\":\"2021-11-22 17:07:02\",\"id\":3,\"name\":\"测试人员\",\"desc\":\"测试人员权限\",\"pid\":1,\"defaultRouter\":\"dashboard\",\"menuIds\":\"1,2,3,4,7,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}]},{\"createdAt\":\"2022-10-01 13:35:31\",\"id\":6,\"name\":\"产品同学\",\"desc\":\"产品同学权限\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"1,2,3,4,5,6,7,11,16,18,19,20,8\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":[]}],\"total\":999,\"page\":1,\"pageSize\":10},\"costTime\":8.68608}', '2022-10-30 09:57:57', '2022-10-30 09:57:57', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (241, 1, 'admin', '127.0.0.1', 'GET', '/users', 200, 17.69600, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询用户列表成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '{\"code\":0,\"msg\":\"分页查询用户列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-29 10:53:41\",\"uuid\":\"3e006a79-a28b-4c6f-a084-99ac9c8ea2e9\",\"userName\":\"admin\",\"nickName\":\"超级管理员\",\"phone\":\"15848606970\",\"email\":\"490513975@qq.com\",\"headImg\":\"https://i.postimg.cc/Wb026mfD/logo.jpg\",\"authorityId\":\"1\",\"authorityIds\":[\"1\",\"3\"],\"enable\":1,\"sideMode\":\"dark\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"token\":\"\",\"authority\":{\"createdAt\":\"0001-01-01 00:00:00\",\"id\":0,\"name\":\"\",\"desc\":\"\",\"pid\":0,\"defaultRouter\":\"\",\"menuIds\":\"\",\"apiIds\":\"\",\"extAuthIds\":\"\",\"children\":null},\"authorities\":null}],\"total\":1,\"page\":1,\"pageSize\":10},\"costTime\":17.65504}', '2022-10-30 09:58:00', '2022-10-30 09:58:00', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (242, 1, 'admin', '127.0.0.1', 'GET', '/common/fileList', 200, 36.82893, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询媒体库列表成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '{\"code\":0,\"msg\":\"分页查询媒体库列表成功\",\"data\":{\"list\":[{\"id\":1,\"createdAt\":\"2022-10-23 14:02:24\",\"name\":\"1\",\"url\":\"https://i.postimg.cc/Wb026mfD/logo.jpg\",\"tag\":\"1\",\"key\":\"1\"}],\"total\":1,\"page\":1,\"pageSize\":10},\"costTime\":36.7529}', '2022-10-30 09:58:03', '2022-10-30 09:58:03', NULL);
INSERT INTO `record` (`id`, `userId`, `userName`, `ip`, `method`, `path`, `status`, `costTime`, `agent`, `msg`, `params`, `resp`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (243, 1, 'admin', '127.0.0.1', 'GET', '/recode', 200, 7.82003, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36', '分页查询操作记录成功', '{\"page\":\"1\",\"pageSize\":\"10\"}', '不记录', '2022-10-30 09:58:06', '2022-10-30 09:58:06', NULL);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `userName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录密码',
  `nickName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `headImg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'https://www.pubgmobile.com/cp/pubgpic/images/100001.png' COMMENT '用户头像',
  `authorityId` int NOT NULL DEFAULT '2' COMMENT '用户角色ID',
  `authorityIds` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '多角色id',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否可用',
  `sideMode` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `baseColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#fff' COMMENT '基础颜色',
  `activeColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `createdAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_userName` (`userName`) USING BTREE,
  KEY `idx_deletedAt` (`deletedAt`) USING BTREE,
  KEY `idx_pasword` (`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `uuid`, `userName`, `password`, `nickName`, `headImg`, `authorityId`, `authorityIds`, `phone`, `email`, `enable`, `sideMode`, `baseColor`, `activeColor`, `createdAt`, `updatedAt`, `deletedAt`) VALUES (1, '3e006a79-a28b-4c6f-a084-99ac9c8ea2e9', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://i.postimg.cc/Wb026mfD/logo.jpg', 1, '1,3', '15848606970', '490513975@qq.com', 1, 'dark', '#fff', '#1890ff', '2022-10-29 10:53:41', '2022-10-29 10:53:42', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
