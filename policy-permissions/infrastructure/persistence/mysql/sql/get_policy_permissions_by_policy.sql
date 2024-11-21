SELECT policy_permissions.id         AS policy_permission_id,
       policy_permissions.enable     AS policy_permission_enable,
       policy_permissions.created_at AS policy_permission_created_at,
       permissions.id                AS permissions_id,
       permissions.code              AS permissions_code,
       permissions.name              AS permissions_name,
       permissions.description       AS permissions_description,
       permissions.created_at        AS permissions_created_at
FROM core_policy_permissions policy_permissions
         LEFT JOIN core_permissions permissions ON permissions.id = policy_permissions.permission_id
WHERE policy_permissions.policy_id = ?
  AND permissions.deleted_at IS NULL
  AND policy_permissions.deleted_at IS NULL
ORDER BY policy_permissions.created_at DESC
LIMIT ? OFFSET ?;
