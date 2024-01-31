/*
Upgrade from v1.1.1 to v1.1.3.
*/
ALTER TABLE `c_config_openai` 
ADD COLUMN `api_url` varchar(1000) NOT NULL DEFAULT '' COMMENT 'api_url' AFTER `title`;