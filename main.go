package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println(editor.CanEditSection("email"))

	fmt.Println(user, "\n", user.Greet(), "\n", admin, "\n", admin.Greet())

	//range over users to check role
	for _, u := range users {
		RoleChecker(u)
	}
	//range over users to check for error in mail
	for _, v := range users {
		_, err := IsValidEmail(v.Email)
		if err != nil {
			fmt.Printf("Wrong email for %v \n", v.Email)
		} else {
			fmt.Println("checked mail for", v.Email)
		}
	}

	//check users email
	// _, err := user.IsValidEmail(user.Email)
	// if err != nil {
	// 	fmt.Println("Wrong email", err)
	// } else {
	// 	fmt.Println("checked mail for", user)
	// }
	//  RoleChecker(admin.User)

	// fmt.Println(c)
	AddtoDb()
	log, err := Login("tami@mail.com", "password")
	if err != nil {
		var LogError LoginError
		if errors.As(err, &LogError) {
			fmt.Println("LoginError:", LogError)
		} else {
			fmt.Println("Generic error", err)
		}
	} else {
		println(log)
	}

	fileU, err := os.Create("user.json")
	fileA, err2 := os.Create("admin.json")
	fileE, err3 := os.Create("editor.json")
	if err != nil {
		fmt.Println(" error creating json file")
		return
	} else {
		fmt.Println(" created user file")
	}

	defer fileU.Close()
	defer fileA.Close()
	defer fileE.Close()


	encoding := json.NewEncoder(fileU)
	encoding2 := json.NewEncoder(fileA)
	encoding3 := json.NewEncoder(fileE)

	
	err = encoding.Encode(users)
	err2 = encoding2.Encode(admin)
	err3 = encoding3.Encode(editors)
		if err != nil || err2  != nil|| err3 != nil{
			fmt.Println(" Error encoding json file")
			return
		}else{
			fmt.Println(" Json file encoded successfully")
		}
}
