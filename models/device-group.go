package models

import (
	"github.com/pasarsakomandiri/shared/database"
	"time"
)

func DeviceGroupCreateNew(executor int64, deviceGroup DeviceGroup) error {
	_, err := database.Db.Exec("INSERT INTO device_groups(raspberry_id, raspberry_ip, camera_id, camera_ip, vehicle_id, vehicle_type, gate_name, created_by, created_date) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", deviceGroup.Raspberry_id, deviceGroup.Raspberry_ip, deviceGroup.Camera_id, deviceGroup.Camera_ip, deviceGroup.Vehicle_id, deviceGroup.Vehicle_type, deviceGroup.Gate_name, executor, time.Now().String())
	return err
}

//device group get all
func DeviceGroupGetAllAPI ()([]DeviceGroup, error){
	var deviceGroup = []DeviceGroup{}
	err  := database.Db.Select(&deviceGroup, "SELECT id, raspberry_id, raspberry_ip, camera_id, camera_ip, vehicle_id, vehicle_type, gate_name FROM device_groups")
	return deviceGroup, err
}

func DeviceGroupDeleteById (deviceId int64) error{
	_, err := database.Db.Exec("DELETE FROM device_groups WHERE id=?", deviceId)
	return err
}