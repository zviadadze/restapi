package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zviadadze/userver/internal/models"
	"github.com/zviadadze/userver/internal/storage"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handleGetUsers)
	mux.HandleFunc("GET /users/{id}", handleGetUser)
	mux.HandleFunc("POST /users", handleAppendUser)
	mux.HandleFunc("DELETE /users/{id}", handleRemoveUser)
	return mux
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(storage.GetUsers()); err != nil {
		http.Error(w, "unable to encode users", http.StatusInternalServerError)
		return
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusNotFound)
		return
	}

	user, err := storage.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.Encode(w)
}

func handleAppendUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
		return
	}

	var dummyUser models.User
	err := json.NewDecoder(r.Body).Decode(&dummyUser)
	if err != nil {
		http.Error(w, "invalid user data", http.StatusBadRequest)
		return
	}

	user := storage.AppendUser(dummyUser.Name, dummyUser.Age)

	user.Encode(w)
}

func handleRemoveUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusNotFound)
		return
	}

	user, err := storage.RemoveUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.Encode(w)
}
