package handlers

import (
	"database/sql"
	//"encoding/json"
	"net/http"
	//"strconv"
	//"strings"

	"github.com/robwatsongtr/dsa-leetcode-tracker/models"
	"github.com/robwatsongtr/dsa-leetcode-tracker/utils"
)

/*
Takes in a database connection and returns the handler. Functional, higher order handling.
Both http.ResponseWriter and *http.Request are required for HTTP handlers
even through the request is not being used. like req res in express I suppose.
*/

// reads in the whole client table from postgres and passes it along to the server as json
func GetAllProblemsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := models.GetProblems(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, clients)
	}
}
