CREATE DATABASE IF NOT EXISTS `kubikvest`
  CHARACTER SET utf8
  COLLATE utf8_general_ci;

USE kubikvest;

DROP TABLE IF EXISTS `kv_user`;

CREATE TABLE `kv_user` (
  `user_id`      CHAR(36)            NOT NULL,
  `provider`     CHAR(2)          DEFAULT NULL,
  `uid`          BIGINT(20) UNSIGNED NOT NULL,
  `access_token` VARCHAR(255)     DEFAULT NULL,
  `group_id`     CHAR(36)         DEFAULT NULL,
  `ttl`          INT(11) UNSIGNED DEFAULT NULL,
  `quest_id`     CHAR(36)         DEFAULT NULL,
  `point_id`     CHAR(36)         DEFAULT NULL,
  `log_quest`    BLOB,
  `start_task`   DATETIME
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `kv_group`;

CREATE TABLE `kv_group` (
  `groupId`    CHAR(36) NOT NULL,
  `gameId`     CHAR(36)             DEFAULT NULL,
  `questId`    CHAR(36)             DEFAULT NULL,
  `pointId`    CHAR(36)             DEFAULT NULL,
  `users`      BLOB,
  `pin`        SMALLINT(4) UNSIGNED DEFAULT NULL,
  `startPoint` DATETIME,
  `active`     TINYINT(2) UNSIGNED  DEFAULT 1
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT kv_user (user_id, provider, uid, access_token)
VALUES ('adff5c92-008c-47ac-bad8-11be43ea1469', 'vk', 1111, 'token');
