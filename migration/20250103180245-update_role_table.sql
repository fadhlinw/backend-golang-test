
-- +migrate Up

INSERT INTO `role` (`id`, `title`, `slug`, `description`, `active`, `created_at`, `updated_at`) VALUES
(1, 'Owner', 'owner', 'owner', 1, '2025-01-03 18:02:45', '2025-01-03 18:02:45'),
(2, 'Librarian', 'librarian', 'librarian', 1, '2025-01-03 18:02:45', '2025-01-03 18:02:45'),
(3, 'Visitor', 'visitor', 'visitor', 1, '2025-01-03 18:02:45', '2025-01-03 18:02:45');

-- +migrate Down

DELETE FROM `role`;