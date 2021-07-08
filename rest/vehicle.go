package rest

import (
	"encoding/json"
	"fmt"
	"go/problem2/entity"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

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
func PostVehicle(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var vehicle map[string]interface{}
	err = json.Unmarshal(bodyBytes, &vehicle)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	vehicleType, ok := vehicle["type"].(string)
	if !ok {
		return
	}
	switch vehicleType {
	case "car":
		fmt.Println("This is a CAR")
	case "bus":
		fmt.Println("This is a BUS")
	case "bike":
		fmt.Println("This is a BIKE")
	}
	fmt.Println(bodyBytes)
	rw.Write(bodyBytes)
}

func PostCanDrive(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var vehicle map[string]interface{}
	err = json.Unmarshal(bodyBytes, &vehicle)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	vehicleType, ok := vehicle["Type"].(string)
	if !ok {
		return
	}
	var canDrive bool
	switch vehicleType {
	case "car":
		var car entity.Car
		canDrive = car.CanDrive()
	case "bus":
		var bus entity.Bus
		canDrive = bus.CanDrive()
	case "bike":
		var bike entity.Bike
		canDrive = bike.CanDrive()
	}
	canDriveString := strconv.FormatBool(canDrive)
	rw.Write([]byte(canDriveString))
	//fmt.Println(bodyBytes)
	//rw.Write(bodyBytes)

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

//s
