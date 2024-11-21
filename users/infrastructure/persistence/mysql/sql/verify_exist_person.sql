SELECT COUNT(*) AS total
FROM hr_people
WHERE id = ?
  AND deleted_at IS NULL;
