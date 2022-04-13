/*
 Navicat MySQL Data Transfer

 Source Server         : centos_mysql
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 192.168.44.100:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 11/04/2022 21:17:01
*/

CREATE DATABASE IF NOT EXISTS blog;

USE blog;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_bgimage
-- ----------------------------
DROP TABLE IF EXISTS `t_bgimage`;
CREATE TABLE `t_bgimage`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片url',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_bgimage
-- ----------------------------
INSERT INTO `t_bgimage` VALUES (1, 'http://192.168.44.100/images/firstPic/wallhaven-v9vv3p_1647674711.png');
INSERT INTO `t_bgimage` VALUES (2, 'http://192.168.44.100/images/firstPic/wallhaven-l37w6q_1647674789.jpg');
INSERT INTO `t_bgimage` VALUES (3, 'http://192.168.44.100/images/firstPic/wallhaven-72ozj3_1647674771.png');
INSERT INTO `t_bgimage` VALUES (4, 'http://192.168.44.100/images/firstPic/bg24.jpg');
INSERT INTO `t_bgimage` VALUES (5, 'http://192.168.44.100/images/firstPic/bg23.jpg');
INSERT INTO `t_bgimage` VALUES (6, 'http://192.168.44.100/images/firstPic/bg21.jpg');
INSERT INTO `t_bgimage` VALUES (7, 'http://192.168.44.100/images/firstPic/bg20.jpg');
INSERT INTO `t_bgimage` VALUES (8, 'http://192.168.44.100/images/firstPic/bg19.jpg');
INSERT INTO `t_bgimage` VALUES (9, 'http://192.168.44.100/images/firstPic/bg17.png');
INSERT INTO `t_bgimage` VALUES (10, 'http://192.168.44.100/images/firstPic/bg14.jpg');
INSERT INTO `t_bgimage` VALUES (11, 'http://192.168.44.100/images/firstPic/bg13.png');
INSERT INTO `t_bgimage` VALUES (12, 'http://192.168.44.100/images/firstPic/bg12.png');
INSERT INTO `t_bgimage` VALUES (13, 'http://192.168.44.100/images/firstPic/bg9.jpg');
INSERT INTO `t_bgimage` VALUES (14, 'http://192.168.44.100/images/firstPic/bg8.jpg');
INSERT INTO `t_bgimage` VALUES (15, 'http://192.168.44.100/images/firstPic/bg5.jpg');
INSERT INTO `t_bgimage` VALUES (16, 'http://192.168.44.100/images/firstPic/bg3.jpg');
INSERT INTO `t_bgimage` VALUES (17, 'http://192.168.44.100/images/firstPic/bg2.png');
INSERT INTO `t_bgimage` VALUES (18, 'http://192.168.44.100/images/firstPic/bg1.png');
INSERT INTO `t_bgimage` VALUES (19, 'http://192.168.44.100/images/firstPic/200_1647675014.jpg');

-- ----------------------------
-- Table structure for t_blog
-- ----------------------------
DROP TABLE IF EXISTS `t_blog`;
CREATE TABLE `t_blog`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `title` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客标题',
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客内容',
  `first_picture` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客首图',
  `flag` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '原创、转载、翻译等',
  `views` int UNSIGNED NOT NULL COMMENT '浏览次数',
  `appreciation` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否开启赞赏',
  `share_statement` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否可以转载',
  `commentabled` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否可以评论',
  `published` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否发布',
  `recommend` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否推荐',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `type_id` bigint NOT NULL COMMENT '博客类型id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_type_id`(`type_id`) USING BTREE COMMENT '类型id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_blog
-- ----------------------------

-- ----------------------------
-- Table structure for t_blog_tags
-- ----------------------------
DROP TABLE IF EXISTS `t_blog_tags`;
CREATE TABLE `t_blog_tags`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `tag_id` int UNSIGNED NOT NULL COMMENT '标签id',
  `blog_id` int UNSIGNED NOT NULL COMMENT '博客id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blog_id`(`blog_id`) USING BTREE COMMENT '博客id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 211 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_blog_tags
-- ----------------------------

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '电子邮件',
  `content` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `avatar` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图像地址',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `blog_id` int NOT NULL COMMENT '博客id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_blog_id`(`blog_id`) USING BTREE COMMENT '博客id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_comment
-- ----------------------------
INSERT INTO `t_comment` VALUES (1, '小白', 'bai@qq.com', '小白的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (2, '小红', 'hong@qq.com', '小红的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (5, '小蓝', 'lan@qq.com', '小蓝的评论', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 4);
INSERT INTO `t_comment` VALUES (7, '小马', '691639910@qq.com', '博主的评论', 'http://5b0988e595225.cdn.sohucs.com/images/20181103/feaa7d14883047fb81bbaa16f681f583.jpeg', '2021-02-05 11:56:19', 2);
INSERT INTO `t_comment` VALUES (8, '安东尼', '2333@qq.com', '不论是我的世界车水马龙繁华盛世 还是它们都瞬间消失化为须臾 我都会坚定地走向你 不慌张 不犹豫', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-02-05 11:56:19', 6);
INSERT INTO `t_comment` VALUES (9, '亚瑟', 'yase@163.com', 'good!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:24:17', 23);
INSERT INTO `t_comment` VALUES (10, '约翰', 'yuehan@163.com', 'excellent!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:26:47', 23);
INSERT INTO `t_comment` VALUES (11, '达奇', 'daqi@163.com', 'I have a plan! A good plan!', 'http://192.168.44.100/images/firstPic/avatar1.png', '2021-03-31 22:47:36', 23);
INSERT INTO `t_comment` VALUES (15, 'admin', 'test', 'test', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-03 20:11:36', 37);
INSERT INTO `t_comment` VALUES (16, 'admin', 'TEST', 'TEST', 'http://192.168.44.100/firstPic/avatar1.png', '2022-03-03 20:12:52', 37);
INSERT INTO `t_comment` VALUES (17, 'test', 'test', 'admin', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-03 20:17:36', 37);
INSERT INTO `t_comment` VALUES (18, '小红', '1789033422@qq.com', '试一试 ', 'http://192.168.44.100/images/firstPic/avatar1.png', '2022-03-18 12:46:41', 4);

-- ----------------------------
-- Table structure for t_essay
-- ----------------------------
DROP TABLE IF EXISTS `t_essay`;
CREATE TABLE `t_essay`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_essay
-- ----------------------------
INSERT INTO `t_essay` VALUES (1, '2022-04-06 18:09:32', '自信人生两百年，会当水击三千里！');
INSERT INTO `t_essay` VALUES (2, '2022-04-07 18:09:50', '如果你累了，学会休息，而不是放弃!');
INSERT INTO `t_essay` VALUES (3, '2022-04-08 18:10:10', '无论生活多么的复杂，总要保持自己的那一份优雅!');
INSERT INTO `t_essay` VALUES (4, '2021-05-01 18:10:48', '自信人生两百年，会当水击三千里！');
INSERT INTO `t_essay` VALUES (5, '2021-05-01 20:10:48', '如果你累了，学会休息，而不是放弃!');
INSERT INTO `t_essay` VALUES (6, '2022-04-06 19:09:15', '学会享受孤独！');

-- ----------------------------
-- Table structure for t_leave_message
-- ----------------------------
DROP TABLE IF EXISTS `t_leave_message`;
CREATE TABLE `t_leave_message`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱地址',
  `url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '个人网站',
  `content` varchar(1600) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `avatar` tinyint UNSIGNED NOT NULL COMMENT '头像',
  `create_time` datetime NOT NULL COMMENT '评论时间',
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父评论，0表示没有',
  `top_parent_id` int NOT NULL DEFAULT 0 COMMENT '顶级评论',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论状态，0：未审核，1：审核通过，2：审核不通过',
  `ip` char(21) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论ip',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_leave_message
-- ----------------------------
INSERT INTO `t_leave_message` VALUES (1, '管理员', '123456@qq.com', '', 'This is a test.', 10, '2022-03-29 13:28:02', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (2, '管理员', '123456@qq.com', '', 'Test2', 1, '2022-03-29 13:29:36', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (3, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub comment.', 6, '2022-03-30 03:07:59', 1, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (4, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub  Comment2.', 2, '2022-03-30 03:09:32', 1, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (5, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 3, '2022-03-30 03:09:32', 3, 1, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (6, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 4, '2022-03-30 03:09:32', 2, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (7, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 5, '2022-03-30 03:09:32', 2, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (8, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 7, '2022-03-30 03:09:32', 0, 0, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (9, '管理员', '456789@qq.com', 'https://www.baidu.com', 'Sub Sub Comment.', 8, '2022-03-30 03:09:32', 6, 2, 1, '127.0.0.1:8082');
INSERT INTO `t_leave_message` VALUES (10, '管理员测试', 'aaa@qq.com', '', '测试发布子级评论', 2, '2022-03-30 09:28:30', 3, 1, 0, '192.168.44.1:56811');
INSERT INTO `t_leave_message` VALUES (11, '管理员测试', 'aaa@qq.com', '', '测试子级评论发布', 26, '2022-03-30 09:31:48', 3, 1, 1, '192.168.44.1:56979');
INSERT INTO `t_leave_message` VALUES (12, '无名', 'ccdd@qq.com', '', '测试', 25, '2022-03-30 11:52:18', 8, 8, 1, '192.168.44.1:63436');
INSERT INTO `t_leave_message` VALUES (13, '大名', 'eeff@qq.com', '', '测试', 12, '2022-03-30 11:55:41', 8, 8, 1, '192.168.44.1:63532');
INSERT INTO `t_leave_message` VALUES (14, '测试人员', 'fff@qq.com', '', 'Test', 12, '2022-03-30 12:14:49', 0, 0, 1, '192.168.44.1:64602');
INSERT INTO `t_leave_message` VALUES (15, '管理猿', '6666@qq.com', '', 'Success!', 27, '2022-03-30 13:48:32', 0, 0, 1, '192.168.44.1:55829');

-- ----------------------------
-- Table structure for t_motto
-- ----------------------------
DROP TABLE IF EXISTS `t_motto`;
CREATE TABLE `t_motto`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `ch` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '中文',
  `en` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '英文',
  `create_time` datetime NOT NULL COMMENT '创建日期',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of t_motto
-- ----------------------------
INSERT INTO `t_motto` VALUES (1, '心有多大，舞台就有多大!', 'How large the heart is,how large the stage is!', '2022-04-05 20:45:20');
INSERT INTO `t_motto` VALUES (2, '人生最可悲的事莫过于胸怀大志，却虚度光阴!', 'The biggest shame for a man is to waste his time having fun with the great dream!', '2022-04-06 20:45:25');
INSERT INTO `t_motto` VALUES (3, '无论生活多么的复杂，总要保持自己的那一份优雅!', 'No matter how complicated the life is ，Please be graceful all the time!', '2022-04-04 20:45:29');
INSERT INTO `t_motto` VALUES (4, '不需要任何人满意，但求于无愧于心!', 'I don\'t need to please everyone ，I just want to do everythingwithout regret for myself!', '2022-04-03 20:45:33');
INSERT INTO `t_motto` VALUES (5, '越努力，越幸运!', 'The harder you are working, the luckier you will be!', '2022-04-02 20:45:37');
INSERT INTO `t_motto` VALUES (6, '当你的才华还撑不起你的野心时，就要静下心来学习；当你的能力还撑不起你的梦想时，就要努力的去工作!', 'Please study hard，if your talent is not enough to achieve your ambition。 Please work hard，if your ability is not enough to make your dream come true!', '2022-04-13 20:45:40');
INSERT INTO `t_motto` VALUES (7, '如果你累了，学会休息，而不是放弃!', 'If you get tired, learn to rest, not to quit!', '2022-04-11 20:45:45');

-- ----------------------------
-- Table structure for t_resource_category
-- ----------------------------
DROP TABLE IF EXISTS `t_resource_category`;
CREATE TABLE `t_resource_category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类别名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_resource_category
-- ----------------------------
INSERT INTO `t_resource_category` VALUES (1, 'Golang学习资料');
INSERT INTO `t_resource_category` VALUES (2, '优质博客');
INSERT INTO `t_resource_category` VALUES (3, '工具网站');
INSERT INTO `t_resource_category` VALUES (21, '前端开发');

-- ----------------------------
-- Table structure for t_resource_link
-- ----------------------------
DROP TABLE IF EXISTS `t_resource_link`;
CREATE TABLE `t_resource_link`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接名称',
  `desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '链接描述',
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'url',
  `category_id` int UNSIGNED NOT NULL COMMENT '类别id',
  `icon` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'icon',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_category_id`(`category_id`) USING BTREE COMMENT '类别id索引'
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of t_resource_link
-- ----------------------------
INSERT INTO `t_resource_link` VALUES (1, 'Go语言高级编程', 'Go语言高级编程(Advanced Go Programming)', 'https://chai2010.cn/advanced-go-programming-book/', 1, 'http://192.168.44.100/images/icons/favicon_1.ico');
INSERT INTO `t_resource_link` VALUES (2, 'Go标准库', 'Golang标准库文档', 'http://doc.golang.ltd/', 1, ' ');
INSERT INTO `t_resource_link` VALUES (3, '李文周的博客', '李文周的博客，里面含有Golang教程', 'https://www.liwenzhou.com/archives/', 1, 'http://192.168.44.100/images/icons/favicon_2.ico');
INSERT INTO `t_resource_link` VALUES (4, '鸟窝', '一位大佬的博客', 'https://colobu.com/', 1, '');
INSERT INTO `t_resource_link` VALUES (5, 'Go语言爱好者周刊', 'Go语言周刊，每周日发布', 'https://github.com/polaris1119/golangweekly', 1, '');
INSERT INTO `t_resource_link` VALUES (6, 'Go语言设计与实现', '深入了解Go语言的实现', 'https://draveness.me/golang/', 1, 'http://192.168.44.100/images/icons/favicon_3.png');
INSERT INTO `t_resource_link` VALUES (7, '程序员的工具箱', '各种程序员使用的在线工具', 'https://tool.lu/', 3, 'http://192.168.44.100/images/icons/favicon_4.ico');
INSERT INTO `t_resource_link` VALUES (8, '编程自学之路', '编程相关知识', 'https://www.r2coding.com/', 2, 'http://192.168.44.100/images/icons/favicon_5.ico');
INSERT INTO `t_resource_link` VALUES (9, '九月网址导航', '各种工具，PPT、高清图库...', 'https://jiuyueppt.com/', 3, 'http://192.168.44.100/images/icons/favicon_6.ico');
INSERT INTO `t_resource_link` VALUES (13, '在线画图工具', '一个在线的画图工具', 'https://www.processon.com/', 3, 'http://192.168.44.100/images/icons/favicon_1649424714.ico');
INSERT INTO `t_resource_link` VALUES (14, 'Animate.css', '一个css动画库，非常好用', 'https://animate.style/', 21, 'http://192.168.44.100/images/icons/favicon_1649486228.ico');
INSERT INTO `t_resource_link` VALUES (15, 'jQuery插件库', '提供各种最新的jQuery插件，各种css特效等', 'https://www.jq22.com/code2732', 21, 'http://192.168.44.100/images/icons/favicon_1649486404.ico');

-- ----------------------------
-- Table structure for t_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_tag`;
CREATE TABLE `t_tag`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_tag
-- ----------------------------
INSERT INTO `t_tag` VALUES (3, '前端');
INSERT INTO `t_tag` VALUES (4, '后端');
INSERT INTO `t_tag` VALUES (5, 'springboot');
INSERT INTO `t_tag` VALUES (6, 'java');
INSERT INTO `t_tag` VALUES (16, '应用部署');
INSERT INTO `t_tag` VALUES (17, '容器');

-- ----------------------------
-- Table structure for t_type
-- ----------------------------
DROP TABLE IF EXISTS `t_type`;
CREATE TABLE `t_type`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '博客类型名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_type
-- ----------------------------
INSERT INTO `t_type` VALUES (1, '博客系统');
INSERT INTO `t_type` VALUES (3, 'Mybatis');
INSERT INTO `t_type` VALUES (7, 'Springboot');
INSERT INTO `t_type` VALUES (8, 'Maven');
INSERT INTO `t_type` VALUES (17, 'Docker');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '昵称',
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮箱',
  `avatar` varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '头像地址',
  `type` int UNSIGNED NOT NULL COMMENT '用户类型，管理员、普通用户',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `github` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'github地址',
  `csdn` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'csdn地址',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE COMMENT '用户名索引'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of t_user user admin password:123456
-- ----------------------------
INSERT INTO `t_user` VALUES (1, 'unknown', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '123456@qq.com', 'http://192.168.44.100/images/firstPic/avatar1.png', 1, '2021-02-01 18:25:26', '2022-03-02 11:52:19', 'https://www.baidu.com/', 'https://blog.csdn.net/Peerless__?type=blog');

SET FOREIGN_KEY_CHECKS = 1;
