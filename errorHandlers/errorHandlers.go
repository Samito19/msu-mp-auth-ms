package errorhandlers

import (
	"fmt"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}
}

func MakeRouteHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Error : %v", err)
			}
		}()
		function(w, r)
	}
}
