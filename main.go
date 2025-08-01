package main

import (
	"errors"
	"fmt"
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

	mar, err := BasicMarsh(users)
	if err != nil {
		fmt.Println("Error with marshalling", err)
	} else {
		fmt.Println(mar)
	}
	fileU, err := WriteToJson("user.json", users)
	fileA, err2 := WriteToJson("admin.json", admins)
	fileE, err3 := WriteToJson("editor.json", editors)

	if err != nil || err2 != nil || err3 != nil {
		fmt.Println(" Error writing json file")
		return
	} else {
		fmt.Println(&fileA, &fileE, &fileU, " written to Json successfully")
	}

	fileU, err = ReadFromJSon("user.json", users)
	if err != nil {
		fmt.Println(" Error reading from json file", err)
		return
	} else {
		fmt.Println(fileU)
	}

	pro, err := BasicMarsh(products)
	if err != nil {
		fmt.Println("error with slice", err)
	} else {
		fmt.Println(pro)
	}

	WriteToJson("products.json", products)
	ReadFromJSon("products.json", products)
	BasicUnMarsh([]byte(pro), &products)

}
