package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/Samito19/msu-mp-auth-ms/user"
)

func main() {
	functionsHandler()
	http.ListenAndServe(":3001", nil)
}

func functionsHandler() {
	http.HandleFunc("/authenticate", handleAuthentication)
}

func handleAuthentication(w http.ResponseWriter, req *http.Request) {
	var user User
	credentials, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(credentials, &user)

}
