package models
import (
	"database/sql"
	"github.com/pasarsakomandiri/shared/database"
)

//ParkingTicket data
type ParkingTicket struct {
	Id int64
	Ticket_number string
	Vehicle_id int
	Vehicle_type string
	Vehicle_number string
	Out_date string
	Parking_cost int
	Verified_by int64
	Created_date string
	Created_by int64
	Last_update_date string
	Updated_by int64
	Picture_in_id int64
	Picture_path_in string
	Picture_out_id int64
	Picture_path_out string
    Duration string
}

type Vehicle struct {
	Id int
	Vehicle_type string
	Created_by int64
	Created_date string
}

func ParkingGetTicketByTicketNumber(ticketNumber string) (ParkingTicket, error) {
	parkingTicket := ParkingTicket{}

	err := database.Db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_id, vehicle_type, created_date, created_by, picture_in_id, picture_out_id FROM parking_transactions WHERE ticket_number=? AND out_date IS NULL", ticketNumber)
	//err := db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_number, out_date, parking_cost, verified_by, created_date, created_by, last_update_date, updated_by FROM parking_transactions WHERE ticket_number=?", ticketNumber)
	return parkingTicket, err
}

func ParkingGetTicketByNumberAndId(ticketId int64, ticketNumber string) (ParkingTicket, error) {
	parkingTicket := ParkingTicket{}

	err := database.Db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_id, vehicle_type, created_date, created_by, picture_in_id, picture_out_id FROM parking_transactions WHERE id=? AND ticket_number=?", ticketId, ticketNumber)
	//err := db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_number, out_date, parking_cost, verified_by, created_date, created_by, last_update_date, updated_by FROM parking_transactions WHERE ticket_number=?", ticketNumber)
	return parkingTicket, err
}

func ParkingGetTicketByTicketID(Id string) (ParkingTicket, error) {
	parkingTicket := ParkingTicket{}

	err :=database.Db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_number, out_date, parking_cost, verified_by, created_date, created_by, last_update_date, updated_by FROM parking_transactions WHERE idr=?", Id)
	return parkingTicket, err
}

func ParkingTicketExistsByNumber(ticketNumber string) (ParkingTicket, error) {
	parkingTicket := ParkingTicket{}

	err := database.Db.Get(&parkingTicket, "SELECT id, ticket_number, vehicle_number, out_date, parking_cost, verified_by, created_date, created_by, last_update_date, updated_by FROM parking_transactions WHERE ticket_number=? AND out_date = NULL", ticketNumber)
	return parkingTicket, err
}

func ParkingCreateNewTicket(parkingTicket ParkingTicket) (sql.Result, error) {
	result, err := database.Db.Exec("INSERT INTO parking_transactions(ticket_number, vehicle_id, vehicle_type, created_by, created_date, picture_in_id) VALUES(?, ?, ?, ?, ?, ?)", parkingTicket.Ticket_number, parkingTicket.Vehicle_id, parkingTicket.Vehicle_type, parkingTicket.Created_by, parkingTicket.Created_date, parkingTicket.Picture_in_id)
	return result, err
}

func ParkingUpdateTicket(parkingTicket ParkingTicket) error {
	_, err := database.Db.Exec("UPDATE parking_transactions SET vehicle_number=?, out_date=?, parking_cost=?, verified_by=?, last_update_date=?, updated_by=?, picture_out_id=?, vehicle_number=? WHERE id=?", parkingTicket.Vehicle_number, parkingTicket.Out_date, parkingTicket.Parking_cost, parkingTicket.Verified_by, parkingTicket.Last_update_date, parkingTicket.Updated_by, parkingTicket.Picture_out_id, parkingTicket.Vehicle_number, parkingTicket.Id)
	return err
}

func ParkingTicketUpdatePictureIn(pictureId int64, ticketId int64) error {
	_, err := database.Db.Exec("UPDATE parking_transactions SET picture_in_id = ? WHERE id = ?", pictureId, ticketId)
	return err
}

func ParkingTicketUpdatePictureOut(pictureId int64, ticketId int64) error {
	_, err := database.Db.Exec("UPDATE parking_transactions SET picture_out_id = ? WHERE id = ?", pictureId, ticketId)
	return err
}