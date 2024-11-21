SELECT COUNT(*) as total
FROM core_document_types
WHERE deleted_at IS NULL
  AND IF(? IS NULL, TRUE, number LIKE CONCAT('%', TRIM(?), '%') OR
                          description LIKE CONCAT('%', TRIM(?), '%') OR
                          abbreviated_description LIKE CONCAT('%', TRIM(?), '%'));
