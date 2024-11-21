SELECT COUNT(*) as total
FROM core_role_policies role_policies
         LEFT JOIN core_policies policies ON role_policies.policy_id = policies.id
WHERE role_policies.deleted_at IS NULL
  AND policies.deleted_at IS NULL
  AND role_policies.role_id = ?;
