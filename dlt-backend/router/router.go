package router

import (
	"database/sql"
	"net/http"

	"github.com/robwatsongtr/dsa-leetcode-tracker/handlers"
	"github.com/robwatsongtr/dsa-leetcode-tracker/utils"
)

func SetupRoutes(db *sql.DB) {
	http.HandleFunc("/problems", problemRouter(db))
	//http.HandleFunc("/clients/", clientIDRouter(db))
}

func problemRouter(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			utils.DisplayLog(handlers.GetAllProblemsHandler(db))(w, r)
		// case http.MethodPost:
		// 	utils.DisplayLog(handlers.CreateClientHandler(db))(w, r)
		// case http.MethodPut:
		// 	utils.DisplayLog(handlers.UpdateClientHandler(db))(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// func clientIDRouter(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodDelete:
// 			utils.DisplayLog(handlers.DeleteClientHandler(db))(w, r)
// 		default:
// 			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		}
// 	}
// }
