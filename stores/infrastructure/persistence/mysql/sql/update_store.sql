UPDATE core_stores
SET name          = TRIM(?),
    shortname     = TRIM(?),
    merchant_id   = ?,
    store_type_id = ?
WHERE id = ?;
