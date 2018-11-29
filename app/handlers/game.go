package handlers

import (
	"log"
	"net/http"

	"github.com/mckdemo/hit_men/app/models"

	"github.com/mckdemo/hit_men/app/services"

	"github.com/gorilla/mux"
)

func CreateUser(connection *services.MySQLDatastore, user *models.User) error {

	transaction, errorMessage := connection.Begin()
	if errorMessage != nil {
		log.Print(errorMessage)
	}

	defer transaction.Rollback()

	preparedStatement, errorMessage := transaction.Prepare("INSERT INTO user(username, team, score) VALUES (?,?,?)")
	if errorMessage != nil {
		return errorMessage
	}

	defer preparedStatement.Close()

	_, errorMessage = preparedStatement.Exec(user.Username, user.Team, user.Score)
	if errorMessage != nil {
		return errorMessage
	}

	errorMessage = transaction.Commit()
	if errorMessage != nil {
		return errorMessage
	}

	return nil
}

func CreateUserAndAllocateTeam(connection *services.MySQLDatastore) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		username := vars["username"]
		newUser := models.NewUser(username)
		err := CreateUser(connection, newUser)
		if err != nil {
			log.Printf("Encountered panic: %+v", err)
			http.Error(response, http.StatusText(500), 500)
		}
		response.Write([]byte(username))
	})
}
