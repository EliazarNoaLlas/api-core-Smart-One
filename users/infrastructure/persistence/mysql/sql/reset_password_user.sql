UPDATE core_users
SET password_hash = TRIM(?)
WHERE id = ?;
