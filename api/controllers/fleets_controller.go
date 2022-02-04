package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/arthur2weber/go_rest/api/responses"
	"github.com/arthur2weber/go_rest/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) Createfleet(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	fleet := models.Fleet{}
	err = json.Unmarshal(body, &fleet)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fleet.Prepare()
	err = fleet.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fleetCreated, err := fleet.SaveFleet(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, fleetCreated.ID))
	responses.JSON(w, http.StatusCreated, map[string]int{"id": int(fleetCreated.ID)})
}

func (server *Server) GetFleets(w http.ResponseWriter, r *http.Request) {

	fleet := models.Fleet{}

	fleets, err := fleet.FindAllFleet(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fleets)
}

func (server *Server) GetFleet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["fleet_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fleet := models.Fleet{}
	fleetGotten, err := fleet.FindFleetByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, fleetGotten)
}
