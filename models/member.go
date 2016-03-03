package models
import "github.com/jmoiron/sqlx"

type Member struct {
	Id int
	Vehicle_number string
	Description string
	Created_date string
	Created_by int64
	Last_update_date string
	Updated_by int64
}

func MemberCreateNew(db *sqlx.DB) {

}

func MemberDelete(db *sqlx.DB) {

}

func MemberUpdate(db *sqlx.DB) {

}