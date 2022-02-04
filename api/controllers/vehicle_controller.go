package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/arthur2weber/go_rest/api/responses"
	"github.com/arthur2weber/go_rest/api/utils/formaterror"
)

func (server *Server) CreateVehicle(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	vehicle := models.Vehicle{}
	err = json.Unmarshal(body, &vehicle)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vehicle.Prepare()

	err = vehicle.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vehicleCreated, err := vehicle.SaveVehicle(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, vehicleCreated.ID))
	responses.JSON(w, http.StatusCreated, map[string]int{"id": int(vehicleCreated.ID)})
}

func (server *Server) GetVehicles(w http.ResponseWriter, r *http.Request) {

	Vehicle := models.Vehicle{}

	alerts, err := Vehicle.FindAllVehicle(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, alerts)
}
