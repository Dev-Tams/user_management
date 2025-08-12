package main

import (
	// "errors"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Dev-Tams/user_management/fake_http"
	"github.com/Dev-Tams/user_management/user"
)

func main() {

	//lets spin a server for offline use
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/", request.GetUserHandler)
	mux.HandleFunc("POST /users/", request.PostUserHandler)

	fmt.Println("server starting on :8000...")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(user.E_editor.CanEditSection("email"))

	// fmt.Println(user.U_user, "\n", user.U_user.Greet(), "\n", user.A_admin, "\n", user.A_admin.Greet())
	fmt.Println(user.U_user.Greet())
	// //range over users to check role
	// for _, u := range user.Users {
	// 	user.RoleChecker(u)
	// }
	//range over users to check for error in mail
	// for _, v := range user.Users {
	// 	_, err := user.IsValidEmail(v.Email)
	// 	if err != nil {
	// 		fmt.Printf("Wrong email for %v \n", v.Email)
	// 	} else {
	// 		fmt.Println("checked mail for", v.Email)
	// 	}
	// }

	//check users email
	// _, err := user.IsValidEmail(user.Email)
	// if err != nil {
	// 	fmt.Println("Wrong email", err)
	// } else {
	// 	fmt.Println("checked mail for", user)
	// }
	//  RoleChecker(admin.User)

	// fmt.Println(c)
	// user.AddtoDb()
	// log, err := user.Login("tami@mail.com", "password")
	// if err != nil {
	// 	var LogError user.LoginError
	// 	if errors.As(err, &LogError) {
	// 		fmt.Println("LoginError:", LogError)
	// 	} else {
	// 		fmt.Println("Generic error", err)
	// 	}
	// } else {
	// 	println(log)
	// }

	// mar, err := user.BasicMarsh(user.Users)
	// if err != nil {
	// 	fmt.Println("Error with marshalling", err)
	// } else {
	// 	fmt.Println(mar)
	// }
	// fileU, err := user.WriteToJson("user.json", user.Users)
	// fileA, err2 := user.WriteToJson("admin.json", user.Admins)
	// fileE, err3 := user.WriteToJson("editor.json", user.Editors)

	// if err != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(" Error writing json file")
	// 	return
	// } else {
	// 	fmt.Println(&fileA, &fileE, &fileU, " written to Json successfully")
	// }

	// fileU, err = user.ReadFromJSon("user.json", user.Users)
	// if err != nil {
	// 	fmt.Println(" Error reading from json file", err)
	// 	return
	// } else {
	// 	fmt.Println(fileU)
	// }

	// pro, err := user.BasicMarsh(user.Products)
	// if err != nil {
	// 	fmt.Println("error with slice", err)
	// } else {
	// 	fmt.Println(pro)
	// }

	// user.WriteToJson("products.json", user.Products)
	// user.ReadFromJSon("products.json", user.Products)
	// user.BasicUnMarsh([]byte(pro), &user.Products)

	// var loaded []user.Product
	// if _, err := user.ReadFromJSon("products.json", &loaded); err != nil {
	// 	fmt.Println("error reading from file:", err)
	// 	return
	// }
	// fmt.Printf("Loaded from file: %+v\n", loaded)

	// Also unmarshal from the earlier string
	// var again []user.Product
	// if err := user.BasicUnMarsh([]byte(pro), &again); err != nil {
	// 	fmt.Println("error unmarshalling from string:", err)
	// 	return
	// }
	// fmt.Printf("Unmarshalled from string: %+v\n", again)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := request.GetReq(ctx, "http://localhost:8000/users/")
	if err != nil {
		fmt.Println("error fetching resource:", err)
		return
	}

	fmt.Println(resp)
	var todo user.Todo
	if err := user.BasicUnMarsh([]byte(resp), &todo); err != nil {
		fmt.Println("error unmarshalling from string:", err)
		return
	}
	fmt.Printf("Unmarshalled from string: %+v\n", todo)

	p := user.Post{
		Title:  "hello",
		Body:   "world",
		UserID: 1,
	}

	resp1, err := request.PostReq(ctx, "https://jsonplaceholder.typicode.com/posts", p)
	if err != nil {
		fmt.Println("error fetching resource:", err)
		return
	}

	b, _ := json.MarshalIndent(resp1, " ", "")
	fmt.Println(string(b))
}
