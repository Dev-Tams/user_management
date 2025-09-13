package main

import (
	// "errors"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	request "github.com/Dev-Tams/user_management/fake_http"
	"github.com/Dev-Tams/user_management/migration"
)

var db *sql.DB
var err error

func main() {

	db, err = sql.Open("sqlite3", "usersdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migration before using the DB
	migration.InitDB(db)

	// Now pass a live DB to the handler
	handler := request.Handler{DB: db}

	mux := http.NewServeMux()
	logger := request.Logger(mux)

	mux.HandleFunc("GET /users/", handler.GetUser)
	mux.HandleFunc("GET /users/{id}", handler.GetUserById)
	mux.HandleFunc("POST /users/", handler.PostUser)
	mux.HandleFunc("PUT /users/{id}", handler.PutUser)
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)

	fmt.Println("server started on :8000...")
	if err := http.ListenAndServe(":8000", logger); err != nil {
		log.Fatal(err)
	}

	// time.Sleep(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := request.GetReq(ctx, "http://localhost:8000/users/")
	if err != nil {
		fmt.Println("error fetching resource:", err)
		return
	}
	fmt.Println("Response:", resp)
}
