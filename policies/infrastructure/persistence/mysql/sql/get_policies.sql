SELECT policies.id               AS policy_id,
       policies.name             AS policy_name,
       policies.description      AS policy_description,
       policies.level            AS policy_level,
       policies.enable           AS policy_enable,
       policies.created_at       AS policy_created_at,
       modules.id                AS module_id,
       modules.name              AS module_name,
       modules.description       AS module_description,
       modules.code              AS module_code,
       merchants.id              AS merchant_id,
       merchants.name            AS merchant_name,
       merchants.description     AS merchant_description,
       merchants.document        AS merchant_document,
       stores.id                 AS store_id,
       stores.name               AS store_name,
       stores.shortname          AS store_shortname,
       permissions.id            AS permission_id,
       permissions.code          AS permission_code,
       permissions.name          AS permission_name,
       permissions.description   AS permission_description,
       policy_permissions.id     AS policy_permission_id,
       policy_permissions.enable AS policy_permission_enable
FROM core_policies policies
         LEFT JOIN core_modules modules ON policies.module_id = modules.id
         LEFT JOIN core_merchants merchants ON policies.merchant_id = merchants.id
         LEFT JOIN core_stores stores ON policies.store_id = stores.id
         LEFT JOIN core_permissions permissions ON policies.module_id = permissions.module_id
         LEFT JOIN core_policy_permissions policy_permissions ON permissions.id = policy_permissions.permission_id
WHERE policies.deleted_at IS NULL
  AND permissions.deleted_at IS NULL
  AND policy_permissions.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, modules.id = TRIM(?))
  AND IF(? IS NULL, TRUE, merchants.id = TRIM(?))
  AND IF(? IS NULL, TRUE, stores.id = TRIM(?))
  AND IF(? IS NULL, TRUE, policies.name LIKE CONCAT('%', TRIM(?), '%') OR
                          policies.description LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY policies.created_at DESC
LIMIT ? OFFSET ?;
