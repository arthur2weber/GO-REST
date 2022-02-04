package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type VehiclePosition struct {
	ID            uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Vehicle_ID    uint32  `sql:"type:int REFERENCES vehicles(id)" json:"vehicle_id"`
	Latitude      float64 `sql:"type:decimal(10,2)" json:"latitude"`
	Longitude     float64 `sql:"type:decimal(10,2)" json:"longitude"`
	Current_Speed float64 `sql:"type:decimal(10,2)" json:"current_speed"`
	Max_Speed     float64 `sql:"type:decimal(10,2)" json:"max_speed"`
}

func (a *VehiclePosition) Prepare() {
	a.ID = 0
}

func (a *VehiclePosition) Validate(action string) error {

	if a.Latitude == 0 {
		return errors.New("Required latitude")
	}

	if a.Longitude == 0 {
		return errors.New("Required longitude")
	}

	if a.Current_Speed < 0 {
		return errors.New("Required current_speed")
	}

	if a.Vehicle_ID < 1 {
		return errors.New("Invalid Vehicle ID")
	}

	return nil
}

func (a *VehiclePosition) SaveVehiclePosition(db *gorm.DB) (*VehiclePosition, error) {

	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &VehiclePosition{}, err
	}
	return a, nil
}

func (a *VehiclePosition) FindAllVehiclePosition(db *gorm.DB) (*[]VehiclePosition, error) {
	var err error

	positions := []VehiclePosition{}

	db.Raw(`
	SELECT 
		v.id id, 
		v.name name, 
		f.id fleet_id, 
		coalesce(nullif(v.max_speed,0),f.max_speed) max_speed 
	FROM vehicles v
	JOIN vehicle_positions vp on vp.vehicle_id = v.id 
	JOIN fleets f on v.fleet_id = f.id
	`).Scan(&positions)

	return &positions, err
}

func (a *VehiclePosition) FindVehiclePositionByVehicleId(db *gorm.DB, Vehicle_ID uint32) (*[]VehiclePosition, error) {
	var err error

	positions := []VehiclePosition{}

	db.Raw(`
	SELECT 
		v.id id, 
		v.name name, 
		f.id fleet_id, 
		coalesce(nullif(v.max_speed,0),f.max_speed) max_speed 
	FROM vehicles v
	JOIN vehicle_positions vp on vp.vehicle_id = v.id 
	JOIN fleets f on v.fleet_id = f.id
	where v.id = ?
	`, Vehicle_ID).Scan(&positions)

	return &positions, err
}

func (a *VehiclePosition) FindVehiclePositionByID(db *gorm.DB, Vehicle_ID uint32) (*[]VehiclePosition, error) {
	var err error
	positions := []VehiclePosition{}
	err = db.Debug().Model(VehiclePosition{}).Where("vehicle_id = ?", Vehicle_ID).Find(&positions).Error
	if err != nil {
		return &positions, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]VehiclePosition{}, errors.New("Vehicle Not Found")
	}
	return &positions, err
}
