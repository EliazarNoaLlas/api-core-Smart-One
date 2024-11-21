UPDATE core_merchants
SET name        = TRIM(?),
    description = TRIM(?),
    phone       = TRIM(?),
    document    = TRIM(?),
    address     = TRIM(?),
    industry    = TRIM(?),
    image_path  = TRIM(?)
WHERE id = ?;
