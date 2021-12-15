CREATE TABLE IF NOT EXISTS `t_email_config` (
    `id` BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(50) UNIQUE NOT NULL COMMENT 'アカウント名',
    `password` VARCHAR(50) NOT NULL COMMENT 'パスワード',
    `smtp_server` VARCHAR(50) NOT NULL COMMENT 'SMTP送信ホストサーバー',
    `smtp_port` INT NOT NULL COMMENT 'SMTP送信ポート',
    `sender` VARCHAR(50) NULL COMMENT '差出人のメールアドレス',
    `display_name` VARCHAR(50) NULL COMMENT '差出人の表示名',
    `is_default` TINYINT(1) DEFAULT 1 NOT NULL COMMENT '既定アカウント'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;