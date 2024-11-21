SELECT id,
       name,
       description,
       code,
       icon,
       position,
       created_at
FROM core_modules
WHERE deleted_at IS NULL
ORDER BY position;
