SELECT receipt_types.id          AS receipt_type_id,
       receipt_types.description AS receipt_type_description,
       receipt_types.sunat_code  AS receipt_type_sunat_code,
       receipt_types.enable      AS receipt_type_enable,
       receipt_types.created_by  AS receipt_type_created_by,
       receipt_types.created_at  AS receipt_type_created_at
FROM core_receipt_types receipt_types
WHERE receipt_types.deleted_at IS NULL
ORDER BY receipt_types.sunat_code;
