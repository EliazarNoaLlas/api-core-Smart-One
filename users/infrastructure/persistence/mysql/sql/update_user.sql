UPDATE core_users
SET username = TRIM(?),
    type_id  = ?
WHERE id = ?
  AND deleted_at IS NULL;
