UPDATE core_policies
SET name        = TRIM(?),
    description = TRIM(?),
    module_id   = ?,
    merchant_id = ?,
    store_id    = ?,
    level       = TRIM(?),
    enable      = ?
WHERE id = ?;
