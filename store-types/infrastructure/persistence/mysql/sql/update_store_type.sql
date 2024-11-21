UPDATE core_store_types
SET description  = TRIM(?),
    abbreviation = TRIM(?)
WHERE id = ?;