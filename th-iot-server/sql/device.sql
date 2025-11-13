DROP TABLE IF EXISTS `devices`;

CREATE TABLE
    `devices` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID，从1开始',
        `device_id` VARCHAR(64) NOT NULL COMMENT '平台设备唯一标识（唯一）',
        `name` VARCHAR(128) NOT NULL COMMENT '设备名称（唯一）',
        -- 基础信息 -------------------------------------------------------------------------------------------
        `product` VARCHAR(128) DEFAULT NULL COMMENT '所属产品型号',
        `region` VARCHAR(128) DEFAULT NULL COMMENT '地区/地理位置',
        `description` TEXT DEFAULT NULL COMMENT '设备描述',
        `extra` JSON DEFAULT NULL COMMENT '扩展属性（如地理坐标、标签、自定义参数等）',
        -- 状态属性（物模型相关/实时状态）---------------------------------------------------------------------
        `status` TINYINT NOT NULL DEFAULT 2 COMMENT '设备状态(1=在线,2=离线,3=故障,4=停用)',
        `switch` TINYINT (1) NOT NULL DEFAULT 0 COMMENT '主电源开关(1=开,0=关)',
        `relay` TINYINT (1) NOT NULL DEFAULT 0 COMMENT '继电器开关(1=开,0=关)',
        `out_j9` TINYINT (1) NOT NULL DEFAULT 0 COMMENT 'J9继电器开关(1=开,0=关)',
        `signal` INT DEFAULT NULL COMMENT '信号质量(0-31) (映射 csq)',
        `temp` DOUBLE DEFAULT NULL COMMENT '实时温度(-55~125°C) (映射 temperature)',
        `warning` VARCHAR(255) DEFAULT NULL COMMENT '预警事件(断电/过流等) (映射 alarm event)',
        -- 网络与通信信息 --------------------------------------------------------------------------------------
        `cell_info` TEXT DEFAULT NULL COMMENT '基站信息 (映射 cell_info)',
        `sim_number` VARCHAR(32) DEFAULT NULL COMMENT 'SIM卡号 (映射 imsi)',
        `report_cycle` INT DEFAULT NULL COMMENT '上报周期(秒) (映射 interval)',
        `mac` TEXT DEFAULT NULL COMMENT 'MAC地址 (映射 macs)',
        `ip_address` VARCHAR(64) DEFAULT NULL COMMENT '最近在线时的设备IP',
        `firmware_version` VARCHAR(64) DEFAULT NULL COMMENT '固件版本号',
        `last_online_at` DATETIME DEFAULT NULL COMMENT '最后一次上线时间',
        -- 通用字段 -------------------------------------------------------------------------------------------
        `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        `deleted_at` DATETIME DEFAULT NULL COMMENT '软删除时间（NULL表示未删除）',
        PRIMARY KEY (`id`),
        UNIQUE KEY `uk_device_id` (`device_id`),
        UNIQUE KEY `uk_name` (`name`),
        KEY `idx_status` (`status`),
        KEY `idx_region` (`region`),
        KEY `idx_signal` (`signal`),
        KEY `idx_deleted_at` (`deleted_at`),
        KEY `idx_last_online_at` (`last_online_at`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '设备表（结合物模型与业务字段的统一设计）';