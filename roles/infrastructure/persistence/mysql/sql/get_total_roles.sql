SELECT COUNT(*)
FROM core_roles
WHERE deleted_at IS NULL
ORDER BY name;
