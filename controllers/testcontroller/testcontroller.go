package testcontroller

import (
	"net/http"

	"github.com/FireCrackerTeam/wonderjack-web/helper"
)

func Index(w http.ResponseWriter, r *http.Request){
	data := []map[string]interface{}{
		{
			"id": 1,
			"nama": "test nama",
		},
		{
			"id": 2,
			"nama": "test nama 2",
		},
	}
	
	helper.ResponseJSON(w, http.StatusOK, data)
}