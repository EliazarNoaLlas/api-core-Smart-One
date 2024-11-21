SELECT distinct stores.id             AS store_id,
                stores.name           AS store_name,
                merchants.id          AS merchant_id,
                merchants.name        AS merchant_name,
                merchants.description AS merchant_description,
                merchants.image_path  AS merchant_image_path
FROM core_users users
         INNER JOIN core_user_roles users_roles ON users.id = users_roles.user_id
         INNER JOIN core_role_policies role_policies ON users_roles.role_id = role_policies.role_id
         INNER JOIN core_policies policies ON policies.id = role_policies.policy_id
         INNER JOIN core_stores stores ON stores.id = policies.store_id
         INNER JOIN core_merchants merchants ON merchants.id = stores.merchant_id
WHERE users.id = ?
  AND users.deleted_at IS NULL
  AND users_roles.deleted_at IS NULL
  AND role_policies.deleted_at IS NULL
  AND policies.deleted_at IS NULL
  AND users_roles.enable IS TRUE
  AND role_policies.enable IS TRUE
  AND policies.enable IS TRUE
GROUP BY stores.id;



