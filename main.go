package main

import (
	"errors"
	"fmt"
)

var users = []User{
	{Name: "James", Email: "james4pf@mail.com", Password: "securepassword", ID: 8},
	{Name: "John", Email: "johndoe@mail.com", Password: "securepassword", ID: 6},
	{Name: "Doe", Email: "doe@mail.com", Password: "securepassword", ID: 5},
	{Name: "Mary", Email: "marydoe@mail.com", Password: "securepassword", ID: 9},
	{Name: "Tesla", Email: "teslamail.com", Password: "securepassword", ID: 7},
}

var admins = []Admin{
	{
		User: User{
			Name:     "Tami",
			Email:    "tami@mail.com",
			Password: "password",
			ID:       4,
		},
		Privileges: []string{"read", "write", "delete"},
		Editor: Editor{
			CanEdit: true,
			Section: []string{ "email", "blog", "substack"},
		},
	},
	{
		User: User{
			Name:     "James",
			Email:    "james4pf@mail.com",
			Password: "securepassword",
			ID:       8,
		},
		Privileges: []string{"read", "write"},
	},
}

var editors = []Editor{

	{
		User: User{
			Name: "Editor", Email: "editor@mail.com", Password: "password", ID: 9,
		},
		CanEdit: true,
		Section: []string{"Substack", "Blog"},
	},
	{
		User: User{
			Name:     "Editor2",
			Email:    "editor2@mail.com",
			Password: "password",
			ID:       13,
		},
		CanEdit: true,
		Section: []string{"Substack", "Blog"},
	},
	{
		User: User{
			Name:     "Editor3",
			Email:    "editor3@mail.com",
			Password: "securepassword",
			ID:       8,
		},
		CanEdit: true,
		Section: []string{"blog"},
	},
}



func main() {

	user := User{
		Name: "Tami", Email: "tami@mail.com", Password: "password", ID: 4,
	}

	editor := Editor{
		User: User{
			Name: "Editor", Email: "editor@mail.com", ID: 2,
		},
		CanEdit: true,
		Section: []string{"Substack", "Blog"},
	}

	fmt.Println(editor.CanEditSection("email"))

	admin := Admin{
		User: User{
			Name: "Admin", Email: "admin@mail.com", Password: "adminpass", ID: 2,
		},
		Privileges: []string{"read", "write", "delete"},
	}
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
	}else{
		println(log)
	}
}
