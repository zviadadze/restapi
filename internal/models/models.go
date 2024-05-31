package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func (u *User) Encode(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, "unable to encode user", http.StatusInternalServerError)
		return
	}
}
