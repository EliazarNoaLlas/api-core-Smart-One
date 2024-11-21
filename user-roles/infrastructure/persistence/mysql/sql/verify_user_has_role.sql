SELECT COUNT(*) AS total
FROM core_user_roles user_roles
WHERE user_roles.deleted_at IS NULL
  AND user_roles.user_id = ?
  AND user_roles.role_id = ?;
