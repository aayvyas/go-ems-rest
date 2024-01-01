package middleware

import (
	"fmt"
	"go-rest-api/api"
	"go-rest-api/internal/tools"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

var ErrUnAuthorized = fmt.Errorf("INVALID USERNAME OR PASSWORD")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {	
		var idStr string  = r.URL.Query().Get("id")
		var token = r.Header.Get("Authorization")
		
		if idStr =="" ||  token == "" {
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}
		id, err := strconv.Atoi(idStr)
		if(err != nil) {
			logrus.Error("Not able to convert string to int for id", err, id)
		}
		var db *tools.DB = tools.GetDB()
		if(!(*db).CheckAuth(uint(id),token)) {
			api.RequestErrorHandler(w, ErrUnAuthorized)
			return
		}
		// call the next middleware in line
		next.ServeHTTP(w, r)

	})
}