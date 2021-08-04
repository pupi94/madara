CREATE TABLE `inventories` (
  `id` bigint NOT NULL auto_increment,
  `store_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  `source_type` varchar(255) NOT NULL,
  `value` bigint DEFAULT 0,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
   PRIMARY KEY (`id`),
   KEY `index_inventories_on_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
