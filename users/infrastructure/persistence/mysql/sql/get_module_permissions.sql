SELECT DISTINCT DISTINCT core_modules.id, core_permissions.code
FROM core_user_roles
    INNER JOIN core_roles
ON core_user_roles.role_id = core_roles.id AND core_roles.deleted_at IS NULL
    INNER JOIN core_role_policies
    ON core_roles.id = core_role_policies.role_id AND core_role_policies.deleted_at IS NULL
    INNER JOIN core_policies
    ON core_role_policies.policy_id = core_policies.id AND core_policies.deleted_at IS NULL
    INNER JOIN core_merchants
    ON core_policies.merchant_id = core_merchants.id AND core_merchants.deleted_at IS NULL
    INNER JOIN core_stores ON core_merchants.id = core_stores.merchant_id AND core_stores.deleted_at IS NULL
    INNER JOIN core_policy_permissions ON core_policies.id = core_policy_permissions.policy_id AND
    core_policy_permissions.deleted_at IS NULL
    INNER JOIN core_permissions ON core_policy_permissions.permission_id = core_permissions.id AND
    core_permissions.deleted_at IS NULL
    INNER JOIN core_modules ON core_permissions.module_id = core_modules.id
WHERE core_user_roles.deleted_at IS NULL
  AND core_modules.code = ?
  AND core_user_roles.user_id = ?;