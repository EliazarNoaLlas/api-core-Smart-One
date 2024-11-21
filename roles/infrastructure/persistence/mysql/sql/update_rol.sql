UPDATE core_roles
SET name        = TRIM(?),
    description = TRIM(?),
    enable      = TRIM(?)
WHERE id = ?;
