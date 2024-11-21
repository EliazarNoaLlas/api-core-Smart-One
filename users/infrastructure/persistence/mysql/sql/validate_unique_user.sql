SELECT COUNT(*) AS total
FROM hr_people
WHERE user_id = ?
    AND deleted_at IS NULL;