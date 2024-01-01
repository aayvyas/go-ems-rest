package handlers

import (
	"encoding/json"
	"go-rest-api/internal/tools"
	"net/http"
)


func GetLogs(w http.ResponseWriter,r *http.Request) {
	var db *tools.LogDb = tools.GetLogDB()
	var logs tools.LogDatabase = (*db).GetLogs()	
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(logs)
	w.WriteHeader(http.StatusOK)

}