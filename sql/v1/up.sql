CREATE SCHEMA `qt-business` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_german2_ci ;
CREATE TABLE `Users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id自增',
  `uid` varchar(32) NOT NULL COMMENT '用户唯一id',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '手机号码',
  `deleted_at` int(1) DEFAULT '0' COMMENT '1 为已经删除, 0 为未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid_UNIQUE` (`uid`),
  UNIQUE KEY `phone_UNIQUE` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4