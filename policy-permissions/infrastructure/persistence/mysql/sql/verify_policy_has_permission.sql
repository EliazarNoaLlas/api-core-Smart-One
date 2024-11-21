SELECT COUNT(*) AS total
FROM core_policy_permissions policyPermission
WHERE policyPermission.deleted_at IS NULL
  AND policy_id = ?
  AND permission_id = ?;

