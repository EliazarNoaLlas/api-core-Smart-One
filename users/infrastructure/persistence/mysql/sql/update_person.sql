UPDATE hr_people
SET user_id          = ?,
    type_document_id = ?,
    document         = ?,
    names            = ?,
    surname          = ?,
    last_name        = ?,
    phone            = ?,
    email            = ?,
    gender           = ?,
    enable           = ?
WHERE id = ?;