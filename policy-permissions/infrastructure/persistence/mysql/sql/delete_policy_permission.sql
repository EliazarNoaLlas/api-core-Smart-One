UPDATE core_policy_permissions
SET deleted_at = ?
WHERE id = ? AND policy_id = ?;