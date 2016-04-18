package models

import (
	"github.com/pasarsakomandiri/shared/database"
)

const (
	DepartmentParking = "Parking"
	DepartmentToycar  = "Toy Car"
	DepartmentOwner   = "Owner"
)

type Department struct {
	ID          int
	Name        string
	Description string
}

func DepartmentGetAll() ([]Department, error) {
	departments := []Department{}
	err := database.Db.Select(&departments, "SELECT * FROM departments")

	return departments, err
}
