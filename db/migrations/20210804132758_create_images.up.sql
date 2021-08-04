CREATE TABLE `images` (
  `id` varbinary(16) NOT NULL,
  `product_id` varbinary(16) DEFAULT NULL,
  `store_id` varbinary(16) DEFAULT NULL,
  `position` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `properties` text,
  PRIMARY KEY (`id`),
  KEY `index_images_on_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;