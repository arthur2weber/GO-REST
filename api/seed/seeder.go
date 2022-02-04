package seed

import (
	"log"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/jinzhu/gorm"
)

var fleet = []models.Fleet{
	{
		Name:      "Frota 1",
		Max_Speed: 55.55,
	},
	{
		Name:      "Frota 2",
		Max_Speed: 22.22,
	},
	{
		Name:      "Frota 3",
		Max_Speed: 33.00,
	},
}

var fleetAlert = []models.FleetAlert{
	{
		Fleet_ID: 1,
		Webhook:  "http://localhost:8080/fleets",
	},
	{
		Fleet_ID: 2,
		Webhook:  "http://localhost:8080/fleets",
	},
	{
		Fleet_ID: 1,
		Webhook:  "http://localhost:8080/fleets?teste",
	},
}

var vehicles = []models.Vehicle{
	{
		Fleet_ID:  1,
		Name:      "Veículos 1",
		Max_Speed: 10.1,
	},
	{
		Fleet_ID:  2,
		Name:      "Veículos 2",
		Max_Speed: 20.2,
	},
	{
		Fleet_ID:  1,
		Name:      "Veículos 3",
		Max_Speed: 30.3,
	},
	{
		Fleet_ID:  3,
		Name:      "Veículos 3",
		Max_Speed: 40.4,
	},
}

func Load(db *gorm.DB) {

	err := db.DropTableIfExists(&models.FleetAlert{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.DropTableIfExists(&models.VehiclePosition{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.DropTableIfExists(&models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.DropTableIfExists(&models.Fleet{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Fleet{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.FleetAlert{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.VehiclePosition{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/* ---------------------------------------------------------------------- */

	for i := range fleet {
		err = db.Debug().Model(&models.Fleet{}).Create(&fleet[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Fleet table: %v", err)
		}
	}
	for i := range fleetAlert {
		err = db.Debug().Model(&models.FleetAlert{}).Create(&fleetAlert[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Fleet table: %v", err)
		}
	}
	for i := range vehicles {
		err = db.Debug().Model(&models.Vehicle{}).Create(&vehicles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed Vehicle table: %v", err)
		}
	}

}
