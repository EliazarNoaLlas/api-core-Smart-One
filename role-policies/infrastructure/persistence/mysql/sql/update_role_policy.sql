UPDATE core_role_policies
SET role_id   = ?,
    policy_id = ?,
    enable    = ?
WHERE id = ?;
