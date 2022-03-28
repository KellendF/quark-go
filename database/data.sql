/*
QuarkAdmin Data

Target Server Type    : MYSQL
Target Server Version : 50734
File Encoding         : 65001

Date: 2022-01-28 09:27:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for action_logs
-- ----------------------------
DROP TABLE IF EXISTS `action_logs`;
CREATE TABLE `action_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `object_id` int(11) DEFAULT NULL,
  `url` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of action_logs
-- ----------------------------

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nickname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT '1',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `wechat_openid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `wechat_unionid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `qq_openid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `weibo_uid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `last_login_ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `last_login_time` timestamp NULL DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admins_username_unique` (`username`),
  UNIQUE KEY `admins_email_unique` (`email`),
  UNIQUE KEY `admins_phone_unique` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES ('1', 'administrator', '超级管理员', 'admin@yourweb.com', '10086', '1', '$2y$10$O8sNV9/QGo2fSNig5XRuLuMmLHHANA5/36IjsMF6z6Ib8HYpaEvma', null, '', '', '', '', '', null, '1', null, null, null, null);

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sort` int(11) DEFAULT '0',
  `group_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` longtext COLLATE utf8mb4_unicode_ci,
  `remark` longtext COLLATE utf8mb4_unicode_ci,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of configs
-- ----------------------------
INSERT INTO `configs` VALUES ('1', '网站名称', 'text', 'WEB_SITE_NAME', '0', '基本', 'QuarkCMS', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('2', '关键字', 'text', 'WEB_SITE_KEYWORDS', '0', '基本', 'QuarkCMS', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('3', '描述', 'textarea', 'WEB_SITE_DESCRIPTION', '0', '基本', 'QuarkCMS', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('4', 'Logo', 'picture', 'WEB_SITE_LOGO', '0', '基本', '', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('5', '统计代码', 'textarea', 'WEB_SITE_SCRIPT', '0', '基本', '', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('6', '网站版权', 'text', 'WEB_SITE_COPYRIGHT', '0', '基本', '© Company 2018', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('7', '开启SSL', 'switch', 'SSL_OPEN', '0', '基本', '0', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('8', '开启网站', 'switch', 'WEB_SITE_OPEN', '0', '基本', '1', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('9', 'KeyID', 'text', 'OSS_ACCESS_KEY_ID', '0', '阿里云存储', '', '你的AccessKeyID', '1', null, null, null);
INSERT INTO `configs` VALUES ('10', 'KeySecret', 'text', 'OSS_ACCESS_KEY_SECRET', '0', '阿里云存储', '', '你的AccessKeySecret', '1', null, null, null);
INSERT INTO `configs` VALUES ('11', 'EndPoint', 'text', 'OSS_ENDPOINT', '0', '阿里云存储', '', '地域节点', '1', null, null, null);
INSERT INTO `configs` VALUES ('12', 'Bucket域名', 'text', 'OSS_BUCKET', '0', '阿里云存储', '', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('13', '自定义域名', 'text', 'OSS_MYDOMAIN', '0', '阿里云存储', '', '例如：oss.web.com', '1', null, null, null);
INSERT INTO `configs` VALUES ('14', '开启云存储', 'switch', 'OSS_OPEN', '0', '阿里云存储', '0', '', '1', null, null, null);
INSERT INTO `configs` VALUES ('15', '开发者模式', 'switch', 'APP_DEBUG', '0', '基本', '1', '', '1', null, null, null);

-- ----------------------------
-- Table structure for failed_jobs
-- ----------------------------
DROP TABLE IF EXISTS `failed_jobs`;
CREATE TABLE `failed_jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `connection` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of failed_jobs
-- ----------------------------

-- ----------------------------
-- Table structure for file_categories
-- ----------------------------
DROP TABLE IF EXISTS `file_categories`;
CREATE TABLE `file_categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `obj_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'ADMINID' COMMENT '图片归属角色的类型（ADMINID/UID/SHOPID/MCHID等）',
  `obj_id` int(11) DEFAULT '0' COMMENT '图片归属角色id，用于将来多用户角色使用',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sort` int(11) DEFAULT '0',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '分类描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of file_categories
-- ----------------------------

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `obj_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'ADMINID' COMMENT '文件归属角色的类型（ADMINID/UID/SHOPID/MCHID等）',
  `obj_id` int(11) DEFAULT '0' COMMENT '文件归属角色id，用于将来统计用户使用空间大小',
  `file_category_id` int(11) DEFAULT '0' COMMENT '文件目录id',
  `sort` int(11) DEFAULT '0',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `size` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件大小',
  `ext` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '文件扩展名',
  `path` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路径',
  `md5` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'md5唯一标识',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of files
-- ----------------------------

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `guard_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT 'default',
  `pid` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) DEFAULT '0',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `show` tinyint(1) NOT NULL DEFAULT '1',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES ('1', '控制台', 'admin', 'icon-home', 'default', '0', '-2', '/dashboard', '1', '1', null, null);
INSERT INTO `menus` VALUES ('2', '主页', 'admin', '', 'engine', '1', '0', 'admin/dashboard/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('13', '管理员', 'admin', 'icon-admin', 'default', '0', '-1', '/admin', '1', '1', null, null);
INSERT INTO `menus` VALUES ('14', '管理员列表', 'admin', '', 'engine', '13', '0', 'admin/admin/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('15', '菜单列表', 'admin', '', 'engine', '25', '0', 'admin/menu/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('16', '权限列表', 'admin', '', 'engine', '13', '0', 'admin/permission/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('17', '角色列表', 'admin', '', 'engine', '13', '0', 'admin/role/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('25', '系统配置', 'admin', 'icon-setting', 'default', '0', '0', '/system', '1', '1', null, null);
INSERT INTO `menus` VALUES ('26', '设置管理', 'admin', '', 'default', '25', '-1', '/system/config', '1', '1', null, null);
INSERT INTO `menus` VALUES ('27', '网站设置', 'admin', '', 'engine', '26', '0', 'admin/webConfig/setting-form', '1', '1', null, null);
INSERT INTO `menus` VALUES ('28', '配置管理', 'admin', '', 'engine', '26', '0', 'admin/config/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('32', '操作日志', 'admin', '', 'engine', '25', '0', 'admin/actionLog/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('33', '附件空间', 'admin', 'icon-attachment', 'default', '0', '0', '/attachment', '1', '1', null, null);
INSERT INTO `menus` VALUES ('34', '文件管理', 'admin', '', 'engine', '33', '0', 'admin/file/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('35', '图片管理', 'admin', '', 'engine', '33', '0', 'admin/picture/index', '1', '1', null, null);
INSERT INTO `menus` VALUES ('36', '我的账号', 'admin', 'icon-user', 'default', '0', '0', '/account', '1', '1', null, null);
INSERT INTO `menus` VALUES ('37', '个人设置', 'admin', '', 'engine', '36', '0', 'admin/account/setting-form', '1', '1', null, null);

-- ----------------------------
-- Table structure for migrations
-- ----------------------------
DROP TABLE IF EXISTS `migrations`;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of migrations
-- ----------------------------
INSERT INTO `migrations` VALUES ('1', '2014_10_12_000000_create_users_table', '1');
INSERT INTO `migrations` VALUES ('2', '2014_10_12_100000_create_password_resets_table', '1');
INSERT INTO `migrations` VALUES ('3', '2019_01_30_080034_create_admins_table', '1');
INSERT INTO `migrations` VALUES ('4', '2019_01_30_080743_create_action_logs_table', '1');
INSERT INTO `migrations` VALUES ('5', '2019_01_30_083510_create_configs_table', '1');
INSERT INTO `migrations` VALUES ('6', '2019_01_30_083614_create_files_table', '1');
INSERT INTO `migrations` VALUES ('7', '2019_01_30_091501_create_menus_table', '1');
INSERT INTO `migrations` VALUES ('8', '2019_02_11_081942_create_pictures_table', '1');
INSERT INTO `migrations` VALUES ('9', '2019_03_08_081738_create_permission_tables', '1');
INSERT INTO `migrations` VALUES ('10', '2019_08_19_000000_create_failed_jobs_table', '1');
INSERT INTO `migrations` VALUES ('11', '2019_10_17_092404_create_picture_categories_table', '1');
INSERT INTO `migrations` VALUES ('12', '2019_10_17_093935_create_file_categories_table', '1');
INSERT INTO `migrations` VALUES ('13', '2019_12_14_000001_create_personal_access_tokens_table', '1');

-- ----------------------------
-- Table structure for model_has_permissions
-- ----------------------------
DROP TABLE IF EXISTS `model_has_permissions`;
CREATE TABLE `model_has_permissions` (
  `permission_id` int(10) unsigned NOT NULL,
  `model_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`model_id`,`model_type`),
  KEY `model_has_permissions_model_id_model_type_index` (`model_id`,`model_type`),
  CONSTRAINT `model_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of model_has_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for model_has_roles
-- ----------------------------
DROP TABLE IF EXISTS `model_has_roles`;
CREATE TABLE `model_has_roles` (
  `role_id` int(10) unsigned NOT NULL,
  `model_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`model_id`,`model_type`),
  KEY `model_has_roles_model_id_model_type_index` (`model_id`,`model_type`),
  CONSTRAINT `model_has_roles_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of model_has_roles
-- ----------------------------

-- ----------------------------
-- Table structure for password_resets
-- ----------------------------
DROP TABLE IF EXISTS `password_resets`;
CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  KEY `password_resets_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of password_resets
-- ----------------------------

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) NOT NULL DEFAULT 0,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `guard_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of permissions
-- ----------------------------

-- ----------------------------
-- Table structure for personal_access_tokens
-- ----------------------------
DROP TABLE IF EXISTS `personal_access_tokens`;
CREATE TABLE `personal_access_tokens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tokenable_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `tokenable_id` bigint(20) unsigned NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abilities` text COLLATE utf8mb4_unicode_ci,
  `last_used_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `personal_access_tokens_token_unique` (`token`),
  KEY `personal_access_tokens_tokenable_type_tokenable_id_index` (`tokenable_type`,`tokenable_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of personal_access_tokens
-- ----------------------------

-- ----------------------------
-- Table structure for picture_categories
-- ----------------------------
DROP TABLE IF EXISTS `picture_categories`;
CREATE TABLE `picture_categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `obj_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'ADMINID' COMMENT '图片归属角色的类型（ADMINID/UID/SHOPID/MCHID等）',
  `obj_id` int(11) DEFAULT '0' COMMENT '图片归属角色id，用于将来多用户角色使用',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sort` int(11) DEFAULT '0',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '分类描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of picture_categories
-- ----------------------------
INSERT INTO `picture_categories` VALUES ('1', 'ADMINID', '1', '默认分类', '0', '默认分类');

-- ----------------------------
-- Table structure for pictures
-- ----------------------------
DROP TABLE IF EXISTS `pictures`;
CREATE TABLE `pictures` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `obj_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'ADMINID' COMMENT '图片归属角色的类型（ADMINID/UID/SHOPID/MCHID等）',
  `obj_id` int(11) DEFAULT '0' COMMENT '图片归属角色id，用于将来统计用户使用空间大小',
  `picture_category_id` int(11) DEFAULT '0' COMMENT '图片目录id',
  `sort` int(11) DEFAULT '0',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `size` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件大小',
  `width` int(11) DEFAULT '0' COMMENT '宽度',
  `height` int(11) DEFAULT '0' COMMENT '高度',
  `ext` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '文件扩展名',
  `path` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路径',
  `md5` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'md5唯一标识',
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of pictures
-- ----------------------------

-- ----------------------------
-- Table structure for role_has_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_has_permissions`;
CREATE TABLE `role_has_permissions` (
  `permission_id` int(10) unsigned NOT NULL,
  `role_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`role_id`),
  KEY `role_has_permissions_role_id_foreign` (`role_id`),
  CONSTRAINT `role_has_permissions_permission_id_foreign` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  CONSTRAINT `role_has_permissions_role_id_foreign` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of role_has_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `guard_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
SET FOREIGN_KEY_CHECKS=1;
