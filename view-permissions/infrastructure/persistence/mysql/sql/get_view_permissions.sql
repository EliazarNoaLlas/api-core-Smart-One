SELECT view_permissions.id         AS view_permission_id,
       view_permissions.created_by AS view_permission_created_by,
       view_permissions.created_at AS view_permission_created_at,
       views.id                    AS view_id,
       views.name                  AS view_name,
       views.description           AS view_description,
       views.created_at            AS view_created_at,
       permissions.id              AS permission_id,
       permissions.code            AS permission_code,
       permissions.name            AS permission_name,
       permissions.description     AS permission_description,
       permissions.created_at      AS permission_created_at,
       modules.id                  AS module_id,
       modules.name                AS module_name,
       modules.description         AS module_description,
       modules.code                AS module_code,
       modules.icon                AS module_icon,
       modules.position            AS module_position,
       modules.created_at          AS module_created_at
FROM core_view_permissions view_permissions
         INNER JOIN core_views views ON view_permissions.view_id = views.id
         INNER JOIN core_permissions permissions ON view_permissions.permission_id = permissions.id
         INNER JOIN core_modules modules ON permissions.module_id = modules.id
WHERE view_permissions.deleted_at IS NULL
  AND view_permissions.view_id = ?
  AND views.deleted_at IS NULL
  AND permissions.deleted_at IS NULL
  AND modules.deleted_at IS NULL
ORDER BY view_permissions.created_at DESC;
