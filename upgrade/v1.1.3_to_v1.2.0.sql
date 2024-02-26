/*
Upgrade from v1.1.3 to v1.2.0
*/
ALTER TABLE `c_queue_midjourney` 
ADD COLUMN `application_type` int NOT NULL DEFAULT 1 COMMENT '应用机器人类型 1MJ 2Niji' AFTER `action_type`;

ALTER TABLE `c_config_baidu` 
ADD COLUMN `features` varchar(255) NOT NULL DEFAULT '[]' COMMENT '可用功能数组json' AFTER `status`;

ALTER TABLE `c_config_openai` 
ADD COLUMN `gpt3_model` varchar(255) NOT NULL DEFAULT '' COMMENT 'gpt3使用模型' AFTER `max_tokens`,
ADD COLUMN `gpt4_model` varchar(255) NOT NULL DEFAULT '' COMMENT 'gpt4使用模型' AFTER `gpt3_model`;

CREATE TABLE `c_config_sensitive_word` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '敏感词内容',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `content` (`content`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='敏感词';

CREATE TABLE `c_user_sensitive_word` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '会员ID',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '类型: 1对话 2用户名',
  `topic_type` int(11) NOT NULL DEFAULT '0' COMMENT '对话类型: 0无 1gpt3 2gpt4 3mj',
  `content` text COLLATE utf8mb4_general_ci COMMENT '触发敏感词原文',
  `validate_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '敏感词json',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='敏感词触发记录';

UPDATE `c_config` SET `title`='midjourney整站并发任务数' WHERE `config_name`='midjourneyProgressSize';
UPDATE `c_config` SET `title`='midjourney整站排队上限' WHERE `config_name`='midjourneyQueueSize';
INSERT INTO `c_config` (config_name,title,input_type,value,type,sort) VALUES ('midjourneyUserProgressSize','midjourney会员并发任务数',1,3,'midjourney',6);