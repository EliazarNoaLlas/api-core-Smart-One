SELECT id,
       name,
       description,
       code,
       icon,
       position,
       created_at
FROM core_modules
WHERE deleted_at IS NULL
  AND IF(? IS NULL, TRUE, core_modules.code LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, core_modules.name LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY position
LIMIT ? OFFSET ?;
