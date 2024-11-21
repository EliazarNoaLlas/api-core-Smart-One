SELECT COUNT(*) AS total
FROM core_policy_permissions policy_permissions
         LEFT JOIN core_permissions permissions ON permissions.id = policy_permissions.permission_id
WHERE policy_permissions.policy_id = ?
  AND permissions.deleted_at IS NULL
  AND policy_permissions.deleted_at IS NULL;