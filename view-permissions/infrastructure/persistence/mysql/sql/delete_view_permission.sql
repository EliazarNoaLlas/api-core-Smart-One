UPDATE core_view_permissions
SET deleted_at = ?
WHERE id = ?
  AND view_id = ?;
