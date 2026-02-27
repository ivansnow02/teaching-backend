CREATE DATABASE IF NOT EXISTS teaching_exam;
USE teaching_exam;

-- 题库表
CREATE TABLE `question` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_id`   bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '归属课程ID',
    `type`        tinyint(2)   NOT NULL DEFAULT '1' COMMENT '题型 1:单选 2:多选 3:判断 4:填空 5:简答/主观题',
    `content`     text         NOT NULL COMMENT '题目内容(JSON,含选项)',
    `answer`      text         NOT NULL COMMENT '标准答案',
    `analysis`    text         COMMENT '题目解析',
    `knowledge_points` varchar(255) NOT NULL DEFAULT '' COMMENT '知识点(关联RAG)',
    `score`       decimal(5,1) NOT NULL DEFAULT '0.0' COMMENT '默认分值',
    `difficulty`  tinyint(2)   NOT NULL DEFAULT '1' COMMENT '难度 1:简单 2:中等 3:困难',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `ix_course_id` (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='题库表';

-- 试卷表
CREATE TABLE `exam` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `course_id`   bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '关联课程ID',
    `title`       varchar(255) NOT NULL DEFAULT '' COMMENT '试卷名称',
    `total_score` decimal(5,1) NOT NULL DEFAULT '100.0' COMMENT '总分',
    `pass_score`  decimal(5,1) NOT NULL DEFAULT '60.0' COMMENT '及格分',
    `duration`    int(11)      NOT NULL DEFAULT '120' COMMENT '考试时长(分钟)',
    `start_time`  timestamp    NULL COMMENT '考试开始时间',
    `end_time`    timestamp    NULL COMMENT '考试结束时间',
    `exam_type`   tinyint(2)   NOT NULL DEFAULT '1' COMMENT '组卷模式 1:固定 2:随机',
    `rule_json`   text         COMMENT '随机组卷规则JSON',
    `status`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '状态 0:未开始 1:进行中 2:已结束',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `ix_course_id` (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='试卷表';

-- 试卷题目关联表
CREATE TABLE `exam_question` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `exam_id`     bigint(20) UNSIGNED NOT NULL COMMENT '试卷ID',
    `question_id` bigint(20) UNSIGNED NOT NULL COMMENT '题目ID',
    `score`       decimal(5,1) NOT NULL DEFAULT '0.0' COMMENT '本试卷中该题分值',
    `sort`        int(11)      NOT NULL DEFAULT '0' COMMENT '题号排序',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_exam_question` (`exam_id`, `question_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='试卷题目关系表';

-- 学生考试记录表
CREATE TABLE `user_exam_record` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `exam_id`     bigint(20) UNSIGNED NOT NULL COMMENT '试卷ID',
    `user_id`     bigint(20) UNSIGNED NOT NULL COMMENT '学生ID',
    `score`       decimal(5,1) NOT NULL DEFAULT '0.0' COMMENT '最终得分',
    `status`      tinyint(4)   NOT NULL DEFAULT '0' COMMENT '状态 0:答题中 1:已交卷待批改 2:已批改',
    `submit_time` timestamp    NULL COMMENT '交卷时间',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_exam_user` (`exam_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='学生考试记录表';

-- 学生答题明细表
CREATE TABLE `user_answer` (
    `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `record_id`   bigint(20) UNSIGNED NOT NULL COMMENT '答卷记录ID',
    `question_id` bigint(20) UNSIGNED NOT NULL COMMENT '题目ID',
    `user_answer` text         COMMENT '用户提交的答案',
    `is_correct`  tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否正确 0:错 1:对 2:半对/部分得分',
    `score`       decimal(5,1) NOT NULL DEFAULT '0.0' COMMENT '此题得分',
    `ai_status`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT 'AI批改状态 0:未批改 1:评卷中 2:已完成',
    `ai_comment`  text         COMMENT 'AI评语和解析',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  timestamp    NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_record_question` (`record_id`, `question_id`),
    KEY `ix_record_id` (`record_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='考生答题明细表';
