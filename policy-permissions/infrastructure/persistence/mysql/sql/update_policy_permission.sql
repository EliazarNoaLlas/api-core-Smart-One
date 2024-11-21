UPDATE core_policy_permissions
SET policy_id     = ?,
    permission_id = TRIM(?),
    enable        = ?
WHERE id = ?;
