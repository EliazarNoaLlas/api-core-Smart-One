UPDATE core_modules
SET name        = TRIM(?),
    description = TRIM(?),
    code        = TRIM(?),
    icon        = TRIM(?),
    position    =?
WHERE id = ?;
