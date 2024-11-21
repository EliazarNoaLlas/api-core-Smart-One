SELECT permissions.id          AS permission_id,
       permissions.code        AS permission_code,
       permissions.name        AS permission_name,
       permissions.description AS permission_description,
       permissions.created_at  AS permission_created_at,
       modules.id              AS module_id,
       modules.name            AS module_name,
       modules.description     AS module_description,
       modules.code            AS module_code
FROM core_permissions permissions
         INNER JOIN core_modules modules ON permissions.module_id = modules.id
WHERE permissions.module_id = ?
  AND permissions.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, permissions.code LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, permissions.name LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY permissions.created_at DESC
LIMIT ? OFFSET ?;
