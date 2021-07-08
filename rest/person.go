package rest

import (
	"encoding/json"
	"fmt"
	"go/problem2/db"
	"go/problem2/entity"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostPerson(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError1(rw, err, "Internal Issue") {
		return
	}

	var person entity.Person
	err = json.Unmarshal(bodyBytes, &person)
	if hasError1(rw, err, "Internal Issue") {
		return
	}

	db.GetDB().Create(&person)
	fmt.Println(person)
	rw.Write(bodyBytes)
}

func GetPerson(rw http.ResponseWriter, r *http.Request) {
	idPerson := r.URL.Query().Get("id")
	var person entity.Person
	result := db.GetDB().Where("id= ?", idPerson).Find(&person)
	if result.RecordNotFound() {
		http.Error(rw, " Record not foundd", http.StatusInternalServerError)
		return
	}
	//if result.Error
	personData, _ := json.Marshal(person)
	rw.Write([]byte(personData))
}

func hasError1(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}
