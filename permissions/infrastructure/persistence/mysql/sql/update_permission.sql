UPDATE core_permissions
SET code        = TRIM(?),
    name        = TRIM(?),
    description = TRIM(?),
    module_id   = ?
WHERE id = ?;
