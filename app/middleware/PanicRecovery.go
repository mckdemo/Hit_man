package middleware

import (
	"log"
	"net/http"
)

func PanicRecoveryHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				log.Printf("Encountered panic: %+v", err)
				http.Error(response, http.StatusText(500), 500)
			}

		}()

		next.ServeHTTP(response, request)

	})

}
