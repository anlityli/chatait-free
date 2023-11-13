/*
 Navicat Premium Data Transfer

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : chatgpt_db

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 02/08/2023 21:16:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for c_admin
-- ----------------------------
DROP TABLE IF EXISTS `c_admin`;
CREATE TABLE `c_admin` (
  `id` bigint NOT NULL,
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '前台会员ID',
  `admin_name` varchar(20) NOT NULL COMMENT '管理员名',
  `real_name` varchar(128) NOT NULL COMMENT '真实姓名',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `is_enable` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '是否启用',
  `auth_key` varchar(255) DEFAULT NULL COMMENT '授权key',
  `password_hash` varchar(255) NOT NULL COMMENT 'hash密码',
  `password_reset_token` varchar(255) DEFAULT NULL,
  `dont_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否允许删除',
  `login_nums` int unsigned NOT NULL DEFAULT '0' COMMENT '登陆次数',
  `fail_nums` int unsigned NOT NULL DEFAULT '0' COMMENT '失败次数',
  `last_login_ip` varchar(16) DEFAULT NULL COMMENT '最后登录IP',
  `last_login_at` int unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `bind_ip` text COMMENT '绑定IP',
  `create_admin` varchar(20) NOT NULL DEFAULT '0' COMMENT '创建管理员ID',
  `update_admin` varchar(20) DEFAULT '0' COMMENT '更新管理员ID',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `deleted_at` int unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for c_admin_role
-- ----------------------------
DROP TABLE IF EXISTS `c_admin_role`;
CREATE TABLE `c_admin_role` (
  `id` bigint unsigned NOT NULL,
  `role_name` varchar(20) NOT NULL COMMENT '角色名',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `permission` text COMMENT '权限',
  `column_permission` text COMMENT '字段权限',
  `dont_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '不允许删除',
  `create_admin` varchar(20) NOT NULL DEFAULT '0' COMMENT '创建管理员ID',
  `update_admin` varchar(20) DEFAULT NULL COMMENT '更新管理员ID',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for c_config
-- ----------------------------
DROP TABLE IF EXISTS `c_config`;
CREATE TABLE `c_config` (
  `config_name` varchar(32) NOT NULL COMMENT '配置参数名',
  `title` varchar(48) NOT NULL DEFAULT '' COMMENT '标题',
  `unit` varchar(32) NOT NULL DEFAULT '' COMMENT '单位',
  `input_type` tinyint NOT NULL DEFAULT '1' COMMENT '表单类型',
  `options` varchar(4000) NOT NULL DEFAULT '' COMMENT '参数配置的选项',
  `value` varchar(4000) NOT NULL DEFAULT '' COMMENT '配置值',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '类型',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`config_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='系统配置表';

-- ----------------------------
-- Table structure for c_config_baidu
-- ----------------------------
DROP TABLE IF EXISTS `c_config_baidu`;
CREATE TABLE `c_config_baidu` (
  `id` bigint NOT NULL COMMENT 'ID',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `api_key` varchar(255) NOT NULL,
  `secret_key` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '启用状态',
  `call_num` int NOT NULL DEFAULT '0' COMMENT '调用次数',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='百度接口';

-- ----------------------------
-- Table structure for c_config_level
-- ----------------------------
DROP TABLE IF EXISTS `c_config_level`;
CREATE TABLE `c_config_level` (
  `id` int NOT NULL COMMENT 'ID',
  `level_name` varchar(50) NOT NULL COMMENT '级别名称',
  `month_gpt3` int NOT NULL DEFAULT '0' COMMENT '月赠送gpt3次数',
  `month_gpt4` int NOT NULL DEFAULT '0' COMMENT '月赠送gpt4次数',
  `month_midjourney` int NOT NULL DEFAULT '0' COMMENT '月赠送midjourney次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for c_config_midjourney
-- ----------------------------
DROP TABLE IF EXISTS `c_config_midjourney`;
CREATE TABLE `c_config_midjourney` (
  `id` bigint NOT NULL COMMENT 'ID',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `guild_id` varchar(255) NOT NULL COMMENT '服务ID',
  `channel_id` varchar(255) NOT NULL COMMENT '频道ID',
  `user_token` varchar(255) NOT NULL COMMENT '用户Token',
  `mj_bot_id` varchar(255) CHARACTER SET utf8mb4  NOT NULL DEFAULT '' COMMENT 'midjourney的BotId',
  `bot_token` varchar(255) CHARACTER SET utf8mb4  NOT NULL COMMENT '自己的BotToken',
  `session_id` varchar(255) NOT NULL COMMENT 'SessionID',
  `user_agent` varchar(255) NOT NULL COMMENT 'UserAgent',
  `hugging_face_token` varchar(255) NOT NULL DEFAULT '' COMMENT 'HuggingFaceToken',
  `proxy` varchar(255) NOT NULL DEFAULT '' COMMENT '代理地址',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用 1启用 0不启用',
  `listen_model` tinyint(1) NOT NULL DEFAULT '1' COMMENT '监听模式 1userWss 0bot',
  `create_model` varchar(50) NOT NULL DEFAULT 'fast' COMMENT '生成图的模式 fast relax turbo',
  `ws_idle_time` int NOT NULL DEFAULT '3600' COMMENT 'websocket闲置时间秒数',
  `call_num` int NOT NULL DEFAULT '0' COMMENT '接口调用次数',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='Midjourney接口配置表';

-- ----------------------------
-- Table structure for c_config_openai
-- ----------------------------
DROP TABLE IF EXISTS `c_config_openai`;
CREATE TABLE `c_config_openai` (
  `id` bigint NOT NULL COMMENT 'ID',
  `title` varchar(50) NOT NULL COMMENT '标题',
  `api_key` varchar(255) CHARACTER SET utf8mb4  NOT NULL COMMENT 'api_key',
  `proxy` varchar(255) NOT NULL DEFAULT '' COMMENT '代理地址',
  `max_tokens` int NOT NULL DEFAULT '500' COMMENT '最大Token',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用 1启用 0不启用',
  `call_num` int NOT NULL DEFAULT '0' COMMENT '接口调用次数',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='Midjourney接口配置表';

-- ----------------------------
-- Table structure for c_config_pay
-- ----------------------------
DROP TABLE IF EXISTS `c_config_pay`;
CREATE TABLE `c_config_pay` (
  `id` int NOT NULL COMMENT 'ID',
  `api_name` varchar(50) NOT NULL COMMENT '支付方式名称',
  `params` text NOT NULL COMMENT '配置参数json',
  `pay_channel` varchar(1000) CHARACTER SET utf8mb4  NOT NULL DEFAULT '' COMMENT '支付渠道的配置json',
  `frontend_description` varchar(255) NOT NULL DEFAULT '' COMMENT '支付方式描述',
  `backend_description` varchar(255) NOT NULL DEFAULT '' COMMENT '支付方式后台描述',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='支付方式';

-- ----------------------------
-- Table structure for c_config_pay_qr
-- ----------------------------
DROP TABLE IF EXISTS `c_config_pay_qr`;
CREATE TABLE `c_config_pay_qr` (
  `id` int NOT NULL COMMENT 'ID',
  `amount` int NOT NULL COMMENT '金额',
  `pay_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '二维码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `amount` (`amount`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='短信转发器支付方式用到的支付二维码';

-- ----------------------------
-- Table structure for c_config_wallet
-- ----------------------------
DROP TABLE IF EXISTS `c_config_wallet`;
CREATE TABLE `c_config_wallet` (
  `field` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '钱包字段名',
  `wallet_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '钱包名称',
  PRIMARY KEY (`field`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包名称';

-- ----------------------------
-- Table structure for c_conversation
-- ----------------------------
DROP TABLE IF EXISTS `c_conversation`;
CREATE TABLE `c_conversation` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `topic_id` bigint NOT NULL COMMENT '话题',
  `role` varchar(50) NOT NULL COMMENT '对话角色',
  `content` text COMMENT '对话内容',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `topic_id` (`topic_id`),
  KEY `role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='对话表';

-- ----------------------------
-- Table structure for c_conversation_midjourney
-- ----------------------------
DROP TABLE IF EXISTS `c_conversation_midjourney`;
CREATE TABLE `c_conversation_midjourney` (
  `conversation_id` bigint NOT NULL COMMENT '对话id',
  `action_type` int NOT NULL DEFAULT '1' COMMENT '行为类型 1生图 2Upsale 3Variate 4Reroll',
  `file_id` bigint NOT NULL DEFAULT '0' COMMENT '图片文件ID',
  `components` text COLLATE utf8mb4_general_ci COMMENT '附加组件json 用于u,v,r等按钮及记录',
  `error_data` varchar(1000) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '错误信息',
  PRIMARY KEY (`conversation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='对话对应Midjourney信息表';


-- ----------------------------
-- Table structure for c_email_code
-- ----------------------------
DROP TABLE IF EXISTS `c_email_code`;
CREATE TABLE `c_email_code` (
  `id` bigint NOT NULL COMMENT 'ID',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `ip` varchar(32) NOT NULL DEFAULT '' COMMENT 'IP',
  `scenario` varchar(255) NOT NULL DEFAULT '' COMMENT '场景',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '验证码',
  `validate_times` int NOT NULL DEFAULT '0' COMMENT '校验次数',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `ip` (`ip`),
  KEY `scenario` (`scenario`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='邮箱验证码';

-- ----------------------------
-- Table structure for c_file_midjourney
-- ----------------------------
DROP TABLE IF EXISTS `c_file_midjourney`;
CREATE TABLE `c_file_midjourney` (
  `id` bigint NOT NULL COMMENT 'id',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '会员id',
  `queue_id` bigint NOT NULL DEFAULT '0' COMMENT '生成该图片的队列id',
  `file_name` varchar(255) CHARACTER SET utf8mb4  NOT NULL COMMENT '文件名',
  `path` varchar(1000) NOT NULL COMMENT '本地储存路径',
  `thumbnail` varchar(1000) NOT NULL DEFAULT '' COMMENT '缩略图的本地保存路径',
  `prompt` text COMMENT '生成该图片的提示词',
  `mj_file_name` varchar(255) NOT NULL COMMENT 'midjourney的文件名',
  `mj_url` varchar(2000) NOT NULL COMMENT 'midjourney的路径',
  `width` int NOT NULL DEFAULT '0' COMMENT '宽',
  `height` int NOT NULL DEFAULT '0' COMMENT '高',
  `size` int NOT NULL DEFAULT '0' COMMENT '大小',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='midjourney文件表';

-- ----------------------------
-- Table structure for c_hand_out_error_flow
-- ----------------------------
DROP TABLE IF EXISTS `c_hand_out_error_flow`;
CREATE TABLE `c_hand_out_error_flow` (
  `id` bigint NOT NULL COMMENT 'id',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `level_id` int NOT NULL COMMENT '会员级别',
  `error_data` text COLLATE utf8mb4_general_ci COMMENT '错误内容',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='每日给会员赠送钱包余额错误记录';

-- ----------------------------
-- Table structure for c_log_operation
-- ----------------------------
DROP TABLE IF EXISTS `c_log_operation`;
CREATE TABLE `c_log_operation` (
  `id` bigint NOT NULL COMMENT 'ID',
  `status_code` varchar(10) NOT NULL COMMENT '状态码',
  `router` varchar(500) NOT NULL DEFAULT '' COMMENT '请求路径',
  `request_header` text COMMENT '请求头',
  `content` text COMMENT '操作内容',
  `admin_name` varchar(255) NOT NULL COMMENT '管理员名',
  `created_at` int NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='后台操作日志';

-- ----------------------------
-- Table structure for c_pay_flow
-- ----------------------------
DROP TABLE IF EXISTS `c_pay_flow`;
CREATE TABLE `c_pay_flow` (
  `id` bigint NOT NULL COMMENT 'ID',
  `flow_type` int NOT NULL DEFAULT '1' COMMENT '记录类型 1商城订单',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT '目标ID',
  `config_pay_id` int NOT NULL DEFAULT '0' COMMENT '支付接口ID',
  `pay_channel` varchar(64) CHARACTER SET utf8mb4  NOT NULL DEFAULT '' COMMENT '支付渠道以支付接口为准',
  `order_amount` int NOT NULL DEFAULT '0' COMMENT '订单原金额',
  `pay_amount` int NOT NULL DEFAULT '0' COMMENT '实付金额',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0已创建 1已支付 2支付失败 3已过期',
  `payment_response` text CHARACTER SET utf8mb4  COMMENT '调用接口返回内容',
  `payment_fail_reason` varchar(1000) NOT NULL DEFAULT '' COMMENT '调用接口失败原因',
  `notify_response` text COMMENT '接口回调原文',
  `notify_fail_reason` varchar(1000) NOT NULL DEFAULT '' COMMENT '回调失败原因',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  `paid_at` int NOT NULL DEFAULT '0' COMMENT '支付时间',
  `due_expire_at` int NOT NULL DEFAULT '0' COMMENT '应到期时间',
  `expired_at` int NOT NULL DEFAULT '0' COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for c_queue_midjourney
-- ----------------------------
DROP TABLE IF EXISTS `c_queue_midjourney`;
CREATE TABLE `c_queue_midjourney` (
  `id` bigint NOT NULL COMMENT 'ID',
  `conversation_id` bigint DEFAULT NULL COMMENT '对话ID',
  `config_id` bigint NOT NULL COMMENT '接口配置ID',
  `action_type` int NOT NULL DEFAULT '1' COMMENT '行为类型  1生图 2Upscale',
  `nonce` bigint NOT NULL DEFAULT '0' COMMENT 'nonceID',
  `message_id` bigint NOT NULL DEFAULT '0' COMMENT '消息结束ID(生成图片完成时的消息ID)',
  `refer_message_id` bigint NOT NULL DEFAULT '0' COMMENT '提到的消息ID(生图动作为0)',
  `interaction_id` bigint NOT NULL DEFAULT '0' COMMENT '交互ID',
  `refer_index` int NOT NULL DEFAULT '0' COMMENT '处理提到的消息的索引',
  `message_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息hash',
  `message_type` int NOT NULL DEFAULT '0' COMMENT '消息type',
  `message_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '消息内容(提示词内容用于匹配任务)',
  `request_type` int NOT NULL DEFAULT '2' COMMENT '请求消息时用到的类型 2生图 3Upscale 3variation',
  `request_url` varchar(1000) COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口的url',
  `request_data` text COLLATE utf8mb4_general_ci COMMENT '请求接口的数据内容',
  `response_data` text COLLATE utf8mb4_general_ci COMMENT '接口返回的数据内容',
  `error_data` varchar(1000) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '错误数据内容',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0任务进入队列 1 任务开始 2任务正常结束 3任务出错',
  `progress` int NOT NULL DEFAULT '0' COMMENT '任务执行进度',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `started_at` int NOT NULL DEFAULT '0' COMMENT '任务开始时间',
  `ended_at` int NOT NULL DEFAULT '0' COMMENT '任务结束时间',
  `error_at` int NOT NULL DEFAULT '0' COMMENT '任务发生错误时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Midjourney队列表';

-- ----------------------------
-- Table structure for c_shop_goods
-- ----------------------------
DROP TABLE IF EXISTS `c_shop_goods`;
CREATE TABLE `c_shop_goods` (
  `id` bigint NOT NULL COMMENT 'ID',
  `title` varchar(255) NOT NULL COMMENT '商品标题',
  `content` text COMMENT '商品内容',
  `feat_items` text COMMENT '商品特色条目JSON',
  `buy_type` int NOT NULL DEFAULT '1' COMMENT '购买类型 1购买级别 2购买balance 3购买gpt3 4购买gpt4 5购买midjourney',
  `active_level_id` int NOT NULL COMMENT '购买的级别',
  `active_expire_type` int NOT NULL DEFAULT '0' COMMENT '激活有效期类型 0无 1一天 2一月 3一年',
  `active_expire_value` int NOT NULL DEFAULT '0' COMMENT '激活有效期值',
  `buy_value` int NOT NULL DEFAULT '0' COMMENT '购买的提问次数的值分单位',
  `market_price` int NOT NULL DEFAULT '0' COMMENT '市场价',
  `real_price` int NOT NULL DEFAULT '0' COMMENT '实际价格',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否上架',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='商品表';

-- ----------------------------
-- Table structure for c_shop_order
-- ----------------------------
DROP TABLE IF EXISTS `c_shop_order`;
CREATE TABLE `c_shop_order` (
  `id` bigint NOT NULL COMMENT 'ID',
  `order_sn` varchar(64) NOT NULL COMMENT '订单编号',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `order_amount` int NOT NULL DEFAULT '0' COMMENT '订单金额',
  `pay_amount` int NOT NULL DEFAULT '0' COMMENT '实付金额',
  `status` int NOT NULL DEFAULT '0' COMMENT '状态 0创建 1已支付 2已发货 3已收货 4已完成 9已取消',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  `paid_at` int NOT NULL DEFAULT '0' COMMENT '支付时间',
  `due_expire_at` int NOT NULL DEFAULT '0' COMMENT '应过期时间',
  `expired_at` int NOT NULL DEFAULT '0' COMMENT '过期时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `order_sn` (`order_sn`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='商品订单表';

-- ----------------------------
-- Table structure for c_shop_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `c_shop_order_goods`;
CREATE TABLE `c_shop_order_goods` (
  `id` bigint NOT NULL COMMENT 'ID',
  `order_id` bigint NOT NULL COMMENT '订单ID',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `goods_id` bigint NOT NULL COMMENT '商品ID',
  `goods_num` int NOT NULL COMMENT '商品数量',
  `goods_snapshot` text CHARACTER SET utf8mb4  NOT NULL COMMENT '商品快照',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `order_id` (`order_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='订单商品';

-- ----------------------------
-- Table structure for c_topic
-- ----------------------------
DROP TABLE IF EXISTS `c_topic`;
CREATE TABLE `c_topic` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `title` varchar(255) NOT NULL COMMENT '话题标题',
  `type` int NOT NULL DEFAULT '1' COMMENT '话题类型 1:gpt3.5 2:gpt4 3:midjourney',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='话题表';

-- ----------------------------
-- Table structure for c_user
-- ----------------------------
DROP TABLE IF EXISTS `c_user`;
CREATE TABLE `c_user` (
  `id` bigint NOT NULL COMMENT 'ID',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `level_id` int NOT NULL DEFAULT '1' COMMENT '级别ID',
  `level_expire_date` date DEFAULT NULL COMMENT '级别到期日期',
  `level_expire_year` int NOT NULL DEFAULT '0' COMMENT '级别到期日期年',
  `level_expire_month` int NOT NULL DEFAULT '0' COMMENT '级别到期日期月',
  `level_expire_day` int NOT NULL DEFAULT '0' COMMENT '级别到期日期日',
  `last_login_at` int NOT NULL DEFAULT '0' COMMENT '最后一次登录时间',
  `is_ban` tinyint NOT NULL DEFAULT '0' COMMENT '是否被禁用',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `level_id` (`level_id`),
  KEY `level_expire_date` (`level_expire_date`),
  KEY `level_expire_year` (`level_expire_year`),
  KEY `level_expire_month` (`level_expire_month`),
  KEY `level_expire_day` (`level_expire_day`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='会员表';

-- ----------------------------
-- Table structure for c_user_info
-- ----------------------------
DROP TABLE IF EXISTS `c_user_info`;
CREATE TABLE `c_user_info` (
  `user_id` bigint NOT NULL COMMENT 'ID',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for c_user_level_flow
-- ----------------------------
DROP TABLE IF EXISTS `c_user_level_flow`;
CREATE TABLE `c_user_level_flow` (
  `id` bigint NOT NULL COMMENT 'id',
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `old_level_id` int NOT NULL COMMENT '原级别',
  `new_level_id` int NOT NULL COMMENT '新级别',
  `old_expire_date` date DEFAULT NULL COMMENT '原有效期',
  `new_expire_date` date DEFAULT NULL COMMENT '新有效期',
  `admin_name` varchar(20) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作管理员名称',
  `remark` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '变更描述',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for c_wallet
-- ----------------------------
DROP TABLE IF EXISTS `c_wallet`;
CREATE TABLE `c_wallet` (
  `user_id` bigint NOT NULL COMMENT '会员ID',
  `balance` int unsigned NOT NULL DEFAULT '0' COMMENT '余额',
  `gpt3` int unsigned NOT NULL DEFAULT '0' COMMENT 'gpt3提问次数',
  `gpt4` int unsigned NOT NULL DEFAULT '0' COMMENT 'gpt4提问次数',
  `midjourney` int unsigned NOT NULL DEFAULT '0' COMMENT 'midjourney提问次数',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='会员钱包';

-- ----------------------------
-- Table structure for c_wallet_flow_balance
-- ----------------------------
DROP TABLE IF EXISTS `c_wallet_flow_balance`;
CREATE TABLE `c_wallet_flow_balance` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `amount` int NOT NULL DEFAULT '0' COMMENT '变动金额',
  `total` int NOT NULL DEFAULT '0' COMMENT '变动后的余额',
  `is_incr` tinyint(1) NOT NULL DEFAULT '1' COMMENT '增加减少',
  `target_type` varchar(50) NOT NULL DEFAULT '' COMMENT '目标类型',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT '目标ID',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `admin_name` varchar(255) NOT NULL DEFAULT '' COMMENT '操作管理员',
  `year` int NOT NULL DEFAULT '0' COMMENT '年',
  `month` int NOT NULL DEFAULT '0' COMMENT '月',
  `day` int NOT NULL DEFAULT '0' COMMENT '日',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `is_incr` (`is_incr`),
  KEY `year` (`year`,`month`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='余额流水表';

-- ----------------------------
-- Table structure for c_wallet_flow_gpt3
-- ----------------------------
DROP TABLE IF EXISTS `c_wallet_flow_gpt3`;
CREATE TABLE `c_wallet_flow_gpt3` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `amount` int NOT NULL DEFAULT '0' COMMENT '变动金额',
  `total` int NOT NULL DEFAULT '0' COMMENT '变动后的余额',
  `is_incr` tinyint(1) NOT NULL DEFAULT '1' COMMENT '增加减少',
  `target_type` varchar(50) NOT NULL DEFAULT '' COMMENT '目标类型',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT '目标ID',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `admin_name` varchar(255) NOT NULL DEFAULT '' COMMENT '操作管理员',
  `year` int NOT NULL DEFAULT '0' COMMENT '年',
  `month` int NOT NULL DEFAULT '0' COMMENT '月',
  `day` int NOT NULL DEFAULT '0' COMMENT '日',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `is_incr` (`is_incr`),
  KEY `year` (`year`,`month`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='token流水表';

-- ----------------------------
-- Table structure for c_wallet_flow_gpt4
-- ----------------------------
DROP TABLE IF EXISTS `c_wallet_flow_gpt4`;
CREATE TABLE `c_wallet_flow_gpt4` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `amount` int NOT NULL DEFAULT '0' COMMENT '变动金额',
  `total` int NOT NULL DEFAULT '0' COMMENT '变动后的余额',
  `is_incr` tinyint(1) NOT NULL DEFAULT '1' COMMENT '增加减少',
  `target_type` varchar(50) NOT NULL DEFAULT '' COMMENT '目标类型',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT '目标ID',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `admin_name` varchar(255) NOT NULL DEFAULT '' COMMENT '操作管理员',
  `year` int NOT NULL DEFAULT '0' COMMENT '年',
  `month` int NOT NULL DEFAULT '0' COMMENT '月',
  `day` int NOT NULL DEFAULT '0' COMMENT '日',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `is_incr` (`is_incr`),
  KEY `year` (`year`,`month`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='token流水表';

-- ----------------------------
-- Table structure for c_wallet_flow_midjourney
-- ----------------------------
DROP TABLE IF EXISTS `c_wallet_flow_midjourney`;
CREATE TABLE `c_wallet_flow_midjourney` (
  `id` bigint NOT NULL COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '会员ID',
  `amount` int NOT NULL DEFAULT '0' COMMENT '变动金额',
  `total` int NOT NULL DEFAULT '0' COMMENT '变动后的余额',
  `is_incr` tinyint(1) NOT NULL DEFAULT '1' COMMENT '增加减少',
  `target_type` varchar(50) NOT NULL DEFAULT '' COMMENT '目标类型',
  `target_id` bigint NOT NULL DEFAULT '0' COMMENT '目标ID',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `admin_name` varchar(255) NOT NULL DEFAULT '' COMMENT '操作管理员',
  `year` int NOT NULL DEFAULT '0' COMMENT '年',
  `month` int NOT NULL DEFAULT '0' COMMENT '月',
  `day` int NOT NULL DEFAULT '0' COMMENT '日',
  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `is_incr` (`is_incr`),
  KEY `year` (`year`,`month`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='token流水表';

SET FOREIGN_KEY_CHECKS = 1;

# c_config
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('shopOrderExpireIn', '订单过期时间', '秒', 1, '', '3600', 'shop', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('shopPayExpireIn', '支付过期时间', '秒', 1, '', '300', 'shop', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('emailCodeEnable', '启用邮箱验证码', '', 2, '[{\"label\":\"是\",\"value\":\"1\"},{\"label\":\"否\",\"value\":\"0\"}]', '0', 'email', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('smtpHost', 'SMTP服务器', '', 1, '', 'smtp.qiye.aliyun.com', 'email', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('smtpPort', 'SMTP服务器端口', '', 1, '', '465', 'email', 2, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('smtpEmail', 'SMTP邮箱地址', '', 1, '', 'xxx@xxx.com', 'email', 3, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('smtpEmailPassword', 'SMTP邮箱密码', '', 1, '', 'xxxxxx', 'email', 4, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyProgressSize', 'midjourney的并发任务数', '', 1, '', '2', 'midjourney', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyQueueSize', 'midjourney排队上限', '', 1, '', '10', 'midjourney', 2, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyQueueTimeout', 'midjourney队列任务超时', '秒', 1, '', '300', 'midjourney', 3, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneySaveImage', 'midjourney图片是否本地保存', '', 2, '[{\"label\":\"是\",\"value\":\"1\"},{\"label\":\"否\",\"value\":\"0\"}]', '0', 'midjourney', 4, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyShowRemoteImage', 'midjourney展示图片', '', 2, '[{\"label\":\"远程Discord的图片\",\"value\":\"1\"},{\"label\":\"本地保存的图片\",\"value\":\"0\"}]', '1', 'midjourney', 5, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('newUserAddBalance', '新会员赠送balance', '', 1, '', '0', 'user', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('newUserAddGpt3', '新会员赠送gpt3', '次', 1, '', '100', 'user', 2, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('newUserAddGpt4', '新会员赠送gpt4', '次', 1, '', '100', 'user', 3, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('newUserAddMidjourney', '新会员赠送midjourney', '次', 1, '', '100', 'user', 4, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('gpt3UseBalance', 'gpt3提问使用balance', '', 1, '', '20', 'conversation', 1, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('gpt4UseBalance', 'gpt4提问使用balance', '', 1, '', '50', 'conversation', 2, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyUseBalance', 'midjourney提问使用balance', '', 1, '', '50', 'conversation', 3, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('midjourneyDailyLimit', 'midjourney每天最多提问次数', '次', 1, '', '10', 'midjourney', 7, 0, 0);
INSERT INTO `c_config` (`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('gptSystemContent', 'gpt对话式系统身份描述', '', 1, '', 'You are an interesting and helpful assistant who can give accurate answers.', 'conversation', 3, 0, 0);

# c_config_level
INSERT INTO `c_config_level` (`id`, `level_name`, `month_gpt3`, `month_gpt4`, `month_midjourney`) VALUES (1, 'member', 0, 0, 0);
INSERT INTO `c_config_level` (`id`, `level_name`, `month_gpt3`, `month_gpt4`, `month_midjourney`) VALUES (2, 'plus', 10000, 0, 0);

# c_config_wallet
INSERT INTO `c_config_wallet` (`field`, `wallet_name`) VALUES ('balance', '现金');
INSERT INTO `c_config_wallet` (`field`, `wallet_name`) VALUES ('gpt3', 'Gpt3次数');
INSERT INTO `c_config_wallet` (`field`, `wallet_name`) VALUES ('gpt4', 'Gpt4次数');
INSERT INTO `c_config_wallet` (`field`, `wallet_name`) VALUES ('midjourney', 'Midjourney次数');

# c_config_pay
INSERT INTO `c_config_pay` (`id`, `api_name`, `params`, `pay_channel`, `frontend_description`, `backend_description`, `status`, `created_at`, `updated_at`) VALUES (1, '微免签', '[{\"param\":\"apiKey\",\"param_name\":\"通讯密钥\",\"value\":\"xxxxxxxxx\"},{\"param\":\"host\",\"param_name\":\"host\",\"value\":\"http://xxx.xxx.com\"}]', '[{\"id\":1,\"channel\":\"1\",\"channel_name\":\"微信\",\"status\":1},{\"id\":2,\"channel\":\"2\",\"channel_name\":\"支付宝\",\"status\":1}]', '', '', 1, 0, 0);

# c_admin_role
INSERT INTO `c_admin_role` (`id`, `role_name`, `remark`, `permission`, `column_permission`, `dont_del`, `create_admin`, `update_admin`, `created_at`, `updated_at`) VALUES (1, '超级管理员', '全局管理', NULL, NULL, 1, '1', '1', 0, 0);

# c_admin
INSERT INTO `c_admin` (`id`, `user_id`, `admin_name`, `real_name`, `remark`, `role_id`, `is_enable`, `auth_key`, `password_hash`, `password_reset_token`, `dont_del`, `login_nums`, `fail_nums`, `last_login_ip`, `last_login_at`, `bind_ip`, `create_admin`, `update_admin`, `created_at`, `updated_at`, `is_delete`, `deleted_at`) VALUES (1, 1650268959049523200, 'admin', '主管理员', '全局管理', 1, 1, NULL, '$2a$10$rV4HME6.88BDuUFvpUwFBO2oiF5eiKMdc9Hx4zxLrvJR4aJ02AmJi', NULL, 0, 435, 23, '127.0.0.1', 1691923845, '', '0', 'admin', 0, 1687397934, 0, 0);
