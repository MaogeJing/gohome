CREATE TABLE cost_detail(
-- 公共基础字段
 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键'
,`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
,`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
,`deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
-- 业务字段 
`remarks` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
`shopping_id` bigint(20) unsigned NOT NULL COMMENT '关联账单',
`good_id` bigint(20) unsigned NOT NULL COMMENT '关联物品',
`cost` bigint(20) unsigned NOT NULL COMMENT '消费金额(分)',
`count` bigint(20) unsigned NOT NULL COMMENT '数量',
`comment` varchar(100) DEFAULT NULL COMMENT '针对消费项的评价',
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购买记录';
