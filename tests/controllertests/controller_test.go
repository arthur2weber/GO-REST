package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/arthur2weber/go_rest/api/controllers"
	"github.com/arthur2weber/go_rest/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var fleetInstance = models.Fleet{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database\n")
	}

}

func RefreshTable() error {

	err := server.DB.DropTableIfExists(&models.FleetAlert{}).Error
	if err != nil {
		return err
	}

	err = server.DB.DropTableIfExists(&models.Fleet{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.Fleet{}).Error
	if err != nil {
		return err
	}

	err = server.DB.Debug().AutoMigrate(&models.FleetAlert{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneFleet() (models.Fleet, error) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}

	fleet := models.Fleet{
		Name:      "Frota teste 1",
		Max_Speed: 12.34,
	}

	err = server.DB.Model(&models.Fleet{}).Create(&fleet).Error
	if err != nil {
		return models.Fleet{}, err
	}
	return fleet, nil
}

func seedFleets() ([]models.Fleet, error) {

	fleets := []models.Fleet{
		{
			Name:      "Frota teste 10",
			Max_Speed: 34.56,
		},
		{
			Name:      "Frota teste 20",
			Max_Speed: 88.88,
		},
	}

	for i := range fleets {
		err := server.DB.Model(&models.Fleet{}).Create(&fleets[i]).Error
		if err != nil {
			return []models.Fleet{}, err
		}
	}
	return fleets, nil
}

func seedFleetsAlerts() ([]models.FleetAlert, error) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}
	seedOneFleet()

	fleetsAlerts := []models.FleetAlert{
		{
			Fleet_ID: 1,
			Webhook:  "http://test2.com",
		},
		{
			Fleet_ID: 1,
			Webhook:  "http://test3.com",
		},
	}

	for i := range fleetsAlerts {
		err := server.DB.Model(&models.FleetAlert{}).Create(&fleetsAlerts[i]).Error
		if err != nil {
			return []models.FleetAlert{}, err
		}
	}
	return fleetsAlerts, nil
}
