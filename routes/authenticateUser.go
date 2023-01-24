package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Samito19/msu-mp-auth-ms/database"
	. "github.com/Samito19/msu-mp-auth-ms/database/models"
	"github.com/Samito19/msu-mp-auth-ms/encryption"
	. "github.com/Samito19/msu-mp-auth-ms/errorHandlers"
	. "github.com/Samito19/msu-mp-auth-ms/jwt"
)

type credentials struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

func AuthenticateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		//Receive user credentials
		var receivedCredentials credentials
		credentialsBody, reqBodyErr := io.ReadAll(req.Body)
		CheckError(reqBodyErr)
		json.Unmarshal(credentialsBody, &receivedCredentials)

		//Pull user data from the database for verification
		var user User
		db := database.CreateConnection()
		db.Where("email_address = ?", receivedCredentials.EmailAddress).First(&user)

		//Verify user credentials
		if encryption.CheckPassword(receivedCredentials.Password, user.Password) {
			log.Printf("%s, logged in.\n\r", user.EmailAddress)
			newToken, tokenErr := CreateNewToken()
			CheckError(tokenErr)
			accessTokenJson, jsonMshlErr := json.Marshal(map[string]string{"auth_status": "success", "accessToken": newToken})
			CheckError(jsonMshlErr)
			w.Write(accessTokenJson)
		} else {
			authStatus, jsonMshlErr := json.Marshal(map[string]string{"auth_status": "failed"})
			CheckError(jsonMshlErr)
			w.Write(authStatus)
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(req.Method + " method is " + http.StatusText(http.StatusNotImplemented)))
	}
}
