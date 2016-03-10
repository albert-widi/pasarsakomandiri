package models
import (
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

func MemberGetByPoliceNumber(policeNumber string) (Member, error) {
	member := Member{}
	err := database.Db.Get(&member, "SELECT id, vehicle_id, vehicle_type, police_number, description, created_by, created_date FROM members WHHERE police_number=?", policeNumber)
	return member, err
}

func MemberGetAll()([]Member, error)  {
	members := []Member{}
	err := database.Db.Select(&members, "SELECT id, vehicle_id, vehicle_type, police_number, description, created_by, created_date FROM members ORDER BY id")
	return members, err
}

func MemberDelete(memberId int) error {
	_, err := database.Db.Exec("DELETE FROM members WHERE id=?", memberId)
	return err
}