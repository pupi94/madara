CREATE TABLE `variants` (
  `id` varbinary(16) NOT NULL,
  `product_id` varbinary(16) DEFAULT NULL,
  `store_id` varbinary(16) DEFAULT NULL,
  `position` int(11) DEFAULT NULL,
  `compare_at_price` decimal(15,2) DEFAULT '0.00',
  `price` decimal(15,2) DEFAULT '0.00',
  `barcode` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_variants_on_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
