CREATE TABLE `users` (
    `id` CHAR(36) PRIMARY KEY,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX `idx_users_first_name_last_name_created_at` ON `users` (`first_name`, `last_name`, `created_at`);
