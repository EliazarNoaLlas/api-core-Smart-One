SELECT views.id          AS view_id,
       views.name        AS view_name,
       views.description AS view_description,
       views.url         AS view_url,
       views.icon        AS view_icon,
       views.created_at  AS view_created_at
FROM core_views views
WHERE views.deleted_at IS NULL
  AND views.module_id = ?
  AND IF(? IS NULL, TRUE, name LIKE CONCAT('%', TRIM(?), '%')
    OR description LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY views.created_at DESC
LIMIT ? OFFSET ?;
