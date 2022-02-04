package controllers

import "github.com/arthur2weber/go_rest/api/middlewares"

func (s *Server) initializeRoutes() {

	// drop
	s.Router.HandleFunc("/database", middlewares.SetMiddlewareJSON(s.RefreshDB)).Methods("DELETE")

	// Fleet
	s.Router.HandleFunc("/fleets", middlewares.SetMiddlewareJSON(s.Createfleet)).Methods("POST")
	s.Router.HandleFunc("/fleets", middlewares.SetMiddlewareJSON(s.GetFleets)).Methods("GET")

	// Fleets Alerts
	s.Router.HandleFunc("/fleets/{id:[0-9]+}/alerts", middlewares.SetMiddlewareJSON(s.CreateFleetAlert)).Methods("POST")
	s.Router.HandleFunc("/fleets/{id:[0-9]+}/alerts", middlewares.SetMiddlewareJSON(s.GetFleetAlerts)).Methods("GET")

	// Vehicle
	s.Router.HandleFunc("/vehicles", middlewares.SetMiddlewareJSON(s.CreateVehicle)).Methods("POST")
	s.Router.HandleFunc("/vehicles", middlewares.SetMiddlewareJSON(s.GetVehicles)).Methods("GET")

	// Vehicle Position
	s.Router.HandleFunc("/vehicles/{id:[0-9]+}/positions", middlewares.SetMiddlewareJSON(s.CreateVehiclePosition)).Methods("POST")
	s.Router.HandleFunc("/vehicles/{id:[0-9]+}/positions", middlewares.SetMiddlewareJSON(s.GetVehiclePositions)).Methods("GET")

}
