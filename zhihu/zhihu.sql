/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : zhihu

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 23/02/2021 14:17:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `className` varchar(32) NOT NULL,
  `title` varchar(32) NOT NULL,
  `createTime` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of categories
-- ----------------------------
BEGIN;
INSERT INTO `categories` VALUES (1, 'fa-glass', '奇趣事2', '');
INSERT INTO `categories` VALUES (2, 'fa-glass', '奇趣事', '2021-02-23 13:04:56');
INSERT INTO `categories` VALUES (3, 'fa-glass3', '奇趣事3', '2021-02-23 14:06:51');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `nickName` varchar(32) NOT NULL COMMENT '用户昵称',
  `email` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `role` varchar(8) NOT NULL COMMENT 'admin normal',
  `avatar` varchar(64) DEFAULT '' COMMENT '头像',
  `status` int(2) NOT NULL COMMENT '0 未激活 1 激活',
  `createTime` varchar(32) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, '张三', '12345678@qq.com', '123', 'admin', '', 1, '');
INSERT INTO `users` VALUES (4, 'coder', 'coder@itcast.cn', '123456', 'admin', '', 0, '2021-02-22 14:47:32');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
