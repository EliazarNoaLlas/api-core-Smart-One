SELECT role_policies.id         AS role_policy_id,
       role_policies.enable     AS role_policy_enable,
       role_policies.created_at AS role_policy_created_at,
       policies.id              AS policy_id,
       policies.name            AS policy_name,
       policies.description     AS policy_description,
       policies.level           AS policy_level,
       policies.enable          AS policy_enable,
       policies.created_at      AS policy_created_at
FROM core_role_policies role_policies
         LEFT JOIN core_policies policies ON role_policies.policy_id = policies.id
WHERE role_policies.deleted_at IS NULL
  AND policies.deleted_at IS NULL
  AND role_policies.role_id = ?
ORDER BY role_policies.created_at DESC
LIMIT ? OFFSET ?;
