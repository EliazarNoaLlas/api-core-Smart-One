SELECT COUNT(*) AS total
FROM core_policies policies
         LEFT JOIN core_modules modules ON policies.module_id = modules.id
         LEFT JOIN core_merchants merchants ON policies.merchant_id = merchants.id
         LEFT JOIN core_stores stores ON policies.store_id = stores.id
WHERE policies.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, modules.id = TRIM(?))
  AND IF(? IS NULL, TRUE, merchants.id = TRIM(?))
  AND IF(? IS NULL, TRUE, stores.id = TRIM(?))
  AND IF(? IS NULL, TRUE, policies.name LIKE CONCAT('%', TRIM(?), '%') OR
                          policies.description LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY policies.created_at DESC;
