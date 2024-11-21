SELECT COUNT(*) AS total
FROM core_store_types
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

