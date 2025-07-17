package main

import "fmt"

func main() {
	

	user := User{
		Name: "Tami", Email: "tami@mail.com", ID: 4,
	}

	editor := Editor{
		User :User{
			Name: "Editor", Email: "editor@mail.com", ID: 2,
			}, 
			CanEdit: true,
			Section: []string{"Substack", "Blog"},
	}

	fmt.Println(editor.CanEditSection("email"))
	
	admin := Admin{
		User :User{
			Name: "Admin", Email: "admin@mail.com", ID: 2,
			}, 
		Privileges: []string{"read", "write", "delete"},

	}
	fmt.Println(user, "\n", user.Greet(), "\n", admin, "\n", admin.Greet())


		users := []User{
		{Name: "Tami", Email: "tami@mail.com", ID: 3},
		{Name: "James", Email: "james4pf@mail.com", ID: 8},
		{Name: "John", Email: "johndoe@mail.com", ID: 6},
		{Name: "Doe", Email: "Doe@mail.com", ID: 5},
		{Name: "Mary", Email: "marydoe@mail.com", ID: 9},
		{Name: "Tesla", Email: "teslamail.com", ID: 7},
	}

//range over users to check role	
	for _, u := range users{
		 RoleChecker(u)
	}

	for _, v := range users{
		 _, err := IsValidEmail(v)
		 if err != nil{
			fmt.Printf("Wrong email for %v \n", v.Email)
		 }else{
			fmt.Println("checked mail for", v.Email)
		 }
	}
//check users email
	 _, err:= user.IsValidEmail(user.Email)
	if err != nil{
		fmt.Println("Wrong email", err)
	}else{
		fmt.Println("checked mail for", user)
	}
	//  RoleChecker(admin.User)
	 
	// fmt.Println(c)
}
