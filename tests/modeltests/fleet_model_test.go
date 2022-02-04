package modeltests

import (
	"log"
	"testing"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllFleet(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error refreshing Fleet table %v\n", err)
	}

	err = seedFleets()
	if err != nil {
		log.Fatalf("Error seeding Fleet table %v\n", err)
	}

	fleets, err := fleetInstance.FindAllFleet(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Fleets: %v\n", err)
		return
	}
	assert.Equal(t, len(*fleets), 2)
}

func TestSaveFleet(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error fleet refreshing table %v\n", err)
	}
	newFleet := models.Fleet{
		Name:      "Frota teste 1",
		Max_Speed: 12.34,
	}
	savedFleet, err := newFleet.SaveFleet(server.DB)
	if err != nil {
		t.Errorf("Error while saving a Fleet: %v\n", err)
		return
	}
	assert.Equal(t, uint32(1), savedFleet.ID)
}

func TestGetFleetByID(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error fleet refreshing table %v\n", err)
	}

	fleet, err := seedOneFleet()
	if err != nil {
		log.Fatalf("cannot seed fleet table: %v", err)
	}
	foundFleet, err := fleetInstance.FindFleetByID(server.DB, fleet.ID)
	if err != nil {
		t.Errorf("this is the error getting one Fleet: %v\n", err)
		return
	}
	assert.Equal(t, foundFleet.ID, fleet.ID)
	assert.Equal(t, foundFleet.Name, fleet.Name)
	assert.Equal(t, foundFleet.Max_Speed, fleet.Max_Speed)
}
