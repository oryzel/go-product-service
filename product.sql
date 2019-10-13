-- -----------------------------------------------------
-- Schema product_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `product_db` DEFAULT CHARACTER SET utf8 ;

USE `product_db` ;

-- -----------------------------------------------------
-- table products
-- -----------------------------------------------------
CREATE TABLE `products` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(191) NOT NULL,
    `sku` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
    `type` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
    `category` int(10) DEFAULT 1,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;