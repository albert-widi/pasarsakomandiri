package models

import (
	"github.com/pasarsakomandiri/shared/database"
)

type ParkingPrice struct {
	Id int64
	Vehicle_id int
	Vehicle_type string
	First_hour_price int
	Promo_jam_pertama int
	Next_hour_price int
	Promo_jam_berikutnya int
	Maximum_price int
	Created_by int64
	Created_date string
}

func ParkingPriceGetById (id int64)(ParkingPrice, error){
	parkingPrice := ParkingPrice{}
	err := database.Db.Get(&parkingPrice, "SELECT id, vehicle_id, vehicle_type, first_hour_price, next_hour_price, maximum_price FROM parking_price WHERE id=?", id)
	return parkingPrice, err
}

func ParkingPriceGetAllAPI ()([]ParkingPrice, error){
	parkingPrice := []ParkingPrice{}
	err := database.Db.Select(&parkingPrice, "SELECT id, vehicle_id, vehicle_type, first_hour_price, promo_jam_pertama, next_hour_price, promo_jam_berikutnya, maximum_price FROM parking_price")
	return parkingPrice, err
}

func ParkingPriceCreateNew (parkingprice ParkingPrice) error {
	_, err := database.Db.Exec("INSERT INTO parking_price(id, vehicle_id, vehicle_type, first_hour_price, promo_jam_pertama, next_hour_price, promo_jam_berikutnya, maximum_price, created_by, created_date) VALUES(?,?,?,?,?,?,?,?,?,?)", parkingprice.Id, parkingprice.Vehicle_id, parkingprice.Vehicle_type, parkingprice.First_hour_price, parkingprice.Promo_jam_pertama, parkingprice.Next_hour_price, parkingprice.Promo_jam_berikutnya, parkingprice.Maximum_price, parkingprice.Created_by, parkingprice.Created_date)
	return err
}

func ParkingPriceGetByVehicleId (vehicleId int) (ParkingPrice, error){
	parkingPrice := ParkingPrice{}
	err := database.Db.Get(&parkingPrice, "SELECT id, vehicle_type, first_hour_price, next_hour_price, maximum_price FROM parking_price WHERE vehicle_id=?", vehicleId)
	return parkingPrice, err
}

func ParkingPriceUpdate(parkingPrice ParkingPrice) error{
	_, err := database.Db.Exec("UPDATE parking_price SET first_hour_price=?,promo_jam_pertama=?, next_hour_price=?, promo_jam_berikutnya=?, maximum_price=? WHERE id=?", parkingPrice.First_hour_price, parkingPrice.Promo_jam_pertama, parkingPrice.Next_hour_price, parkingPrice.Promo_jam_berikutnya, parkingPrice.Maximum_price, parkingPrice.Id)
	return err
}

