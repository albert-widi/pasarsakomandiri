package models

import (
	"github.com/jmoiron/sqlx"
	"github.com/pasarsakomandiri/shared/database"
)

type Device struct {
	Id int64
	Device_type_id int64
	Device_type string
	Device_name string
	Host string
	Token string
	Description string
	Created_by int64
	Created_date string
}

type DeviceType struct {
	Id int64
	Device_type string
}

type ParkingGate struct {
	id int
	Device_group_id int64
	Vehicle_type_id int
	Vehicle_type string
	Gate_name  string
	Created_by int64
	Created_date string
}

//GET
func DeviceGetAll() ([]Device, error) {
	device := []Device{}
	err := database.Db.Select(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list ORDER BY id")
	return device, err
}

func DeviceGetAllDeviceType() ([]DeviceType, error){
	device := []DeviceType{}
	err := database.Db.Select(&device, "SELECT id, device_type FROM device_type")
	return device, err
}

func DeviceGetDeviceTypeByName(deviceType string) (DeviceType, error) {
	device := DeviceType{}
	err := database.Db.Get(&device, "SELECT id, device_type FROM device_type WHERE device_type=?", deviceType)
	return device, err
}

func DeviceGetDeviceTypeByID(id string) (DeviceType, error) {
	device := DeviceType{}
	err := database.Db.Get(&device, "SELECT id, device_type FROM device_type WHERE id=?", id)
	return device, err
}

func DeviceGetByID(id int64) (Device, error) {
	device := Device{}
	err := database.Db.Get(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list WHERE id=?", id)

	return device, err
}

func DeviceGetByHost(host string) (Device, error) {
	device := Device{}
	err := database.Db.Get(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list WHERE host=?", host)

	return device, err
}

func DeviceGetByHostType(host string, deviceType string) (Device, error) {
	device := Device{}
	err := database.Db.Get(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list WHERE host=? AND device_type=?", host, deviceType)

	return device, err
}

func DeviceGetByType(deviceId int, deviceType string) (Device, error) {
	device := Device{}
	err := database.Db.Get(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list WHERE id=? AND device_type=?", deviceId, deviceType)
	return device, err
}

func DevicesGetByType(deviceType string) ([]Device, error) {
	device := []Device{}
	err := database.Db.Select(&device, "SELECT id, device_type_id, device_type, device_name, host, token, description FROM device_list WHERE device_type=?", deviceType)

	return device, err
}

//device group get by host
func DeviceGroupGetByHost(param string, host string) (DeviceGroup, error) {
	deviceGroup := DeviceGroup{}
	selectString := "SELECT id, raspberry_id, raspberry_ip, camera_id, camera_ip, vehicle_id, vehicle_type FROM device_groups WHERE " + param + "=?"
	err := database.Db.Get(&deviceGroup, selectString, host)
	return deviceGroup, err
}

//device group get by id
func DeviceGroupGetByID( id int64) (DeviceGroup, error) {
	deviceGroup := DeviceGroup{}
	err := database.Db.Get(&deviceGroup, "SELECT id, raspberry_id, raspberry_ip, camera_id, camera_ip FROM device_groups WHERE id=?", id)
	return deviceGroup, err
}

//GET END --------------------------

//INSERT
func DeviceCreateNew(device Device) error {
	_, err := database.Db.Exec("INSERT INTO device_list(device_type_id, device_type, device_name, host, token, description, created_by, created_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", device.Device_type_id, device.Device_type, device.Device_name, device.Host, device.Token, device.Description, device.Created_by, device.Created_date)
	return err
}

//EDIT
func DeviceEdit(device Device) error {
	_, err := database.Db.Exec("UPDATE device_list SET device_type=?, device_name=?, host=?, token=?, description=? WHERE id=?", device.Device_type, device.Device_name, device.Host, device.Token, device.Description, device.Id)
	return err
}

//DELETE
func DeviceDeleteByID(db *sqlx.DB, device Device) bool {
	tx := database.Db.MustBegin()
	trx := tx.MustExec("DELETE FROM device_list WHERE id=?", device.Id)

	if _, err := trx.RowsAffected(); err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()

	return true
}

func DeleteDeviceById(device Device) error {
	_, err := database.Db.Exec("DELETE FROM device_list WHERE id=?", device.Id)
	return err
}