CREATE TABLE IF NOT EXISTS `t_email_log` (
   `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
   `action_time` datetime(6) NOT NULL,
   `username` varchar(50) DEFAULT NULL COMMENT '送信ユーザー名',
   `sender` varchar(200) NOT NULL,
   `recipient` varchar(200) NOT NULL,
   `cc` varchar(200) DEFAULT NULL,
   `bcc` varchar(200) DEFAULT NULL,
   `title` varchar(50) NOT NULL,
   `body` longtext NOT NULL,
   `password_body` longtext,
   `attachments` varchar(2000) DEFAULT NULL,
   `server_name` varchar(100) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;