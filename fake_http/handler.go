package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Dev-Tams/user_management/user"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	  w.Header().Set("Content-Type", "application/json")


	
	role := r.URL.Query().Get("role")
	if role == "" {
		response := map[string]any{
		"users":   user.Users,
		"message": "All users",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
		
	}

	var fUser []user.User
	for _, u := range user.Users{
		if u.Role == role{
			 fUser = append(fUser, u)
		}
	}
	if fUser == nil{
			http.Error(w, "No User with role found",http.StatusBadRequest )
			return
		}
	

	  response := map[string]any{
        "users":   fUser,
        "message": fmt.Sprintf("Users with role '%s'", role),
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	  w.Header().Set("Content-Type", "application/json")


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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	  w.Header().Set("Content-Type", "application/json")


	idParam := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idParam)
	if err != nil {
        http.Error(w, "Inavlid user ID", http.StatusBadRequest)
		return
	}


	var fUser *user.User
	for _, u := range user.Users{
		 if u.ID == id {
            fUser = &u
            break
        }
	}

	if fUser == nil{
			http.Error(w, "Can't find user",http.StatusBadRequest )
			return
		}
	

	response := map[string]any{
		"user":  fUser,
		"id": fmt.Sprintf("user %v", id),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
