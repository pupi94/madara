CREATE TABLE `products` (
  `id` varbinary(16) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `published` tinyint(1) DEFAULT '0' COMMENT '是否上架',
  `published_at` datetime DEFAULT NULL COMMENT '发布时间',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `spu` varchar(255) DEFAULT NULL,
  `tags` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;