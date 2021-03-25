CREATE TABLE `inventories` (
  `id` bigint NOT NULL auto_increment,
  `store_id` bigint NOT NULL,
  `product_id` bigint NOT NULL,
  `value` bigint DEFAULT 0,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
   PRIMARY KEY (`id`),
   KEY `index_inventories_on_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;