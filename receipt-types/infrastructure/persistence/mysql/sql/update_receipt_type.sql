UPDATE core_receipt_types
SET description = ?,
    sunat_code  = ?,
    enable      = ?
WHERE id = ?;