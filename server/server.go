package server

import (
	"2gis/config"
	"log"
	"net/http"
)

func Init() {
	var configuration = config.GetConfig()
	var err = http.ListenAndServe(":" + configuration.Server.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/company/search/building/address", GetCompanyByBuildingAddress)
	http.HandleFunc("/company/search/building/coordinates", GetCompanyByBuildingCoordinates)
	http.HandleFunc("/company/search/rubric", GetCompanyByRubric)
}

func GetCompanyByBuildingAddress(w http.ResponseWriter, r *http.Request) {

}

func GetCompanyByBuildingCoordinates(w http.ResponseWriter, r *http.Request) {

}

func GetCompanyByRubric(w http.ResponseWriter, r *http.Request) {

}