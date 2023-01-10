package errorhandlers

import (
	"log"
	"net/http"
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
		function(w, r)
	}
}
