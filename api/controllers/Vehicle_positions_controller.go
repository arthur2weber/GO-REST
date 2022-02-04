package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/arthur2weber/go_rest/api/responses"
	"github.com/arthur2weber/go_rest/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateVehiclePosition(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Vehicle_ID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	vehicle := models.Vehicle{}
	_, err = vehicle.FindVehicleByID(server.DB, uint32(Vehicle_ID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Vehicle not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	VehiclePosition := models.VehiclePosition{}
	err = json.Unmarshal(body, &VehiclePosition)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	VehiclePosition.Vehicle_ID = uint32(Vehicle_ID)

	VehiclePosition.Prepare()

	err = VehiclePosition.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	VehiclePositionCreated, err := VehiclePosition.SaveVehiclePosition(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	// start calls
	if VehiclePosition.Current_Speed > vehicle.Max_Speed {
		go server.CallFleetAlerts(int32(vehicle.Fleet_ID), VehiclePositionCreated)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, VehiclePositionCreated.ID))
	responses.JSON(w, http.StatusCreated, map[string]int{"id": int(VehiclePositionCreated.ID)})
}

func (server *Server) CallFleetAlerts(Fleet_id int32, body *models.VehiclePosition) {
	FleetAlert := models.FleetAlert{}
	urls := []string{}
	urls, _ = FleetAlert.FindWebhooksByFleetID(server.DB, uint32(Fleet_id))

	for _, webhook := range urls {
		go server.CallUrl(webhook, body)
	}
}

func (server *Server) CallUrl(url string, body *models.VehiclePosition) {
	jsonValue, _ := json.Marshal(body)
	for _, n := range [4]int{0, 1, 5, 15} {
		time.Sleep(time.Duration(n) * time.Second)

		resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if resp.StatusCode == http.StatusOK {
			return
		}
	}
}

func (server *Server) GetVehiclePositions(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Vehicle_ID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	VehiclePosition := models.VehiclePosition{}

	alerts, err := VehiclePosition.FindVehiclePositionByVehicleId(server.DB, uint32(Vehicle_ID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, alerts)
}
