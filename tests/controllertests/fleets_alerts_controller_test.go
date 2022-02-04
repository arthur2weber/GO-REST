package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateFleetAlert(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}

	seedFleets()

	samples := []struct {
		inputJSON    string
		statusCode   int
		errorMessage string
		ID           int
		Fleet_ID     int
		Webhook      string
	}{
		{
			inputJSON:  `{"Webhook":"http://localhost:8081/fleet/alert"}`,
			statusCode: http.StatusCreated,
			Fleet_ID:   1,
			ID:         1,
		},
		{
			inputJSON:  `{"Webhook":"www.google.com"}`,
			statusCode: http.StatusBadRequest,
			Fleet_ID:   1,
		},
		{
			inputJSON:  `{"Webhook":"http://localhost:8081/fleet/alert"}`,
			statusCode: http.StatusBadRequest,
			Fleet_ID:   0,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/fleets", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(v.Fleet_ID)})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateFleetAlert)
		handler.ServeHTTP(rr, req)
		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, float64(v.ID), responseMap["id"])
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, v.errorMessage, responseMap["error"])
		}

	}
}

func TestGetFleetsAlerts(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedFleetsAlerts()
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/fleets/1/alerts", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetFleet)
	handler.ServeHTTP(rr, req)

	var fleetsAlerts []models.FleetAlert
	err = json.Unmarshal([]byte(rr.Body.String()), &fleetsAlerts)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(fleetsAlerts), 2)
}
