package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Fleet struct {
	ID        uint32  `gorm:"primary_key;auto_increment" json:"id"`
	Name      string  `gorm:"size:255;not null;" json:"name"`
	Max_Speed float64 `sql:"type:decimal(10,2)" json:"max_speed"`
}

func (f *Fleet) Prepare() {
	f.ID = 0
	f.Name = html.EscapeString(strings.TrimSpace(f.Name))
}

func (f *Fleet) Validate(action string) error {
	if f.Name == "" {
		return errors.New("Required Name")
	}
	if f.Max_Speed <= 0 {
		return errors.New("Required Max_Speed")
	}
	return nil
}

func (f *Fleet) SaveFleet(db *gorm.DB) (*Fleet, error) {

	var err error
	err = db.Debug().Create(&f).Error
	if err != nil {
		return &Fleet{}, err
	}
	return f, nil
}

func (f *Fleet) FindAllFleet(db *gorm.DB) (*[]Fleet, error) {
	var err error
	fleets := []Fleet{}
	err = db.Debug().Model(&Fleet{}).Find(&fleets).Error
	if err != nil {
		return &[]Fleet{}, err
	}
	return &fleets, err
}

func (f *Fleet) FindFleetByID(db *gorm.DB, ID uint32) (*Fleet, error) {
	var err error
	err = db.Debug().Model(Fleet{}).Where("id = ?", ID).Take(&f).Error
	if err != nil {
		return &Fleet{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Fleet{}, errors.New("Fleet Not Found")
	}
	return f, err
}
