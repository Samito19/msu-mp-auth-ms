package main

import (
	"net/http"

	. "github.com/Samito19/msu-mp-auth-ms/errorHandlers"
	"github.com/Samito19/msu-mp-auth-ms/routes"
)

func main() {
	functionsHandler()
	http.ListenAndServe(":3001", nil)
}

func functionsHandler() {
	http.HandleFunc("/createUser", MakeRouteHandler(routes.CreateUser))
	http.HandleFunc("/authenticateUser", MakeRouteHandler(routes.AuthenticateUser))

}
