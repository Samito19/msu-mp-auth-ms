package errorhandlers

import (
	"log"
	"net/http"

	cors "github.com/Samito19/msu-mp-auth-ms/cors"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func MakeRouteHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		cors.Enable(w, r)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			function(w, r)
		}
	}
}
