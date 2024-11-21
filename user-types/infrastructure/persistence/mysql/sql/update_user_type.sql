UPDATE core_user_types
SET description = TRIM(?),
    code        = TRIM(?),
    enable      = ?
WHERE id = ?;
