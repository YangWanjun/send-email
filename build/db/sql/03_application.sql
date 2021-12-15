CREATE TABLE IF NOT EXISTS `t_application` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `client_id` varchar(100) UNIQUE NOT NULL COMMENT 'クライアントID',
    `client_secret` varchar(255) NOT NULL COMMENT 'クライアントシークレット',
    `domain` varchar(100) NOT NULL COMMENT 'ドメイン',
    `name` varchar(50) DEFAULT NULL COMMENT '名称',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;