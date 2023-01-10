package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Samito19/msu-mp-auth-ms/database"
	"github.com/Samito19/msu-mp-auth-ms/encryption"
	"github.com/google/uuid"

	. "github.com/Samito19/msu-mp-auth-ms/database/models"
	. "github.com/Samito19/msu-mp-auth-ms/errorHandlers"
)

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Receive user data
	var user User
	userData, reqBodyErr := ioutil.ReadAll(req.Body)
	CheckError(reqBodyErr)
	json.Unmarshal(userData, &user)

	//generate UUID for the user
	id := uuid.New()
	user.Uuid = id

	//Encrypt user data
	hashedPassword, hashErr := encryption.HashPassword(user.Password)
	CheckError(hashErr)
	user.Password = hashedPassword

	//Save new user into the database
	db := database.CreateConnection()
	result := db.Create(&user)
	if result.Error != nil {
		response := make(map[string]string)
		response["message"] = "Failed to create new user !"
		jsonResponse, jsonErr := json.Marshal(response)
		CheckError(jsonErr)
		w.Write(jsonResponse)
		CheckError(result.Error)
	} else {
		response := make(map[string]string)
		response["message"] = "Successfully created new user !"
		jsonResponse, jsonErr := json.Marshal(response)
		CheckError(jsonErr)
		w.Write(jsonResponse)
	}
}
