/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : dbMagicAdmin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 30/09/2022 17:47:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `menuId` int(11) NOT NULL COMMENT '所属菜单id',
  `menuName` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求类型',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api名称',
  `route` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由路径',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of api
-- ----------------------------

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority`  (
  `createdAt` datetime(0) NULL DEFAULT NULL,
  `updatedAt` datetime(0) NULL DEFAULT NULL,
  `deletedAt` datetime(0) NULL DEFAULT NULL,
  `authorityId` int(11) NOT NULL COMMENT '角色ID',
  `authorityName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parentId` int(11) NULL DEFAULT NULL COMMENT '父角色ID',
  `defaultRouter` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  `menuIds` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单ID字符串拼接',
  PRIMARY KEY (`authorityId`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authorityId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authority
-- ----------------------------
INSERT INTO `authority` VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 1, '超级管理员', 0, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');
INSERT INTO `authority` VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 2, '普通管理员', 1, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');
INSERT INTO `authority` VALUES ('2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, '测试人员', 0, 'dashboard', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25');

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `pid` int(11) NULL DEFAULT NULL COMMENT '父级ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文名',
  `value` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'val值',
  `super` tinyint(4) NOT NULL DEFAULT 0 COMMENT '超管才有权限 1-有 2-无',
  `sort` int(11) NOT NULL DEFAULT 50 COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `createdAt` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_val`(`value`) USING BTREE COMMENT '字典目录val唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典目录表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO `dictionary` VALUES (9, 0, '菜单权限类型', 'aaa', 1, 50, 'aaaa', '2022-09-30 16:22:32', '2022-09-30 08:22:32', NULL);
INSERT INTO `dictionary` VALUES (16, 0, 'vvv', 'vvv', 2, 50, 'vvv', '2022-09-30 16:18:08', '2022-09-30 08:05:23', NULL);
INSERT INTO `dictionary` VALUES (19, 0, 'bbb', 'bbb', 2, 50, 'bbbb', '2022-09-30 16:18:07', '2022-09-30 08:05:26', NULL);
INSERT INTO `dictionary` VALUES (20, 0, 'asd', 'asdf', 2, 50, 'adsf', '2022-09-30 16:22:05', '2022-09-30 08:22:05', NULL);
INSERT INTO `dictionary` VALUES (21, 0, '111112', '111112', 1, 50, '111112', '2022-09-30 16:40:29', '2022-09-30 08:12:52', '2022-09-30 08:40:29');

-- ----------------------------
-- Table structure for dictionaryDetail
-- ----------------------------
DROP TABLE IF EXISTS `dictionaryDetail`;
CREATE TABLE `dictionaryDetail`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(11) NULL DEFAULT NULL COMMENT '所属目录ID',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '中文名',
  `value` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'val值',
  `super` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否只有超管才有权限',
  `sort` int(11) NOT NULL DEFAULT 50 COMMENT '排序',
  `desc` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `createdAt` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_pid`(`pid`) USING BTREE COMMENT '查询索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典详情表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dictionaryDetail
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parentId` int(11) NOT NULL COMMENT '父菜单ID',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单图标',
  `title` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由path',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由name',
  `hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否是普通页面',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `sort` int(11) NOT NULL COMMENT '排序标记',
  `apis` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'API权限',
  `extAuth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '拓展权限',
  `createdAt` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 0, 'setting', '仪表盘', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '[]', '[]', '2022-09-27 15:43:25', '2022-09-27 07:43:26', NULL);
INSERT INTO `menu` VALUES (2, 0, 'message', '关于我们', 'about', 'about', 0, 'view/about/index.vue', 7, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (3, 0, 'user', '超级管理员', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', NULL, '2022-09-27 11:55:35', '2022-09-27 03:55:26', NULL);
INSERT INTO `menu` VALUES (4, 3, 'hot-water', '角色管理', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (5, 3, 'box', '菜单管理', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (6, 3, 'filter', 'API管理', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (7, 3, 'coordinate', '用户管理', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (8, 0, 'basketball', '个人信息', 'person', 'person', 1, 'view/person/person.vue', 4, NULL, NULL, '2022-09-28 12:40:46', '2022-09-28 04:40:47', NULL);
INSERT INTO `menu` VALUES (9, 0, 'chat-square', '示例文件', 'example', 'example', 0, 'view/example/index.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (10, 9, 'chat-line-square', 'excel导入导出', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (11, 9, 'upload', '媒体库（上传下载）', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (12, 9, 'upload', '断点续传', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (13, 9, 'coffee-cup', '客户列表（资源示例）', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (14, 0, 'briefcase', '系统工具', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (15, 14, 'cpu', '代码生成器', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (16, 14, 'magic-stick', '表单生成器', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (17, 14, 'film', '系统配置', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (18, 3, 'files', '字典管理', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/dictionary.vue', 5, NULL, NULL, '2022-09-27 16:11:23', '2022-09-27 08:07:50', NULL);
INSERT INTO `menu` VALUES (19, 3, 'cherry', '字典详情', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/dictionaryDetail.vue', 1, NULL, NULL, '2022-09-28 11:12:56', '2022-09-28 03:12:57', NULL);
INSERT INTO `menu` VALUES (20, 3, 'alarm-clock', '操作历史', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (22, 0, 'add-location', '官方网站', 'https://www.baidu.com', 'https://www.baidu.com', 0, '/', 0, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (23, 0, 'cloudy', '服务器状态', 'state', 'state', 0, 'view/system/state.vue', 6, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (24, 14, 'coin', '自动化代码管理', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (25, 14, 'calendar', '自动化代码（复用）', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (26, 0, 'apple', 'aaabbbbbb', 'aaabbb', 'aaabbb', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (27, 0, 'aim', 'bbbb', 'bbbb', 'bbbb', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `menu` VALUES (28, 27, 'aim', 'ccc', 'ccc', 'ccc', 0, 'view/routerHolder.vue', 50, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `createdAt` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `userName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录密码',
  `nickName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `headImg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://www.pubgmobile.com/cp/pubgpic/images/100001.png' COMMENT '用户头像',
  `authorityId` int(11) NOT NULL DEFAULT 2 COMMENT '用户角色ID',
  `sideMode` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `baseColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `activeColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_userName`(`userName`) USING BTREE,
  INDEX `idx_deletedAt`(`deletedAt`) USING BTREE,
  INDEX `idx_pasword`(`password`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2022-09-17 04:38:45', NULL, NULL, '3e006a79-a28b-4c6f-a084-99ac9c8ea2e9', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '系统管理员', 'https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi.qqkou.com%2Fi%2F1a3077830965x2286750911b26.jpg&refer=http%3A%2F%2Fi.qqkou.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1665981503&t=c8fefca6768618f0b72fe29f770b4225', 1, 'dark', '#fff', '#1890ff');

SET FOREIGN_KEY_CHECKS = 1;


-- 记录表
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
) ENGINE=InnoDB AUTO_INCREMENT=239 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;