package models
import (
	"github.com/jmoiron/sqlx"
	"github.com/pasarsakomandiri/shared/database"
)

type Member struct {
	Id int
	Vehicle_id int
	Vehicle_type string
	Police_number string
	Description string
	Created_date string
	Created_by int64
}

func MemberCreateNew(member Member) error {
	_, err := database.Db.Exec("INSERT INTO members(id, vehicle_id, vehicle_type, police_number, description, created_by, created_date)VALUES(?, ?, ?, ?, ?, ?, ?)", member.Id, member.Vehicle_id, member.Vehicle_type, member.Police_number, member.Description, member.Created_by, member.Created_date)
	return err
}

func MemberGetAll()([]Member, error)  {
	member := []Member{}
	err := database.Db.Select(&member, "SELECT id, vehicle_id, vehicle_type, police_number, description, created_by, created_date FROM members ORDER BY id")
	return member, err
}

func MemberDelete(memberId int) error {
	_, err := database.Db.Exec("DELETE FROM members WHERE id=?", memberId)
	return err
}

func MemberUpdate(db *sqlx.DB) {

}