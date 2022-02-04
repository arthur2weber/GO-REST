package models

import (
	"errors"
	"net/url"

	"github.com/jinzhu/gorm"
)

type FleetAlert struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Fleet_ID uint32 `sql:"type:int REFERENCES fleets(id)" json:"fleet_id"`
	Webhook  string `gorm:"size:255;not null;" json:"webhook"`
}

func (a *FleetAlert) Prepare() {
	a.ID = 0
}

func (a *FleetAlert) Validate(action string) error {

	_, err := url.ParseRequestURI(a.Webhook)
	u, errr := url.Parse(a.Webhook)
	if err != nil || errr != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("Invalid URL for Webhook")
	}

	if a.Fleet_ID < 0 {
		return errors.New("Invalid Fleet ID")
	}

	return nil
}

func (a *FleetAlert) SaveFleetAlert(db *gorm.DB) (*FleetAlert, error) {

	var err error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &FleetAlert{}, err
	}
	return a, nil
}

func (a *FleetAlert) FindAllFleetAlert(db *gorm.DB) (*[]FleetAlert, error) {
	var err error
	alerts := []FleetAlert{}
	err = db.Debug().Model(&FleetAlert{}).Find(&alerts).Error
	if err != nil {
		return &[]FleetAlert{}, err
	}
	return &alerts, err
}

func (a *FleetAlert) FindFleetAlertsByID(db *gorm.DB, ID uint32) (*FleetAlert, error) {
	var err error
	err = db.Debug().Model(FleetAlert{}).Where("id = ?", ID).Take(&a).Error
	if err != nil {
		return &FleetAlert{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &FleetAlert{}, errors.New("FleetAlert Not Found")
	}
	return a, err
}

func (a *FleetAlert) FindFleetAlertsByFleetID(db *gorm.DB, Fleet_ID uint32) (*[]FleetAlert, error) {
	var err error
	alerts := []FleetAlert{}
	err = db.Debug().Model(&FleetAlert{}).Where("fleet_id = ?", Fleet_ID).Find(&alerts).Error
	if err != nil {
		return &[]FleetAlert{}, err
	}
	return &alerts, err
}

func (a *FleetAlert) FindWebhooksByFleetID(db *gorm.DB, Fleet_ID uint32) ([]string, error) {
	var err error
	webhooks := []string{}
	err = db.Debug().Model(&FleetAlert{}).Select("webhook").Where("fleet_id = ?", Fleet_ID).Find(&webhooks).Error
	if err != nil {
		return []string{}, err
	}
	return webhooks, err
}
