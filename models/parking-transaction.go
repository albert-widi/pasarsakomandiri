package models

import (
	"github.com/pasarsakomandiri/shared/database"
)

func ParkingTransactionGetAllAPI(condition string) ([]ParkingTicket, error) {
	parkingTicket := []ParkingTicket{}

	//default condition
	if condition == "" {
		condition = "1=1"
	}

	queryString := "SELECT id" +
			", ticket_number" +
			", vehicle_id" +
			", vehicle_type" +
			", COALESCE(vehicle_number, ' ') vehicle_number" +
			", COALESCE(out_date, ' ') out_date" +
			", COALESCE(parking_cost, 0) parking_cost" +
			", COALESCE(verified_by, 0) verified_by" +
			", created_date, created_by, COALESCE(last_update_date, ' ') last_update_date, COALESCE(updated_by, 0) updated_by" +
			", COALESCE((SELECT CONCAT('/', pic.filepath, '/', pic.filename, '.', pic.format) FROM pictures pic WHERE pic.id = picture_in_id), ' ') picture_path_in" +
			", COALESCE((SELECT CONCAT('/', pic.filepath, '/', pic.filename, '.', pic.format) FROM pictures pic WHERE pic.id = picture_out_id), ' ') picture_path_out" +
			" FROM parking_transactions WHERE " + condition
	//err := database.Db.Select(&parkingTicket, "SELECT id, ticket_number, vehicle_id, vehicle_type, COALESCE(vehicle_number, ' ') vehicle_number, COALESCE(out_date, ' ') out_date, COALESCE(parking_cost, 0) parking_cost, COALESCE(verified_by, 0) verified_by, created_date, created_by, COALESCE(last_update_date, ' ') last_update_date, COALESCE(updated_by, 0) updated_by FROM parking_transactions WHERE "+condition)
	err := database.Db.Select(&parkingTicket, queryString)
	return parkingTicket, err
}

func ParkingTransactions() ([]ParkingTicket, error)  {
	parkingtrans := []ParkingTicket{}
	err := database.Db.Select(&parkingtrans, "SELECT id, ticket_number, vehicle_id, vehicle_type, COALESCE(vehicle_number, ' ') vehicle_number, COALESCE(out_date, ' ') out_date, COALESCE(parking_cost, 0) parking_cost, COALESCE(verified_by, 0) verified_by, created_date, created_by, COALESCE(last_update_date, ' ') last_update_date, COALESCE(updated_by, 0) updated_by FROM parking_transactions")
	return parkingtrans, err
}

func ParkingTransGetByTgl(created_date string) ([]ParkingTicket, error) {
	parkingtgl := []ParkingTicket{}
	err := database.Db.Select(&parkingtgl, "SELECT id, ticket_number, vehicle_id, vehicle_type, COALESCE(vehicle_number, ' ') vehicle_number, COALESCE(out_date, ' ') out_date, COALESCE(parking_cost, 0) parking_cost, COALESCE(verified_by, 0) verified_by, created_date, created_by, COALESCE(last_update_date, ' ') last_update_date, COALESCE(updated_by, 0) updated_by FROM parking_transactions WHERE created_date=?", created_date)
	return parkingtgl, err
}