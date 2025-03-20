
-- im.users definition

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `mobile` varchar(20) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `passwd` varchar(255) DEFAULT NULL,
  `salt` blob,
  `nickname` varchar(100) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `sex` tinyint DEFAULT NULL,
  `memo` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- im.friendships definition

CREATE TABLE `friendships` (
  `user_id` int NOT NULL COMMENT '用户ID',
  `friend_id` int NOT NULL COMMENT '好友ID',
  `comment` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`,`friend_id`),
  KEY `friend_id` (`friend_id`),
  CONSTRAINT `friendships_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `friendships_ibfk_2` FOREIGN KEY (`friend_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='好友关系表';

-- im.friend_apply definition

CREATE TABLE `friend_apply` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '申请ID',
  `from_id` int NOT NULL COMMENT '申请者ID',
  `to_id` int NOT NULL COMMENT '被申请者ID',
  `status` int DEFAULT '0' COMMENT '申请状态',
  `apply_reason` varchar(255) DEFAULT NULL COMMENT '申请理由',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `from_id` (`from_id`,`to_id`) COMMENT '避免重复申请',
  KEY `to_id` (`to_id`),
  CONSTRAINT `friend_apply_ibfk_1` FOREIGN KEY (`from_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `friend_apply_ibfk_2` FOREIGN KEY (`to_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='好友申请表';


-- im.`group` definition

CREATE TABLE `group` (
  `group_id` bigint NOT NULL AUTO_INCREMENT COMMENT '群聊ID，主键，自增',
  `group_name` varchar(255) NOT NULL COMMENT '群聊名称',
  `creator_id` int NOT NULL COMMENT '创建者用户ID',
  `avatar_url` varchar(255) DEFAULT NULL COMMENT '群聊头像URL',
  `description` text COMMENT '群聊描述',
  `max_members` int DEFAULT '200' COMMENT '群聊最大成员数',
  `status` tinyint DEFAULT '0' COMMENT '群聊状态（0:正常，1:解散）',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`group_id`),
  KEY `fk_group_member_group_1` (`creator_id`),
  CONSTRAINT `fk_group_member_group_1` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='群聊表';

-- im.group_member definition

CREATE TABLE `group_member` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键，自增',
  `group_id` bigint NOT NULL COMMENT '群聊ID，外键',
  `user_id` int NOT NULL COMMENT '用户ID',
  `join_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
  `role` tinyint DEFAULT '0' COMMENT '成员角色（0:普通成员，1:管理员，2:群主）',
  `status` tinyint DEFAULT '0' COMMENT '成员状态（0:正常，1:已退出）',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_group_user` (`group_id`,`user_id`),
  KEY `fk_group_member_ibfk_1` (`user_id`),
  CONSTRAINT `fk_group_member_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_group_member_ibfk_2` FOREIGN KEY (`group_id`) REFERENCES `group` (`group_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='群聊成员表';


-- im.group_application definition

CREATE TABLE `group_apply` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '申请ID，主键，自增',
  `group_id` bigint NOT NULL COMMENT '群聊ID，外键',
  `applicant_id` int NOT NULL COMMENT '申请人用户ID',
  `status` tinyint DEFAULT '0' COMMENT '申请状态（0:待处理，1:已通过，2:已拒绝）',
  `handler_id` int DEFAULT NULL COMMENT '处理人用户ID',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_applicant_id` (`applicant_id`),
  KEY `fk_group_application_ibfk_1` (`handler_id`),
  CONSTRAINT `fk_group_application_group` FOREIGN KEY (`group_id`) REFERENCES `group` (`group_id`) ON DELETE CASCADE,
  CONSTRAINT `fk_group_application_ibfk_1` FOREIGN KEY (`handler_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_group_application_ibfk_2` FOREIGN KEY (`applicant_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='群聊申请表';