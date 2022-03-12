-- 用户表 --
CREATE TABLE `user`
(
    `id`          bigint(10) NOT NULL AUTO_INCREMENT,
    `user`        varchar(50)                             NOT NULL DEFAULT '' COMMENT '用户',
    `name`        varchar(255) COLLATE utf8mb4_general_ci NULL COMMENT '用户\t名称',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户\n密码',
    `mobile`      varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `gender`      char(5) COLLATE utf8mb4_general_ci      NOT NULL COMMENT '男｜女｜未公\r开',
    `nickname`    varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT '用户昵称',
    `type`    tinyint(1) COLLATE utf8mb4_general_ci          DEFAULT 0 COMMENT '用户类型',
    `create_time` timestamp NULL,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_index` (`name`),
    UNIQUE KEY `name_index2` (`name`),
    UNIQUE KEY `user_index` (`user`),
    UNIQUE KEY `type_index` (`type`),
    UNIQUE KEY `mobile_index` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `student`
(
    `id`    bigint                           NOT NULL,
    `class` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `name`  varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `age`   tinyint                                   DEFAULT NULL,
    `score` float(10, 0
) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `class_name_index` (`class`,`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;