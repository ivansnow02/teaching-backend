CREATE DATABASE IF NOT EXISTS teaching_course;
USE teaching_course;

-- 课程表
CREATE TABLE `course` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '课程名称',
    `cover`       varchar(255) NOT NULL DEFAULT '' COMMENT '课程封面URL',
    `description` text         COMMENT '课程描述',
    `teacher_id`  bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '教师ID',
    `status`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '状态 0:未发布 1:已发布',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `ix_teacher_id` (`teacher_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='课程表';

-- 课程章节表
CREATE TABLE `course_chapter` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_id`   bigint(20) UNSIGNED NOT NULL COMMENT '课程ID',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '章节标题',
    `sort`        int(11)      NOT NULL DEFAULT '0' COMMENT '排序',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_course_id` (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='课程章节表';

-- 课件资料表
CREATE TABLE `course_material` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `chapter_id`  bigint(20) UNSIGNED NOT NULL COMMENT '章节ID',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '课件标题',
    `type`        tinyint(2)   NOT NULL DEFAULT '2' COMMENT '资料类型 2:PDF 3:文档',
    `url`         varchar(512) NOT NULL DEFAULT '' COMMENT 'OSS存储链接',
    `file_hash`   varchar(64)  NOT NULL DEFAULT '' COMMENT '文件Hash(MD5等)',
    `file_size`   bigint(20)   NOT NULL DEFAULT '0' COMMENT '文件大小(字节)',
    `ai_status`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '向量化状态 0:未处理 1:处理中 2:已完成 3:失败',
    `sort`        int(11)      NOT NULL DEFAULT '0' COMMENT '排序',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `ix_chapter_id` (`chapter_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='课件资料表';

-- 学习进度表
CREATE TABLE `study_progress` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id`     bigint(20) UNSIGNED NOT NULL COMMENT '学生ID',
    `course_id`   bigint(20) UNSIGNED NOT NULL COMMENT '课程ID',
    `chapter_id`  bigint(20) UNSIGNED NOT NULL COMMENT '章节ID',
    `material_id` bigint(20) UNSIGNED NOT NULL COMMENT '课件ID',
    `progress`    int(11)      NOT NULL DEFAULT '0' COMMENT '进度百分比 0~100',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_material` (`user_id`, `material_id`),
    KEY `ix_user_course` (`user_id`, `course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='学习进度表';
