SELECT COUNT(*)
FROM core_user_roles user_roles
         LEFT JOIN core_roles roles ON roles.id = user_roles.role_id
WHERE user_roles.user_id = ?
  AND roles.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL;
