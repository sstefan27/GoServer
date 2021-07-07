package rest

import (
	"encoding/json"
	"fmt"
	"go/problem2/entity"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostVehicle(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var vehicle entity.Vehicle
	err = json.Unmarshal(bodyBytes, &vehicle)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	fmt.Println(vehicle)
	rw.Write(bodyBytes)
}

func GetVehicle(rw http.ResponseWriter, r *http.Request) {
	typeOfVehicle := r.URL.Query().Get("type")
	var vehicleType entity.Vehicle
	switch typeOfVehicle {
	case "car":
		vehicleType = entity.Car{
			Name:  "BMW",
			Type:  "car",
			Model: "2020",
		}
	case "bike":
		vehicleType = entity.Bike{
			Name:  "Mountain Bike",
			Type:  "bike",
			Model: "2013",
		}
	case "bus":
		vehicleType = entity.Bus{
			Name:  "Volvo",
			Type:  "bus",
			Model: "1992",
		}
	}
	val, _ := json.Marshal(vehicleType)
	rw.Write(val)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}
