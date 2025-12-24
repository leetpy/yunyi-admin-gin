CREATE TABLE users (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(64) NOT NULL UNIQUE COMMENT '用户名',
    `password` VARCHAR(255) NOT NULL COMMENT '用户密码',
    `state` TINYINT DEFAULT 0 COMMENT '状态',
    `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `last_login_at` datetime(3) DEFAULT NULL COMMENT '最后登录时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    KEY `idx_users_created_at` (`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
