package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Comment struct {
	Message string `json:"message"`
}

type Blog struct {
	Author   string    `json:"author"`
	Title    string    `json:"title"`
	Comments []Comment `json:"comments"`
}

var blog = Blog{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/blogs", PostBlog).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func PostBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewDecoder(r.Body).Decode(&blog)
	json.NewEncoder(w).Encode(&blog)
}
