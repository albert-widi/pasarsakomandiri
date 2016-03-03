package models

import (
	"github.com/pasarsakomandiri/shared/database"
	"time"
)

func ParkingCreateNewVehicle(executor int64, vehicle string) bool {
	tx := database.Db.MustBegin()
	trx := tx.MustExec("INSERT INTO vehicle_type(vehicle_type, created_by, created_date) VALUES(?, ?, ?)", vehicle, executor, time.Now().String())

	if _, err := trx.RowsAffected(); err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}

func ParkingVehicleGetAll() ([]Vehicle, error) {
	vehicle := []Vehicle{}
	err := database.Db.Select(&vehicle, "SELECT id, vehicle_type, created_date, created_by FROM vehicle_type")
	return vehicle, err
}

func ParkingVehicleGetByID(id int) (Vehicle, error) {
	vehicle := Vehicle{}
	err := database.Db.Get(&vehicle, "SELECT id, vehicle_type, created_date, created_by FROM vehicle_type WHERE id=?", id)
	//err := db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_number, out_date, parking_cost, verified_by, created_date, created_by, last_update_date, updated_by FROM parking_transactions WHERE ticket_number=?", ticketNumber)
	return vehicle, err
}