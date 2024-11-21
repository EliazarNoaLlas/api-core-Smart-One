SELECT users.id                               AS user_id,
       users.username                         AS user_name,
       users.created_at                       AS user_created_at,
       people.id                              AS person_id,
       people.document                        AS person_document,
       people.names                           AS person_names,
       people.surname                         AS person_surname,
       people.last_name                       AS person_last_name,
       people.phone                           AS person_phone,
       people.email                           AS person_email,
       people.gender                          AS person_gender,
       people.enable                          AS person_enable,
       people.created_at                      AS person_created_at,
       document_types.id                      AS document_type_id,
       document_types.number                  AS document_type_number,
       document_types.description             AS document_type_description,
       document_types.abbreviated_description AS document_type_abbreviated_description,
       document_types.enable                  AS document_type_enable,
       document_types.created_at              AS document_type_created_at,
       roles.id                               AS role_id,
       roles.name                             AS role_name,
       roles.description                      AS role_description,
       roles.enable                           AS role_enable,
       roles.created_at                       AS role_created_at
FROM core_users users
         LEFT JOIN hr_people people ON users.id = people.user_id
         LEFT JOIN core_document_types document_types ON people.type_document_id = document_types.id
         INNER JOIN core_user_roles user_roles ON users.id = user_roles.user_id
         INNER JOIN core_roles roles ON user_roles.role_id = roles.id
WHERE users.id = ?
  AND users.deleted_at IS NULL
  AND people.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND document_types.deleted_at is NULL
  AND roles.deleted_at IS NULL
ORDER BY users.created_at;
