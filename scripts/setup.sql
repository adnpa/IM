CREATE DATABASE `goim` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- goim.chat_log definition

CREATE TABLE `chat_log` (
                            `msg_id` varchar(128) NOT NULL,
                            `send_id` varchar(255) NOT NULL,
                            `session_type` int NOT NULL,
                            `recv_id` varchar(255) NOT NULL,
                            `content_type` int NOT NULL,
                            `msg_from` int NOT NULL,
                            `content` varchar(1000) NOT NULL,
                            `remark` varchar(100) DEFAULT NULL,
                            `sender_platform_id` int NOT NULL,
                            `send_time` datetime NOT NULL,
                            PRIMARY KEY (`msg_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- goim.friend definition

CREATE TABLE `friend` (
                          `owner_id` varchar(255) NOT NULL,
                          `friend_id` varchar(255) NOT NULL,
                          `comment` varchar(255) DEFAULT NULL,
                          `friend_flag` int NOT NULL,
                          `create_time` datetime NOT NULL,
                          PRIMARY KEY (`owner_id`,`friend_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- goim.friend_request definition

CREATE TABLE `friend_request` (
                                  `req_id` varchar(255) NOT NULL,
                                  `user_id` varchar(255) NOT NULL,
                                  `flag` int NOT NULL DEFAULT '0',
                                  `req_message` varchar(255) DEFAULT NULL,
                                  `create_time` datetime NOT NULL,
                                  PRIMARY KEY (`user_id`,`req_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- goim.`group` definition

CREATE TABLE `group` (
                         `group_id` varchar(64) NOT NULL,
                         `name` varchar(255) DEFAULT NULL,
                         `introduction` varchar(255) DEFAULT NULL,
                         `notification` varchar(255) DEFAULT NULL,
                         `face_url` varchar(255) DEFAULT NULL,
                         `create_time` datetime DEFAULT NULL,
                         `ex` varchar(255) DEFAULT NULL,
                         PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- goim.group_member definition

CREATE TABLE `group_member` (
                                `group_id` varchar(64) NOT NULL,
                                `uid` varchar(64) NOT NULL,
                                `nickname` varchar(255) DEFAULT NULL,
                                `user_group_face_url` varchar(255) DEFAULT NULL,
                                `administrator_level` int NOT NULL,
                                `join_time` datetime NOT NULL,
                                PRIMARY KEY (`group_id`,`uid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


-- goim.group_request definition

CREATE TABLE `group_request` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `group_id` varchar(64) NOT NULL,
                                 `from_user_id` varchar(255) NOT NULL,
                                 `to_user_id` varchar(255) NOT NULL,
                                 `flag` int NOT NULL DEFAULT '0',
                                 `req_msg` varchar(255) DEFAULT '',
                                 `handled_msg` varchar(255) DEFAULT '',
                                 `create_time` datetime NOT NULL,
                                 `from_user_nickname` varchar(255) DEFAULT '',
                                 `to_user_nickname` varchar(255) DEFAULT NULL,
                                 `from_user_face_url` varchar(255) DEFAULT '',
                                 `to_user_face_url` varchar(255) DEFAULT '',
                                 `handled_user` varchar(255) DEFAULT '',
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- goim.register definition

CREATE TABLE `register` (
                            `account` varchar(255) NOT NULL,
                            `password` varchar(255) NOT NULL,
                            PRIMARY KEY (`account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;


-- goim.`user` definition

CREATE TABLE `user` (
                        `uid` varchar(64) NOT NULL,
                        `name` varchar(64) DEFAULT NULL,
                        `icon` varchar(1024) DEFAULT NULL,
                        `gender` int(11) unsigned zerofill DEFAULT NULL,
                        `mobile` varchar(32) DEFAULT NULL,
                        `birth` varchar(16) DEFAULT NULL,
                        `email` varchar(64) DEFAULT NULL,
                        `ex` varchar(1024) DEFAULT NULL,
                        `create_time` datetime DEFAULT NULL,
                        PRIMARY KEY (`uid`),
                        UNIQUE KEY `uk_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- goim.user_black_list definition

CREATE TABLE `user_black_list` (
                                   `owner_id` varchar(64) NOT NULL,
                                   `block_id` varchar(64) NOT NULL,
                                   `create_time` datetime NOT NULL,
                                   PRIMARY KEY (`owner_id`,`block_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;