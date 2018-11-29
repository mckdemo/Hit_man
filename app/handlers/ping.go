package handlers

import (
	"net/http"
)

func Ping(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("ok"))
}
