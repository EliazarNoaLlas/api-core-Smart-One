SELECT COUNT(*) AS total
FROM core_role_policies role_policies
WHERE role_policies.deleted_at IS NULL
  AND role_policies.role_id = ?
  AND role_policies.policy_id = ?;
