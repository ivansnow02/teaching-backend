CREATE DATABASE IF NOT EXISTS teaching_user;
USE teaching_user;

CREATE TABLE `user` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `email`       varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
    `password`    varchar(128) NOT NULL DEFAULT '' COMMENT '密码(加密)',
    `nickname`    varchar(64)  NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`      varchar(255) NOT NULL DEFAULT '' COMMENT '头像URL',
    `role`        tinyint(4)   NOT NULL DEFAULT '1' COMMENT '角色 1:学生 2:教师 3:管理员',
    `status`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '状态 0:正常 1:禁用',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
