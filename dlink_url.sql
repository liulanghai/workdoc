

USE `scan_tool`;

DROP TABLE IF EXISTS `dlink_scan_result`;

CREATE TABLE `dlink_scan_result` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` varchar(32) NOT NULL COMMENT '用户id',
  `tool_id` int(10) unsigned NOT NULL COMMENT '工具id',
  `tool_ver` varchar(32) NOT NULL COMMENT '工具版本',
  `host_ip` varchar(32) DEFAULT NULL COMMENT '工具端ip',
  `host_os` varchar(32) DEFAULT NULL COMMENT '操作系统',
  `host_mac` varchar(32) DEFAULT NULL COMMENT 'mac地址',
  `scan_file_total` int(10) unsigned NOT NULL COMMENT '扫描文件数',
  `scan_time_total` int(10) unsigned NOT NULL COMMENT '扫描总耗时',
  `find_threats` int(10) unsigned NOT NULL COMMENT '发现威胁数',
  `scan_time_start` datetime DEFAULT NULL COMMENT '开始扫描时间',
  `details` varchar(1024) DEFAULT NULL COMMENT 'json字符串，记录每个引擎当次扫描出的威胁个数',
  `state` tinyint(4) DEFAULT '0' COMMENT '扩展字段默认为0',
  `reserve` varchar(32) DEFAULT NULL COMMENT '扩展字段',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*Table structure for table `dlink_urls` */

DROP TABLE IF EXISTS `dlink_urls`;

CREATE TABLE `dlink_urls` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `scan_id` varchar(64) NOT NULL COMMENT '扫描的记录id',
  `url` varchar(128) NOT NULL COMMENT 'urls',
  `file_md5` varchar(32) DEFAULT NULL COMMENT 'url对应文件的md5',
  `state` tinyint(4) DEFAULT '0' COMMENT '扩展字段',
  `reserve` varchar(32) DEFAULT NULL COMMENT '扩展字段',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
