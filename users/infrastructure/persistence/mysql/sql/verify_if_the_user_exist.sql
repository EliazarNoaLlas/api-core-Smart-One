SELECT COUNT(*) AS total
FROM core_users
WHERE id = ?
  AND deleted_at IS NULL;