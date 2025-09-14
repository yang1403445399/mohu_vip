-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        9.1.0 - MySQL Community Server - GPL
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.11.0.7065
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 mohu_vip 的数据库结构
CREATE DATABASE IF NOT EXISTS `mohu_vip` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `mohu_vip`;

-- 导出  表 mohu_vip.tp_article 结构
CREATE TABLE IF NOT EXISTS `tp_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `cate_id` int unsigned NOT NULL COMMENT '栏目ID',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
  `keywords` varchar(255) DEFAULT NULL COMMENT '关键词',
  `intro` varchar(255) DEFAULT NULL COMMENT '描述',
  `thumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '缩略图',
  `author` varchar(255) DEFAULT NULL COMMENT '作者',
  `source` varchar(50) DEFAULT NULL COMMENT '来源',
  `content` text NOT NULL COMMENT '内容',
  `hits` int unsigned NOT NULL DEFAULT '0' COMMENT '点击量',
  `hot` int unsigned NOT NULL DEFAULT '0' COMMENT '热门',
  `top` int unsigned NOT NULL DEFAULT '0' COMMENT '置顶',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `state` int unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `FK_tp_article_tp_user` (`user_id`) USING BTREE,
  KEY `FK_tp_article_tp_column` (`cate_id`) USING BTREE,
  CONSTRAINT `FK_tp_article_tp_cate` FOREIGN KEY (`cate_id`) REFERENCES `tp_cate` (`id`),
  CONSTRAINT `FK_tp_article_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文章';

-- 正在导出表  mohu_vip.tp_article 的数据：~3 rows (大约)
INSERT INTO `tp_article` (`id`, `user_id`, `cate_id`, `title`, `keywords`, `intro`, `thumb`, `author`, `source`, `content`, `hits`, `hot`, `top`, `sort`, `uuid`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, 8, '阿松大阿松大啊', '阿松大阿三的', '撒旦阿三', '452452', '阿松大', '阿松大啊', '啊实打实大', 0, 0, 0, 0, 'bf87435a-a336-4dd9-83c0-2bc07abc6aca', '2024-11-05 10:02:26', '2024-11-05 10:02:27', 1),
	(2, 1, 8, '阿松大阿松大啊', '阿松大阿三的', '撒旦阿三', '452452', '阿松大', '阿松大啊', '啊实打实大', 0, 0, 0, 0, 'bf87435a-a336-4dd9-83c0-2bc07abc6aca', '2025-07-05 10:02:26', '2025-07-05 10:02:27', 1),
	(3, 1, 8, '阿松大阿松大啊', '阿松大阿三的', '撒旦阿三', '452452', '阿松大', '阿松大啊', '啊实打实大', 0, 0, 0, 0, 'bf87435a-a336-4dd9-83c0-2bc07abc6aca', '2025-09-05 10:02:26', '2025-09-05 10:02:27', 1);

-- 导出  表 mohu_vip.tp_banner 结构
CREATE TABLE IF NOT EXISTS `tp_banner` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int unsigned NOT NULL COMMENT '类型',
  `name` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片',
  `url` varchar(255) DEFAULT NULL COMMENT '链接',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `type_id` (`type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='轮播图';

-- 正在导出表  mohu_vip.tp_banner 的数据：~4 rows (大约)
INSERT INTO `tp_banner` (`id`, `type_id`, `name`, `src`, `url`, `sort`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, '1111', 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg', '1111', 0, '2025-09-14 18:18:45', '2025-09-12 14:56:05', 1),
	(2, 1, '4444', 'http://127.0.0.1:3000/uploads/20250911/1757562278_企业微信截图_20250910103207 - 副本 (2).png', '2222', 0, '2025-09-14 18:18:53', '2025-09-14 17:54:30', 1),
	(3, 1, '3333', 'http://127.0.0.1:3000/uploads/20250912/1757644283_企业微信截图_20250912103046.png', '3333', 0, '2025-09-14 18:19:00', '2025-09-14 17:55:22', 1),
	(4, 2, '2222', 'http://127.0.0.1:3000/uploads/20250911/1757562240_1 - 副本.jpg', '4444', 0, '2025-09-14 18:19:07', '2025-09-14 18:03:20', 1);

-- 导出  表 mohu_vip.tp_banner_type 结构
CREATE TABLE IF NOT EXISTS `tp_banner_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `create_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT (now()) COMMENT '更新时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  mohu_vip.tp_banner_type 的数据：~2 rows (大约)
INSERT INTO `tp_banner_type` (`id`, `name`, `create_at`, `update_at`, `state`) VALUES
	(1, '电脑端', '2025-09-14 11:46:48', '2025-09-14 11:46:47', 1),
	(2, '移动端', '2025-09-14 11:46:48', '2025-09-14 11:46:47', 1);

-- 导出  表 mohu_vip.tp_basic 结构
CREATE TABLE IF NOT EXISTS `tp_basic` (
  `site` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  UNIQUE KEY `label` (`label`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='基本信息';

-- 正在导出表  mohu_vip.tp_basic 的数据：~15 rows (大约)
INSERT INTO `tp_basic` (`site`, `label`, `value`) VALUES
	('site', 'site_addr', ''),
	('site', 'site_cellphone', ''),
	('site', 'site_copyright', ''),
	('site', 'site_email', ''),
	('site', 'site_ewm', 'http://127.0.0.1:3000/uploads/20250911/1757562278_企业微信截图_20250910103207 - 副本.png'),
	('site', 'site_icp', ''),
	('site', 'site_intro', '444'),
	('site', 'site_keywords', '33'),
	('site', 'site_logo', 'http://127.0.0.1:3000/uploads/20250911/1757562278_企业微信截图_20250910103207 - 副本.png'),
	('site', 'site_mode', '1'),
	('site', 'site_name', '111'),
	('site', 'site_person', ''),
	('site', 'site_script', ''),
	('site', 'site_telephone', ''),
	('site', 'site_title', '222');

-- 导出  表 mohu_vip.tp_browse 结构
CREATE TABLE IF NOT EXISTS `tp_browse` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `province` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省',
  `city` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '市',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` int NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='访问';

-- 正在导出表  mohu_vip.tp_browse 的数据：~6 rows (大约)
INSERT INTO `tp_browse` (`id`, `ip`, `province`, `city`, `update_at`, `create_at`, `state`) VALUES
	(1, '111.178.57.181', '湖北省', '孝感市', '2025-08-02 11:41:15', '2025-08-02 11:41:15', 1),
	(2, '111.178.57.181', '湖北省', '孝感市', '2025-08-02 11:41:15', '2025-08-02 11:41:15', 1),
	(3, '111.178.57.181', '湖北省', '孝感市', '2025-08-02 11:41:15', '2025-08-02 11:41:15', 1),
	(4, '111.178.57.181', '湖北省', '孝感市', '2025-08-02 11:41:15', '2025-08-02 11:41:15', 1),
	(5, '111.178.57.181', '湖南省', '湘潭市', '2025-09-02 11:41:15', '2025-09-02 11:41:20', 1),
	(6, '111.178.57.181', '湖南省', '湘潭市', '2025-09-02 11:41:15', '2025-09-02 11:41:20', 1);

-- 导出  表 mohu_vip.tp_cate 结构
CREATE TABLE IF NOT EXISTS `tp_cate` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `pattern_id` int unsigned NOT NULL COMMENT '模板ID',
  `parent_id` int unsigned NOT NULL COMMENT '父级ID',
  `name` varchar(50) NOT NULL COMMENT '栏目名',
  `alias` varchar(50) DEFAULT NULL COMMENT '别名',
  `keywords` varchar(255) DEFAULT NULL COMMENT '关键词',
  `intro` varchar(255) DEFAULT NULL COMMENT '描述',
  `picture` varchar(255) DEFAULT NULL COMMENT '缩略图',
  `content` text COMMENT '内容',
  `level` int unsigned NOT NULL DEFAULT '0' COMMENT '层级',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `size` int unsigned NOT NULL COMMENT '分页大小',
  `uuid` varchar(50) NOT NULL,
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `state` int unsigned NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `FK_tp_column_tp_page` (`pattern_id`) USING BTREE,
  KEY `FK_tp_column_tp_user` (`user_id`) USING BTREE,
  CONSTRAINT `FK_tp_cate_tp_pattern` FOREIGN KEY (`pattern_id`) REFERENCES `tp_pattern` (`id`),
  CONSTRAINT `FK_tp_cate_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='栏目';

-- 正在导出表  mohu_vip.tp_cate 的数据：~8 rows (大约)
INSERT INTO `tp_cate` (`id`, `user_id`, `pattern_id`, `parent_id`, `name`, `alias`, `keywords`, `intro`, `picture`, `content`, `level`, `sort`, `size`, `uuid`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, 2, 0, '开发工具', 'kaifagongju', '开发工具', '开发工具', '', '', 0, 0, 12, '5755cf91-5b3d-4f61-8be0-c73e5d52caa1', '2024-02-26 16:44:44', '2024-02-26 08:29:52', 1),
	(2, 1, 2, 1, '前端开发', 'qianduankaifa', '前端开发', '前端开发', '', '', 1, 0, 12, '718f138b-810b-4e05-8e51-0099aec06ce1', '2024-03-05 14:00:58', '2024-02-26 08:38:58', 1),
	(3, 1, 2, 1, '后端开发', 'houduankaifa', '后端开发', '后端开发', '', '', 1, 0, 12, '8c3c9ca4-a609-4123-9561-73f2ed9f98fc', '2024-02-26 11:13:49', '2024-02-26 11:13:49', 1),
	(4, 1, 2, 1, 'CSS工具', 'cssgongju', 'CSS工具', 'CSS工具', '', '', 1, 0, 12, '4bd6827a-729e-4709-89c3-6b7035e126f0', '2024-02-26 11:14:41', '2024-02-26 11:14:07', 1),
	(5, 1, 2, 1, 'UI灵感', 'uilinggan', 'UI灵感', 'UI灵感', '', '', 1, 0, 12, '044cc9c2-c637-465d-8f62-5630395a703e', '2024-02-26 11:15:08', '2024-02-26 11:15:08', 1),
	(6, 1, 2, 1, '图片工具', 'tupiangongju', '图片工具', '图片工具', '', '', 1, 0, 12, 'cedc416e-ce85-4e80-aa03-cdd47e9107f4', '2024-02-26 11:15:30', '2024-02-26 11:15:30', 1),
	(7, 1, 2, 1, 'GIT工具', 'gitgongju', 'GIT工具', 'GIT工具', '', '', 1, 0, 12, '04bf5ff5-767b-4948-9d44-678c8701ac1b', '2024-02-26 11:15:55', '2024-02-26 11:15:55', 1),
	(8, 1, 2, 0, '开发笔记', 'kaifabiji', '开发笔记', '开发笔记', '', '', 0, 0, 12, '691cdcf3-7fe9-4c17-9aed-53ecfcd939f4', '2024-03-05 14:00:19', '2024-02-26 11:16:45', 1);

-- 导出  表 mohu_vip.tp_file_type 结构
CREATE TABLE IF NOT EXISTS `tp_file_type` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `state` int unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件类型';

-- 正在导出表  mohu_vip.tp_file_type 的数据：~4 rows (大约)
INSERT INTO `tp_file_type` (`id`, `name`, `update_at`, `create_at`, `state`) VALUES
	(1, '图片', '2025-09-09 17:01:17', '2025-09-09 17:01:18', 1),
	(2, '视频', '2025-09-09 17:01:17', '2025-09-09 17:01:18', 1),
	(3, '音频', '2025-09-09 17:01:17', '2025-09-09 17:01:18', 1),
	(4, '文档', '2025-09-09 17:01:17', '2025-09-09 17:01:18', 1);

-- 导出  表 mohu_vip.tp_gbook 结构
CREATE TABLE IF NOT EXISTS `tp_gbook` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `mobile` varchar(50) NOT NULL COMMENT '手机号',
  `email` varchar(80) DEFAULT NULL COMMENT '电子邮箱',
  `addr` text COMMENT '地址',
  `content` text COMMENT '内容',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `state` int unsigned NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='留言';

-- 正在导出表  mohu_vip.tp_gbook 的数据：~0 rows (大约)

-- 导出  表 mohu_vip.tp_image 结构
CREATE TABLE IF NOT EXISTS `tp_image` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户',
  `column_id` int unsigned NOT NULL COMMENT '栏目',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '网址',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `column_id` (`column_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `FK_tp_image_tp_image_column` FOREIGN KEY (`column_id`) REFERENCES `tp_image_column11111` (`id`),
  CONSTRAINT `FK_tp_image_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='文件';

-- 正在导出表  mohu_vip.tp_image 的数据：~10 rows (大约)
INSERT INTO `tp_image` (`id`, `user_id`, `column_id`, `name`, `url`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, 1, '1757562240_1 - 副本 (2).jpg', 'http://127.0.0.1:3000/uploads/20250911/1757562240_1 - 副本 (2).jpg', '2025-09-11 11:44:01', '2025-09-11 11:44:01', 1),
	(2, 1, 1, '1757562240_1.jpg', 'http://127.0.0.1:3000/uploads/20250911/1757562240_1.jpg', '2025-09-11 11:44:01', '2025-09-11 11:44:01', 1),
	(3, 1, 1, '1757562240_1 - 副本.jpg', 'http://127.0.0.1:3000/uploads/20250911/1757562240_1 - 副本.jpg', '2025-09-11 11:44:01', '2025-09-11 11:44:01', 1),
	(4, 1, 1, '1757562278_企业微信截图_20250910103207 - 副本 (2).png', 'http://127.0.0.1:3000/uploads/20250911/1757562278_企业微信截图_20250910103207 - 副本 (2).png', '2025-09-11 11:44:38', '2025-09-11 11:44:38', 1),
	(6, 1, 1, '1757562278_企业微信截图_20250910103207 - 副本.png', 'http://127.0.0.1:3000/uploads/20250911/1757562278_企业微信截图_20250910103207 - 副本.png', '2025-09-11 11:44:38', '2025-09-11 11:44:38', 1),
	(7, 1, 1, '1757644265_企业微信截图_20250912103046.png', 'http://127.0.0.1:3000/uploads/20250912/1757644265_企业微信截图_20250912103046.png', '2025-09-12 10:31:05', '2025-09-12 10:31:05', 1),
	(8, 1, 1, '1757644275_企业微信截图_20250912103046 - 副本.png', 'http://127.0.0.1:3000/uploads/20250912/1757644275_企业微信截图_20250912103046 - 副本.png', '2025-09-12 10:31:15', '2025-09-12 10:31:15', 1),
	(9, 1, 1, '1757644275_企业微信截图_20250912103046.png', 'http://127.0.0.1:3000/uploads/20250912/1757644275_企业微信截图_20250912103046.png', '2025-09-12 10:31:15', '2025-09-12 10:31:15', 1),
	(10, 1, 1, '1757644283_企业微信截图_20250912103046 - 副本.png', 'http://127.0.0.1:3000/uploads/20250912/1757644283_企业微信截图_20250912103046 - 副本.png', '2025-09-12 10:31:23', '2025-09-12 10:31:23', 1),
	(11, 1, 1, '1757644283_企业微信截图_20250912103046.png', 'http://127.0.0.1:3000/uploads/20250912/1757644283_企业微信截图_20250912103046.png', '2025-09-12 10:31:23', '2025-09-12 10:31:23', 1);

-- 导出  表 mohu_vip.tp_image_column 结构
CREATE TABLE IF NOT EXISTS `tp_image_column` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `name` char(50) NOT NULL,
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP,
  `create_at` datetime NOT NULL DEFAULT (now()),
  `state` tinyint unsigned DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `FK_tp_image_column_copy_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 正在导出表  mohu_vip.tp_image_column 的数据：~2 rows (大约)
INSERT INTO `tp_image_column` (`id`, `user_id`, `name`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, '风景1', '2025-09-14 17:20:40', '2025-09-14 17:14:34', 1),
	(2, 1, '科技', '2025-09-14 17:14:47', '2025-09-14 17:14:47', 1);

-- 导出  表 mohu_vip.tp_link 结构
CREATE TABLE IF NOT EXISTS `tp_link` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `picture` varchar(255) DEFAULT NULL COMMENT '图片',
  `url` varchar(255) NOT NULL COMMENT '链接',
  `uuid` varchar(50) NOT NULL,
  `sort` int unsigned NOT NULL COMMENT '排序',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `state` int unsigned NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `FK_tp_links_tp_user` (`user_id`) USING BTREE,
  CONSTRAINT `FK_tp_link_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='友链';

-- 正在导出表  mohu_vip.tp_link 的数据：~0 rows (大约)

-- 导出  表 mohu_vip.tp_log 结构
CREATE TABLE IF NOT EXISTS `tp_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户',
  `route` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路由',
  `method` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '方法',
  `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '参数',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '备注',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `FK_tp_logs_tp_user` (`user_id`) USING BTREE,
  CONSTRAINT `FK_tp_logs_tp_user` FOREIGN KEY (`user_id`) REFERENCES `tp_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='日志';

-- 正在导出表  mohu_vip.tp_log 的数据：~18 rows (大约)
INSERT INTO `tp_log` (`id`, `user_id`, `route`, `method`, `params`, `remark`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, '1111', '2222', '3333', '77777777777777777777777777777777777777777777777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(2, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(3, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(4, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(5, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(6, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(7, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(8, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(9, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(10, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(11, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(12, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(13, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(14, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(15, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(16, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(17, 1, '1111', '2222', '3333', '7777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1),
	(18, 1, '1111', '2222', '3333', '77777777777777777777777777777777777777777777777', '2025-09-05 16:51:32', '2025-09-05 16:51:33', 1);

-- 导出  表 mohu_vip.tp_menu 结构
CREATE TABLE IF NOT EXISTS `tp_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int unsigned NOT NULL COMMENT '上级',
  `name` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路径',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- 正在导出表  mohu_vip.tp_menu 的数据：~9 rows (大约)
INSERT INTO `tp_menu` (`id`, `parent_id`, `name`, `icon`, `path`, `update_at`, `create_at`, `state`) VALUES
	(1, 0, '网站概况', 'Monitor', '/home', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(2, 0, '网站设置', 'Setting', '/setup', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(3, 2, '基本设置', NULL, '/setup/basic', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(4, 2, '轮播图', NULL, '/setup/banner', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(5, 0, '栏目管理', 'CopyDocument', '/column', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(6, 0, '文章管理', 'Document', '/article', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(7, 0, '留言管理', 'ChatDotSquare', '/comment', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(8, 0, '友链管理', 'Connection', '/link', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1),
	(9, 0, '用户管理', 'User', '/user', '2025-09-09 09:54:49', '2025-09-09 09:54:53', 1);

-- 导出  表 mohu_vip.tp_user 结构
CREATE TABLE IF NOT EXISTS `tp_user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int unsigned NOT NULL DEFAULT '1' COMMENT '角色',
  `username` char(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `nickname` char(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '电话',
  `uuid` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `role_id` (`role_id`),
  CONSTRAINT `FK_tp_user_tp_user_role` FOREIGN KEY (`role_id`) REFERENCES `tp_user_role` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

-- 正在导出表  mohu_vip.tp_user 的数据：~0 rows (大约)
INSERT INTO `tp_user` (`id`, `role_id`, `username`, `password`, `nickname`, `avatar`, `mobile`, `uuid`, `update_at`, `create_at`, `state`) VALUES
	(1, 1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'mohu', '/static/images/avatar_pic.png', '19151004901', 'bf87435a-a336-4dd9-83c0-2bc07abc6aca', '2024-02-28 15:17:53', '2023-09-13 08:27:27', 1);

-- 导出  表 mohu_vip.tp_user_role 结构
CREATE TABLE IF NOT EXISTS `tp_user_role` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` char(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `update_at` datetime NOT NULL DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT (now()) COMMENT '创建时间',
  `state` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色';

-- 正在导出表  mohu_vip.tp_user_role 的数据：~2 rows (大约)
INSERT INTO `tp_user_role` (`id`, `name`, `update_at`, `create_at`, `state`) VALUES
	(1, '管理员', '2025-08-22 09:50:22', '2025-08-22 09:50:15', 1),
	(2, '用户', '2025-08-22 09:51:55', '2025-08-22 09:51:51', 1);

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
