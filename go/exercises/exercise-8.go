package exercises

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Exercise8() {
	r := mux.NewRouter()
	r.HandleFunc("/users", UsersHandler).Methods("GET")
	r.HandleFunc("/user", CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

var (
	users = []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
	}
	mu sync.Mutex
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	users = append(users, user)
	log.Printf("User created: %+v\n", user)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if id != updatedUser.ID {
		http.Error(w, "User ID mismatch", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, user := range users {
		if user.ID == id {
			users[i] = updatedUser
			log.Printf("User updated: %+v\n", updatedUser)
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
				http.Error(w, "Failed to encode user", http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			log.Printf("User deleted: %d\n", id)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
