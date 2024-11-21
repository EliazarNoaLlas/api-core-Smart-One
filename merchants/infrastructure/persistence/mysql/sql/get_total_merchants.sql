SELECT COUNT(*) AS total
FROM core_merchants
WHERE deleted_at IS NULL
ORDER BY name;
