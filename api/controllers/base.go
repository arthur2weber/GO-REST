package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthur2weber/go_rest/api/models"
	"github.com/arthur2weber/go_rest/api/responses"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to postgres database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the postgres database")
	}

	server.DB.AutoMigrate()

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) RefreshDB(w http.ResponseWriter, r *http.Request) {

	err := server.DB.DropTableIfExists(&models.VehiclePosition{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.DropTableIfExists(&models.Vehicle{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.DropTableIfExists(&models.FleetAlert{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.DropTableIfExists(&models.Fleet{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.Debug().AutoMigrate(&models.Fleet{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.Debug().AutoMigrate(&models.FleetAlert{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.Debug().AutoMigrate(&models.Vehicle{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	err = server.DB.Debug().AutoMigrate(&models.VehiclePosition{}).Error
	if err != nil {
		responses.JSON(w, http.StatusNotModified, err)
		return
	}

	log.Printf("Successfully refreshed database")
	responses.JSON(w, http.StatusOK, "")
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
