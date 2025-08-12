package request

import (
	"encoding/json"
	"net/http"
	"github.com/Dev-Tams/user_management/user"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]any{
		"users":   user.Users,
		"message": "All users",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    defer r.Body.Close()

   	 var newUser user.User 
    if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

 
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

