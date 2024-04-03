/*
Upgrade from v1.2.0 to v1.2.1
*/

INSERT INTO `c_config`(`config_name`, `title`, `unit`, `input_type`, `options`, `value`, `type`, `sort`, `created_at`, `updated_at`) VALUES ('speakContentMaxLength', '提问最大长度', '字符', 1, '', '500', 'conversation', 5, 0, 0);