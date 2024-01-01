package middleware

import (
	"net/http"
	"time"

	"go-rest-api/internal/tools"

	"github.com/sirupsen/logrus"
)

type LogEntry struct{
	Timestamp string
	Query string
}


func LogRequest(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
		var logDb *tools.LogDb = tools.GetLogDB()
		(*logDb).LogRequestToDB(time.Now().String(), r.URL.String())

		logrus.Info("Logging the request to logDb")
		next.ServeHTTP(w, r)
	})



}