CREATE TABLE IF NOT EXISTS `users` (
     `id` bigint(20) NOT NULL,
     `name` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
     `created_at` datetime(3) DEFAULT NULL,
     `updated_at` datetime(3) DEFAULT NULL,
     `deleted_at` datetime(3) DEFAULT NULL,
     `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
     `birth_date` datetime(3) DEFAULT NULL,
     `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;