package modeltests

import (
	"log"
	"testing"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllFleetAlerts(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error refreshing fleet_alerts table %v\n", err)
	}

	_, err = seedFleetsAlerts()
	if err != nil {
		log.Fatalf("Error seeding fleet_alerts table %v\n", err)
	}

	fleetsAlerts, err := fleetAlertInstance.FindAllFleetAlert(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the fleet_alerts: %v\n", err)
		return
	}
	assert.Equal(t, len(*fleetsAlerts), 2)
}

func TestSaveFleetAlerts(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error fleet_alerts refreshing table %v\n", err)
	}
	newFleetAlert := models.FleetAlert{
		Fleet_ID: 1,
		Webhook:  "http://test.com",
	}
	savedAlert, err := newFleetAlert.SaveFleetAlert(server.DB)
	if err != nil {
		t.Errorf("Error while saving a fleet_alerts: %v\n", err)
		return
	}
	assert.Equal(t, uint32(1), savedAlert.ID)
}

func TestGetFleetAlertByFleetID(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatalf("Error fleet_alerts refreshing table %v\n", err)
	}

	alert, err := seedOneFleetsAlert()
	if err != nil {
		log.Fatalf("cannot seed fleet_alerts table: %v", err)
	}
	foundAlert, err := fleetAlertInstance.FindFleetAlertsByID(server.DB, alert.ID)
	if err != nil {
		t.Errorf("this is the error getting one fleet_alerts: %v\n", err)
		return
	}
	assert.Equal(t, foundAlert.ID, alert.ID)
	assert.Equal(t, foundAlert.Fleet_ID, alert.Fleet_ID)
	assert.Equal(t, foundAlert.Webhook, alert.Webhook)
}
