package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Vehicle struct {
	ID        uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Fleet_ID  uint32  `sql:"type:int REFERENCES fleets(id)" json:"fleet_id"`
	Name      string  `gorm:"size:255;not null;" json:"name"`
	Max_Speed float64 `sql:"type:decimal(10,2)" json:"max_speed"`
}

func (a *Vehicle) Prepare() {
	a.ID = 0
}

func (a *Vehicle) Validate(action string) error {

	if a.Name == "" {
		return errors.New("Required Name")
	}

	if a.Fleet_ID < 1 {
		return errors.New("Invalid Fleet ID")
	}

	return nil
}

func (a *Vehicle) SaveVehicle(db *gorm.DB) (*Vehicle, error) {

	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Vehicle{}, err
	}
	return a, nil
}

func (a *Vehicle) FindAllVehicle(db *gorm.DB) (*[]Vehicle, error) {
	var err error

	vehicles := []Vehicle{}

	db.Raw(`
		SELECT 
			v.id id, 
			v.name name, 
			f.id fleet_id, 
			coalesce(nullif(v.max_speed,0),f.max_speed) max_speed 
		FROM vehicles v 
		JOIN fleets f on v.fleet_id = f.id
	`).Scan(&vehicles)

	return &vehicles, err
}

func (a *Vehicle) FindVehicleByID(db *gorm.DB, ID uint32) (*Vehicle, error) {
	var err error
	err = db.Debug().Model(Vehicle{}).Where("id = ?", ID).Take(&a).Error
	if err != nil {
		return &Vehicle{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Vehicle{}, errors.New("Vehicle Not Found")
	}
	return a, err
}
