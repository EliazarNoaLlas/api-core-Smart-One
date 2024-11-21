SELECT COUNT(*) AS total
FROM core_views views
WHERE views.deleted_at IS NULL
  AND views.module_id = ?
  AND IF(? IS NULL, TRUE, name LIKE CONCAT('%', TRIM(?), '%')
    OR description LIKE CONCAT('%', TRIM(?), '%'));
