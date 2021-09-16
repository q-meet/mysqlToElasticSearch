/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 192.168.1.12:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 16/09/2021 15:28:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '订单ID',
  `amount` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '订单金额',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 33 CHARACTER SET = utf8 COLLATE = utf8_unicode_ci COMMENT = '订单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, '1', 0.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (2, '2', 0.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (3, '3', 0.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (4, '4', 0.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (5, '1', 22.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (6, '2', 22.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (7, '3', 22.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (8, '4', 22.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (9, '1', 33.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (10, '2', 33.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (11, '3', 33.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (12, '4', 33.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (13, '1', 33.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (14, '2', 33.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (15, '3', 33.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (16, '4', 33.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (17, '1', 1.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (18, '2', 2.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (19, '3', 3.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (20, '4', 4.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (21, '5', 5.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (22, '6', 6.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (23, '7', 7.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (24, '1', 8.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (25, '2', 9.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (26, '3', 10.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (27, '4', 11.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (28, '12', 12.00, '2021-09-14 09:23:31');
INSERT INTO `order` VALUES (29, '13', 13.00, '2021-09-14 00:07:21');
INSERT INTO `order` VALUES (30, '14', 14.00, '2021-09-14 00:08:47');
INSERT INTO `order` VALUES (31, '15', 15.00, '2021-09-14 00:09:34');
INSERT INTO `order` VALUES (32, '16', 16.00, '2021-09-14 09:23:31');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '栏目类型',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  `flag` set('hot','index','recommend') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '',
  `image` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '图片',
  `keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '关键字',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '描述',
  `diyname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '自定义名称',
  `createtime` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `updatetime` int(10) NULL DEFAULT NULL COMMENT '更新时间',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
  `status` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `weigh`(`weigh`, `id`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, 0, 'page', '官方新闻', 'news', 'recommend', '/assets/img/qrcode.png', '', '', 'news', 1495262190, 1495262190, 1, 'normal');
INSERT INTO `category` VALUES (2, 0, 'page', '移动应用', 'mobileapp', 'hot', '/assets/img/qrcode.png', '', '', 'mobileapp', 1495262244, 1495262244, 2, 'normal');
INSERT INTO `category` VALUES (3, 2, 'page', '微信公众号', 'wechatpublic', 'index', '/assets/img/qrcode.png', '', '', 'wechatpublic', 1495262288, 1495262288, 3, 'normal');
INSERT INTO `category` VALUES (4, 2, 'page', 'Android开发', 'android', 'recommend', '/assets/img/qrcode.png', '', '', 'android', 1495262317, 1495262317, 4, 'normal');
INSERT INTO `category` VALUES (5, 0, 'page', '软件产品', 'software', 'recommend', '/assets/img/qrcode.png', '', '', 'software', 1495262336, 1499681850, 5, 'normal');
INSERT INTO `category` VALUES (6, 5, 'page', '网站建站', 'website', 'recommend', '/assets/img/qrcode.png', '', '', 'website', 1495262357, 1495262357, 6, 'normal');
INSERT INTO `category` VALUES (7, 5, 'page', '企业管理软件', 'company', 'index', '/assets/img/qrcode.png', '', '', 'company', 1495262391, 1495262391, 7, 'normal');
INSERT INTO `category` VALUES (8, 6, 'page', 'PC端', 'website-pc', 'recommend', '/assets/img/qrcode.png', '', '', 'website-pc', 1495262424, 1495262424, 8, 'normal');
INSERT INTO `category` VALUES (9, 6, 'page', '移动端', 'website-mobile', 'recommend', '/assets/img/qrcode.png', '', '', 'website-mobile', 1495262456, 1495262456, 9, 'normal');
INSERT INTO `category` VALUES (10, 7, 'page', 'CRM系统 ', 'company-crm', 'recommend', '/assets/img/qrcode.png', '', '', 'company-crm', 1495262487, 1495262487, 10, 'normal');
INSERT INTO `category` VALUES (11, 7, 'page', 'SASS平台软件', 'company-sass', 'recommend', '/assets/img/qrcode.png', '', '', 'company-sass', 1495262515, 1495262515, 11, 'normal');
INSERT INTO `category` VALUES (12, 0, 'test', '测试1', 'test1', 'recommend', '/assets/img/qrcode.png', '', '', 'test1', 1497015727, 1497015727, 12, 'normal');
INSERT INTO `category` VALUES (13, 0, 'test', '测试2', 'test2', 'recommend', '/assets/img/qrcode.png', '', '', 'test2', 1497015738, 1497015738, 13, 'normal');

SET FOREIGN_KEY_CHECKS = 1;
