/*
 Navicat Premium Data Transfer

 Source Server         : Localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : dbMagicAdmin

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 16/01/2022 13:01:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for authority
-- ----------------------------
DROP TABLE IF EXISTS `authority`;
CREATE TABLE `authority`  (
  `createdAt` datetime(0) NULL DEFAULT NULL,
  `updatedAt` datetime(0) NULL DEFAULT NULL,
  `deletedAt` datetime(0) NULL DEFAULT NULL,
  `authorityId` int(10) NOT NULL COMMENT '角色ID',
  `authorityName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parentId` int(10) NULL DEFAULT NULL COMMENT '父角色ID',
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
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `createdAt` timestamp(0) NULL DEFAULT NULL,
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  `parentId` int(10) NOT NULL COMMENT '父菜单ID',
  `path` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由path',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由name',
  `hidden` tinyint(1) NOT NULL COMMENT '是否在列表隐藏',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `sort` int(10) NOT NULL COMMENT '排序标记',
  `title` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单图标',
  `keepAlive` tinyint(1) NULL DEFAULT NULL COMMENT '是否缓存',
  `closeTab` tinyint(1) NULL DEFAULT NULL COMMENT '自动关闭tab',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deletedAt`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '仪表盘', 'setting', 0, 0);
INSERT INTO `menu` VALUES (2, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'about', 'about', 0, 'view/about/index.vue', 7, '关于我们', 'info', 0, 0);
INSERT INTO `menu` VALUES (3, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '超级管理员', 'user-solid', 0, 0);
INSERT INTO `menu` VALUES (4, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '角色管理', 's-custom', 0, 0);
INSERT INTO `menu` VALUES (5, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '菜单管理', 's-order', 1, 0);
INSERT INTO `menu` VALUES (6, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 'API管理', 's-platform', 1, 0);
INSERT INTO `menu` VALUES (7, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '用户管理', 'coordinate', 0, 0);
INSERT INTO `menu` VALUES (8, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'person', 'person', 1, 'view/person/person.vue', 4, '个人信息', 'message-solid', 0, 0);
INSERT INTO `menu` VALUES (9, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'example', 'example', 0, 'view/example/index.vue', 6, '示例文件', 's-management', 0, 0);
INSERT INTO `menu` VALUES (10, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 9, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 'excel导入导出', 's-marketing', 0, 0);
INSERT INTO `menu` VALUES (11, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 9, 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '媒体库（上传下载）', 'upload', 0, 0);
INSERT INTO `menu` VALUES (12, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 9, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '断点续传', 'upload', 0, 0);
INSERT INTO `menu` VALUES (13, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 9, 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '客户列表（资源示例）', 's-custom', 0, 0);
INSERT INTO `menu` VALUES (14, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '系统工具', 's-cooperation', 0, 0);
INSERT INTO `menu` VALUES (15, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 14, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '代码生成器', 'cpu', 1, 0);
INSERT INTO `menu` VALUES (16, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 14, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '表单生成器', 'magic-stick', 1, 0);
INSERT INTO `menu` VALUES (17, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 14, 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '系统配置', 's-operation', 0, 0);
INSERT INTO `menu` VALUES (18, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '字典管理', 'notebook-2', 0, 0);
INSERT INTO `menu` VALUES (19, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, '字典详情', 's-order', 0, 0);
INSERT INTO `menu` VALUES (20, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '操作历史', 'time', 0, 0);
INSERT INTO `menu` VALUES (22, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '官方网站', 's-home', 0, 0);
INSERT INTO `menu` VALUES (23, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 0, 'state', 'state', 0, 'view/system/state.vue', 6, '服务器状态', 'cloudy', 0, 0);
INSERT INTO `menu` VALUES (24, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 14, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '自动化代码管理', 's-finance', 0, 0);
INSERT INTO `menu` VALUES (25, '2021-11-22 17:07:02', '2021-11-22 17:07:02', NULL, 14, 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '自动化代码（复用）', 's-finance', 0, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `createdAt` timestamp(0) NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP(0),
  `updatedAt` timestamp(0) NULL DEFAULT NULL,
  `deletedAt` timestamp(0) NULL DEFAULT NULL,
  `uuid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户UUID',
  `userName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户登录密码',
  `nickName` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `headImg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://www.pubgmobile.com/cp/pubgpic/images/100001.png' COMMENT '用户头像',
  `authorityId` int(20) NOT NULL DEFAULT 2 COMMENT '用户角色ID',
  `sideMode` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `baseColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `activeColor` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `defaultRouter` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '默认菜单',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_userName`(`userName`) USING BTREE,
  INDEX `idx_deletedAt`(`deletedAt`) USING BTREE,
  INDEX `idx_pasword`(`password`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2021-11-30 19:07:28', '2021-11-24 09:19:08', NULL, '3e006a79-a28b-4c6f-a084-99ac9c8ea2e9', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '系统管理员', 'https://img2.baidu.com/it/u=685222420,2910423078&fm=26&fmt=auto', 1, 'dark', '#fff', '#1890ff', NULL);

SET FOREIGN_KEY_CHECKS = 1;
