//数据库：thiotdb
//用户表：users
CREATE DATABASE IF NOT EXISTS thiotdb DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
USE thiotdb;

CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID，自增主键',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名，唯一标识用户',
    password VARCHAR(255) NOT NULL COMMENT '用户密码（加密存储）',
    email VARCHAR(100) DEFAULT NULL COMMENT '用户邮箱',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

ALTER TABLE users
ADD COLUMN DeletedAt DATETIME NULL COMMENT '逻辑删除标志，软删除' AFTER updated_at,
ADD INDEX idx_deleted_at (DeletedAt);
-- fix
DeletedAt   datetime  YES  MUL  NULL
deleted_at  datetime  YES  NULL

-- 创建数据库
CREATE DATABASE IF NOT EXISTS thiotdb 
DEFAULT CHARSET = utf8mb4 
COLLATE = utf8mb4_general_ci;

-- 使用数据库
USE thiotdb;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID，自增主键',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名，唯一标识用户',
    password VARCHAR(255) NOT NULL COMMENT '用户密码（加密存储）',
    email VARCHAR(100) DEFAULT NULL COMMENT '用户邮箱',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '逻辑删除标志，软删除',
    PRIMARY KEY (id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
