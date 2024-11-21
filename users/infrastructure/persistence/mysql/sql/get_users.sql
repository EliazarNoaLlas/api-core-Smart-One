SELECT users.id          AS user_id,
       users.username    AS user_name,
       users.created_at  AS user_created_at,
       users.type_id     AS user_type_id,
       users.description AS user_type_description,
       users.code        AS user_type_code,
       user_roles.id     AS user_role_id,
       roles.id          AS role_id,
       roles.name        AS role_name,
       roles.description AS role_description,
       roles.enable      AS role_enable,
       roles.created_at  AS role_created_at
FROM (SELECT users.id,
             max(users.created_at) AS max_created_at,
             users.username,
             users.created_at,
             types.id              AS type_id,
             types.description,
             types.code
      FROM core_users users
               INNER JOIN core_user_types types ON users.type_id = types.id
               LEFT JOIN core_user_roles user_roles ON user_roles.user_id = users.id
               lEFT JOIN core_roles roles ON user_roles.role_id = roles.id
      WHERE users.deleted_at IS NULL
        AND IF(? IS NULL, TRUE, types.id = TRIM(?))
        AND IF(? IS NULL, TRUE, users.username LIKE CONCAT('%', TRIM(?), '%'))
        AND user_roles.deleted_at is null
        AND IF(? = '', TRUE, FIND_IN_SET(roles.id, TRIM(?)))
      GROUP BY users.id
      ORDER BY max_created_at DESC
      LIMIT ? OFFSET ?) users
         LEFT JOIN core_user_roles user_roles ON user_roles.user_id = users.id
         lEFT JOIN core_roles roles ON user_roles.role_id = roles.id
WHERE user_roles.deleted_at IS NULL;