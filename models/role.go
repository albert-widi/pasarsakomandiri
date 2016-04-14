package models

import (
	"github.com/pasarsakomandiri/shared/database"
)

const (
	Role_level_cashier       = int(1)
	Role_level_office        = int(2)
	Role_level_administrator = int(3)
	Role_level_superuser     = int(99)
)

type Role struct {
	Id          int64
	Role_name   string
	Role_level  int
	Description string
}

func RoleGetAllLimitByLevel(level int) ([]Role, error) {
	role := []Role{}
	err := database.Db.Select(&role, "SELECT id, role_name, role_level, description FROM roles WHERE role_level <= ?", level)
	return role, err
}

func RoleGetAllRole() ([]Role, error) {
	role := []Role{}
	//super user will never be exposed
	err := database.Db.Select(&role, "SELECT id, role_name, role_level, description FROM roles WHERE role_name <> 'Super USer' ORDER BY id")

	return role, err
}
