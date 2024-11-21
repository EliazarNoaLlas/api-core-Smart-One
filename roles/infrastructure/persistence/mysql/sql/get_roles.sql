SELECT id,
       name,
       description,
       enable,
       created_at
FROM core_roles
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?;
