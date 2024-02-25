/*
Upgrade from v1.1.3 to v1.2.0
*/
ALTER TABLE `c_queue_midjourney` 
ADD COLUMN `application_type` int NOT NULL DEFAULT 1 COMMENT '应用机器人类型 1MJ 2Niji' AFTER `action_type`;
ALTER TABLE `c_config_baidu` 
ADD COLUMN `features` varchar(255) NOT NULL DEFAULT '[]' COMMENT '可用功能数组json' AFTER `status`;

UPDATE `c_config` SET `title`='midjourney整站并发任务数' WHERE `config_name`='midjourneyProgressSize';
UPDATE `c_config` SET `title`='midjourney整站排队上限' WHERE `config_name`='midjourneyQueueSize';
INSERT INTO `c_config` (config_name,title,input_type,value,type,sort) VALUES ('midjourneyUserProgressSize','midjourney会员并发任务数',1,3,'midjourney',6);