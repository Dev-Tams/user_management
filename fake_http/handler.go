package request

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Dev-Tams/user_management/user"
)

type Handler struct{
	DB *sql.DB
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {

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
	for _, u := range user.Users {
		if u.Role == role {
			fUser = append(fUser, u)
		}
	}
	if fUser == nil {
		http.Error(w, "No User with role found", http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"users":   fUser,
		"message": fmt.Sprintf("Users with role '%s'", role),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func PostUser(h Handler) (w http.ResponseWriter, r *http.Request) {

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

	stmt, err := h.DB.Prepare("INSERT INTO users(name, email, password, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(newUser.Name, newUser.Email, newUser.Password, newUser.Role) 
	if err != nil{
		http.Error(w, "Insert failed", http.StatusInternalServerError)
	}
	
	id, _ := res.LastInsertId()
	newUser.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Inavlid user ID", http.StatusBadRequest)
		return
	}

	var fUser *user.User
	for _, u := range user.Users {
		if u.ID == id {
			fUser = &u
			break
		}
	}

	if fUser == nil {
		http.Error(w, "Can't find user", http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"user": fUser,
		"id":   fmt.Sprintf("user %v", id),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func PutUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	idParam := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var updateUser user.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var fUser *user.User
	for i := range user.Users {
		if user.Users[i].ID == id {
			fUser = &user.Users[i]
			break
		}
	}

	if fUser == nil {
		http.Error(w, "Can't find user", http.StatusBadRequest)
		return
	}

	fUser.Email = updateUser.Email
	fUser.Name = updateUser.Name
	fUser.Password = updateUser.Password
	fUser.Role = updateUser.Role

	response := map[string]any{
		"message": "User updated successfully",
		"user":    updateUser,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid Id", http.StatusBadRequest)
	}

	index := -1
	for i := range user.Users {
		if user.Users[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	user.Users = append(user.Users[:index], user.Users[index+1:]...)
	w.WriteHeader(http.StatusNoContent)

}
