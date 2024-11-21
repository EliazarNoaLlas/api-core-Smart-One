SELECT id,
       name,
       description,
       phone,
       document,
       address,
       industry,
       image_path,
       created_at
FROM core_merchants
WHERE deleted_at IS NULL
ORDER BY name
LIMIT ? OFFSET ?;
