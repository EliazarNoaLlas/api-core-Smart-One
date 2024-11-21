SELECT COUNT(*) AS total
FROM core_permissions permissions
         INNER JOIN core_modules modules ON permissions.module_id = modules.id
WHERE permissions.module_id = ?
  AND IF(? IS NULL, TRUE, permissions.code LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, permissions.name LIKE CONCAT('%', TRIM(?), '%'))
  AND permissions.deleted_at IS NULL
ORDER BY permissions.created_at DESC;
