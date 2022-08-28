/*
Navicat MySQL Data Transfer

Source Server         : JavaEE
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : study

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-08-26 22:07:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(256) NOT NULL COMMENT '名称',
  `code` varchar(64) DEFAULT NULL COMMENT '角色code',
  `gmt_create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted` bigint(20) NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of t_role
-- ----------------------------
INSERT INTO `t_role` VALUES ('1', '超级管理员', 'super_admin', '2021-09-10 19:27:55', '2021-09-10 19:27:55', '0');
INSERT INTO `t_role` VALUES ('2', '管理员', 'admin', '2021-09-10 19:27:55', '2021-09-10 19:27:55', '0');
INSERT INTO `t_role` VALUES ('3', '普通用户', 'user', '2021-09-10 19:27:56', '2021-09-10 19:27:56', '0');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_name` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '密码',
  `real_name` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '真实姓名',
  `phone` bigint(20) DEFAULT NULL COMMENT '手机号',
  `province` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `city` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `county` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `unit_id` bigint(20) DEFAULT NULL COMMENT '单位id',
  `unit_name` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '单位名称',
  `gmt_create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted` bigint(20) NOT NULL DEFAULT '0' COMMENT '是否删除，非0为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1563001523807162371 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='用户表';

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES ('1563001523807162370', 'louzai', '***', '楼', '13123676844', '湖北省', '鄂州市', '葛店开发区', '2', '武汉小米', '2022-08-26 11:12:57', '2022-08-26 11:12:57', '0');

-- ----------------------------
-- Table structure for t_user_role
-- ----------------------------
DROP TABLE IF EXISTS `t_user_role`;
CREATE TABLE `t_user_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `gmt_create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `deleted` bigint(20) NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1563001523903631362 DEFAULT CHARSET=utf8 COMMENT='用户角色关联表';

-- ----------------------------
-- Records of t_user_role
-- ----------------------------
INSERT INTO `t_user_role` VALUES ('1563001523903631361', '1563001523807162370', '2', '2022-08-26 11:12:57', '2022-08-26 11:12:57', '0');
