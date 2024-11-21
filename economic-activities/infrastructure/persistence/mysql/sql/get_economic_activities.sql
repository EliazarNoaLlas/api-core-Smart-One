SELECT id,
       cuui_id,
       description,
       status,
       created_at
FROM core_economic_activities
WHERE deleted_at IS NULL
  AND IF(? IS NULL, TRUE, cuui_id = TRIM(?))
  AND IF(? IS NULL, TRUE, description LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY created_at DESC
LIMIT ? OFFSET ?;
