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
