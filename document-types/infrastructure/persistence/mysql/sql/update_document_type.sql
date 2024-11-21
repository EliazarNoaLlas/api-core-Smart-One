UPDATE core_document_types
SET number = ?,
    description = ?,
    abbreviated_description = ?,
    enable = ?
WHERE id = ?;
