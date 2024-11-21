SELECT COUNT(*) as total
FROM hr_people
WHERE type_document_id = ?
  AND document = ?
  AND deleted_at IS NULL;