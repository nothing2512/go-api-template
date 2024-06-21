package controllers

import (
	"encoding/json"
	"net/http"
)

func (idb *InDb) Create(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    false,
		"message":   "",
		"data":      nil,
		"errorCode": "C202",
	})
}
