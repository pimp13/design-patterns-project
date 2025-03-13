CREATE TABLE `categories` (
    `id` CHAR(36) PRIMARY KEY,
    `name` VARCHAR(160) NOT NULL,
    `description` TEXT,
    `slug` VARCHAR(160) NOT NULL UNIQUE,
    `image` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX `idx_categories_name` ON `categories` (`name`);
CREATE INDEX `idx_categories_created_at` ON `categories` (`created_at`);