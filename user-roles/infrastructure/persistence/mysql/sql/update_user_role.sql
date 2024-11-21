UPDATE core_user_roles
SET user_id = ?,
    role_id = ?,
    enable  = ?
WHERE id = ?;
