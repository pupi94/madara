CREATE TABLE `products` (
  `id` char(36) NOT NULL,
  `store_id` char(36) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `published` tinyint(1) DEFAULT 0 COMMENT '是否上架',
  `published_at` datetime DEFAULT NULL COMMENT '发布时间',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;