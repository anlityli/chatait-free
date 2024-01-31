/*
Upgrade from v1.1.3 to v1.2.0
*/
ALTER TABLE `c_queue_midjourney` 
ADD COLUMN `application_type` int NOT NULL DEFAULT 1 COMMENT '应用机器人类型 1MJ 2Niji' AFTER `action_type`;