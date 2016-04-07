package models

import (
	"github.com/pasarsakomandiri/shared/database"
	_"log"
)

type ParkingTransactionsCashier struct {
    Verified_by int64
    Username string
    Vehicle_id int
    Parking_cost int
}

func ParkingTransactionGetAllAPI(condition string) ([]ParkingTicket, error) {
	parkingTicket := []ParkingTicket{}

	//default condition
	if condition == "" {
		condition = "1=1"
	}

	/*queryString := "SELECT id" +
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
			" FROM parking_transactions WHERE " + condition*/
	//err := database.Db.Select(&parkingTicket, "SELECT id, ticket_number, vehicle_id, vehicle_type, COALESCE(vehicle_number, ' ') vehicle_number, COALESCE(out_date, ' ') out_date, COALESCE(parking_cost, 0) parking_cost, COALESCE(verified_by, 0) verified_by, created_date, created_by, COALESCE(last_update_date, ' ') last_update_date, COALESCE(updated_by, 0) updated_by FROM parking_transactions WHERE "+condition)
	//err := database.Db.Select(&parkingTicket, queryString)
    err := database.Db.Select(&parkingTicket, "SELECT * FROM parking_transactions_all WHERE " + condition)
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

func PTGetVehicleCountByDate(userID int64, date string, vehicleID int) (int, error) {
    var result int
    queryString := "SELECT COUNT(1) result FROM parking_transactions WHERE out_date <= STR_TO_DATE('" + date + "', '%e %M %Y %H:%i') AND out_date >= DATE_SUB(STR_TO_DATE('" +date+"', '%e %M %Y %H:%i'), INTERVAL 15 HOUR)"
    //log.Println(queryString)
	err := database.Db.Get(&result, queryString + "AND verified_by=? AND vehicle_id=?", userID, vehicleID)
    return result, err
}

func UserParkingTransactions(date string) ([]ParkingTransactionsCashier, error)  {
    result := []ParkingTransactionsCashier{}

	/*//default condition
	if condition == "" {
		condition = "1=1"
	}

	queryString := "SELECT park.verified_by"+
					", (SELECT username FROM user WHERE id = park.verified_by) username"+
					", (SELECT COUNT(1) FROM parking_transactions"+
					"WHERE 1=1"+
					"AND verified_by = park.verified_by"+
					"AND out_date <= NOW() AND out_date >= DATE_SUB(NOW(), INTERVAL 15 HOUR)"+ condition +
					"AND vehicle_id = 2) banyak_mobil"+
					", (SELECT COUNT(1) FROM parking_transactions"+
					"WHERE 1=1"+
					"AND verified_by = park.verified_by"+
					"AND out_date <= NOW() AND out_date >= DATE_SUB(NOW(), INTERVAL 15 HOUR)"+ condition +
					"AND vehicle_id = 1) banyak_motor"+
					", SUM(parking_cost) cost"+
					"FROM parking_transactions park"+
					"WHERE out_date IS NOT NULL"+
					"AND out_date <= NOW() AND out_date >= DATE_SUB(NOW(), INTERVAL 15 HOUR)"+ condition +
					"GROUP BY verified_by, username;"*/

    queryString := "SELECT verified_by, username, vehicle_id, SUM(parking_cost) parking_cost FROM parking_transactions_cashier WHERE out_date <= STR_TO_DATE('" + date + "', '%e %M %Y %H:%i') AND out_date >= DATE_SUB(STR_TO_DATE('" +date+"', '%e %M %Y %H:%i'), INTERVAL 15 HOUR) GROUP BY verified_by, username, vehicle_id"
	//log.Println(queryString)
	err := database.Db.Select(&result, queryString)
	return result, err
}