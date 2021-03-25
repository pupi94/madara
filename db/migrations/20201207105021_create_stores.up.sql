CREATE TABLE `stores` (
  `id` bigint NOT NULL auto_increment,
  `created_at` bigint NOT NULL,
  `updated_at` bigint NOT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;