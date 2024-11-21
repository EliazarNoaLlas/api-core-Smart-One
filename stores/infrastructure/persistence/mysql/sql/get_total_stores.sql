SELECT COUNT(*) AS total
FROM core_stores stores
         INNER JOIN core_store_types types ON stores.store_type_id = types.id
WHERE stores.merchant_id = ?
  AND stores.deleted_at IS NULL
ORDER BY stores.created_at DESC;

