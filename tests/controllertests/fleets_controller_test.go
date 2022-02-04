package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateFleet(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		errorMessage string
		ID           int
	}{
		{
			inputJSON:  `{"Name":"Frota teste 1", "Max_Speed": 12.34}`,
			statusCode: http.StatusCreated,
			ID:         1,
		},
		{
			inputJSON:    `{"Name":"Frota sem speed"}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "Required Max_Speed",
			ID:           2,
		},
		{
			inputJSON:    `{ "Max_Speed": 88}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "Required Name",
			ID:           3,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/fleets", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Createfleet)
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

func TestGetFleets(t *testing.T) {

	err := RefreshTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedFleets()
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/fleets", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetFleet)
	handler.ServeHTTP(rr, req)

	var fleets []models.Fleet
	err = json.Unmarshal([]byte(rr.Body.String()), &fleets)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(fleets), 2)
}
