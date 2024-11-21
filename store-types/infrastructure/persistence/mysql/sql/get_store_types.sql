SELECT id,
       description,
       abbreviation
FROM core_store_types
WHERE deleted_at IS NULL
ORDER BY description
LIMIT ? OFFSET ?;