package models

import (
	"github.com/pasarsakomandiri/shared/database"
	"time"
)

type DeviceGroup struct {
	Id int64
	Group_type string
	Raspberry_id int64
	Raspberry_ip string
	Camera_id int64
	Camera_ip string
	Vehicle_id int
	Vehicle_type string
	Group_name string
}

func DeviceGroupCreateNew(executor int64, deviceGroup DeviceGroup) error {
	_, err := database.Db.Exec("INSERT INTO device_groups(raspberry_id, group_type, raspberry_ip, camera_id, camera_ip, vehicle_id, vehicle_type, group_name, created_by, created_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", deviceGroup.Raspberry_id, deviceGroup.Group_type, deviceGroup.Raspberry_ip, deviceGroup.Camera_id, deviceGroup.Camera_ip, deviceGroup.Vehicle_id, deviceGroup.Vehicle_type, deviceGroup.Group_name, executor, time.Now().String())
	return err
}

//device group get all
func DeviceGroupGetAllAPI ()([]DeviceGroup, error){
	var deviceGroup = []DeviceGroup{}
	err  := database.Db.Select(&deviceGroup, "SELECT id, raspberry_id, group_type, raspberry_ip, camera_id, camera_ip, vehicle_id, vehicle_type, group_name FROM device_groups")
	return deviceGroup, err
}

func DeviceGroupDeleteById (deviceId int64) error{
	_, err := database.Db.Exec("DELETE FROM device_groups WHERE id=?", deviceId)
	return err
}