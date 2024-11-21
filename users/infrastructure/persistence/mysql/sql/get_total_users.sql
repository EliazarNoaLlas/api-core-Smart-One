SELECT COUNT(*) AS total
FROM core_users users
         INNER JOIN core_user_types types ON users.type_id = types.id
WHERE users.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, types.id = TRIM(?))
  AND IF(? IS NULL, TRUE, users.username LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY users.created_at DESC;

