package main

import "fmt"

func main() {
	

	user := User{
		Name: "Tami", Email: "tami@mail.com", ID: 4,
	}

	admin := Admin{
		User :User{
			Name: "Admin", Email: "admin@mail.com", ID: 2,
			}, 
		Privileges: []string{"read", "write", "delete"},

	}
	fmt.Println(user, "\n", user.Greet(), "\n", admin, "\n", admin.Greet())
}