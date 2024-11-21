package mysql

import "time"

type DocumentType struct {
	Id                     string     `db:"id"`
	Number                 string     `db:"number"`
	Description            string     `db:"description"`
	AbbreviatedDescription string     `db:"abbreviated_description"`
	Enable                 int        `db:"enable"`
	CreatedAt              *time.Time `db:"created_at"`
}
