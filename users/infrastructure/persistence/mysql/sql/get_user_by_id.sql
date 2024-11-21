SELECT id,
       username,
       created_at
FROM core_users
WHERE id = ?
  AND deleted_at IS NULL;