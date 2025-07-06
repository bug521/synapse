-- Synapse 数据库初始化脚本

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS synapse CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE synapse;

-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    password VARCHAR(128) NOT NULL COMMENT '密码',
    email VARCHAR(100) UNIQUE COMMENT '邮箱',
    active BOOLEAN DEFAULT TRUE COMMENT '是否激活',
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_username (username),
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 通道表
CREATE TABLE IF NOT EXISTS channels (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '通道ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '所属用户的ID',
    name VARCHAR(255) NOT NULL COMMENT '通道名称',
    type VARCHAR(50) NOT NULL COMMENT '通道类型',
    credentials JSON NOT NULL COMMENT '凭证',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_type (type),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='通知通道表';

-- 主题表
CREATE TABLE IF NOT EXISTS topics (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主题ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '所属用户的ID',
    name VARCHAR(255) NOT NULL COMMENT '主题名称',
    webhook_key VARCHAR(36) NOT NULL UNIQUE COMMENT 'Webhook Key',
    sending_strategy VARCHAR(50) DEFAULT 'all' COMMENT '发送策略',
    execution_mode VARCHAR(50) DEFAULT 'async' COMMENT '执行模式',
    description TEXT COMMENT '主题描述',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_webhook_key (webhook_key),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息主题表';

-- 消息表
CREATE TABLE IF NOT EXISTS messages (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '消息ID',
    topic_id BIGINT UNSIGNED NOT NULL COMMENT '来源主题ID',
    content JSON NOT NULL COMMENT '原始消息内容',
    status VARCHAR(50) DEFAULT 'pending' COMMENT '处理状态',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    INDEX idx_topic_id (topic_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息表';

-- 路由表
CREATE TABLE IF NOT EXISTS routings (
    topic_id BIGINT UNSIGNED NOT NULL COMMENT '主题ID',
    channel_id BIGINT UNSIGNED NOT NULL COMMENT '通道ID',
    priority INT DEFAULT 0 COMMENT '优先级',
    variable_mappings JSON COMMENT '变量映射规则',
    message_template TEXT COMMENT '消息模板',
    subject_template TEXT COMMENT '邮件主题模板',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    PRIMARY KEY (topic_id, channel_id),
    INDEX idx_topic_id (topic_id),
    INDEX idx_channel_id (channel_id),
    INDEX idx_priority (priority),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='路由表';

-- 投递日志表
CREATE TABLE IF NOT EXISTS message_delivery_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID',
    message_id BIGINT UNSIGNED NOT NULL COMMENT '消息ID',
    channel_id BIGINT UNSIGNED NOT NULL COMMENT '目标通道ID',
    status VARCHAR(50) NOT NULL COMMENT '投递状态',
    response TEXT COMMENT 'API响应',
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新时间',
    deleted_at DATETIME(3) NULL COMMENT '删除时间',
    INDEX idx_message_id (message_id),
    INDEX idx_channel_id (channel_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='投递日志表';

-- 插入默认数据（可选）
-- INSERT INTO users (username, password, email) VALUES ('admin', 'hashed_password', 'admin@example.com'); 