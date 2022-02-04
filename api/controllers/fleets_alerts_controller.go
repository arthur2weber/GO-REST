package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/arthur2weber/go_rest/api/responses"
	"github.com/arthur2weber/go_rest/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateFleetAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Fleet_ID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fleet := models.Fleet{}
	_, err = fleet.FindFleetByID(server.DB, uint32(Fleet_ID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Fleet not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	fleetAlert := models.FleetAlert{}
	err = json.Unmarshal(body, &fleetAlert)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fleetAlert.Fleet_ID = uint32(Fleet_ID)
	fleetAlert.Prepare()

	err = fleetAlert.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fleetAlertCreated, err := fleetAlert.SaveFleetAlert(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, fleetAlertCreated.ID))
	responses.JSON(w, http.StatusCreated, map[string]int{"id": int(fleetAlertCreated.ID)})
}

func (server *Server) GetFleetAlerts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	Fleet_ID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	fleetAlert := models.FleetAlert{}

	alerts, err := fleetAlert.FindFleetAlertsByFleetID(server.DB, uint32(Fleet_ID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, alerts)
}

func (server *Server) GetFleetAlert(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fleetAlert := models.FleetAlert{}
	fleetGotten, err := fleetAlert.FindFleetAlertsByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, fleetGotten)
}
