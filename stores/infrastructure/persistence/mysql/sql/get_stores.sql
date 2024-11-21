SELECT stores.id          AS store_id,
       stores.name        AS store_name,
       stores.shortname   AS store_shortname,
       stores.merchant_id AS store_merchant_id,
       stores.created_at  AS store_created_at,
       types.id           AS type_id,
       types.description  AS type_description,
       types.abbreviation AS type_abbreviation
FROM core_stores stores
         INNER JOIN core_store_types types ON stores.store_type_id = types.id
WHERE stores.merchant_id = ?
  AND stores.deleted_at IS NULL
ORDER BY stores.created_at DESC
LIMIT ? OFFSET ?;
