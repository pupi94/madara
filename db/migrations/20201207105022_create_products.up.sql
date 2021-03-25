CREATE TABLE `products` (
  `id` bigint NOT NULL auto_increment,
  `store_id` bigint NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `published` tinyint(1) DEFAULT 0 COMMENT '是否上架',
  `published_at` bigint DEFAULT NULL COMMENT '发布时间',
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
   PRIMARY KEY (`id`),
   KEY `index_products_on_store_id` (`store_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;