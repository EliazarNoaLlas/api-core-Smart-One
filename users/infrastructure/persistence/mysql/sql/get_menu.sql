SELECT modules.id            AS module_id,
       modules.name          AS module_name,
       modules.description   AS module_description,
       modules.code          AS module_code,
       modules.icon          AS module_icon,
       MAX(modules.position) AS module_position,
       modules.created_at    AS module_created_at,
       views.id              AS view_id,
       views.name            AS view_name,
       views.description     AS view_description,
       views.url             AS view_url,
       views.icon            AS view_icon,
       views.created_at      AS view_created_at
FROM core_users users
         INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id
         INNER JOIN core_roles roles ON user_roles.role_id = roles.id
         INNER JOIN core_role_policies role_policies ON roles.id = role_policies.role_id
         INNER JOIN core_policies policies ON role_policies.policy_id = policies.id
         INNER JOIN core_policy_permissions policy_permissions ON policies.id = policy_permissions.policy_id
         INNER JOIN core_permissions permissions ON policy_permissions.permission_id = permissions.id
         INNER JOIN core_view_permissions ON permissions.id = core_view_permissions.permission_id
         INNER JOIN core_modules modules ON permissions.module_id = modules.id
         INNER JOIN core_views views ON core_view_permissions.view_id = views.id
WHERE users.id = ?
  AND users.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND roles.deleted_at IS NULL
  AND role_policies.deleted_at IS NULL
  AND policies.deleted_at IS NULL
  AND policy_permissions.deleted_at IS NULL
  AND permissions.deleted_at IS NULL
  AND modules.deleted_at IS NULL
  AND views.deleted_at IS NULL
  AND core_view_permissions.deleted_at IS NULL
GROUP BY modules.id, views.id
ORDER BY module_position;
