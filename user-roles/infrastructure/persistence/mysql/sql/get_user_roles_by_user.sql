SELECT user_roles.id         AS user_role_id,
       user_roles.enable     AS user_role_enable,
       user_roles.created_at AS user_role_created_at,
       roles.id              AS role_id,
       roles.name            AS role_name,
       roles.description     AS role_description,
       roles.enable          AS role_enable,
       roles.created_at      AS role_created_at
FROM core_user_roles user_roles
         LEFT JOIN core_roles roles ON roles.id = user_roles.role_id
WHERE user_roles.user_id = ?
  AND roles.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
ORDER BY user_roles.created_at DESC LIMIT ?
OFFSET ?;