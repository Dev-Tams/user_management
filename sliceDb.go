package main


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






	var user =  User{
		Name: "Tami", Email: "tami@mail.com", Password: "password", ID: 4,
	}

	var editor = Editor{
		User: User{
			Name: "Editor", Email: "editor@mail.com", ID: 2,
		},
		CanEdit: true,
		Section: []string{"Substack", "Blog"},
	}

	var admin = Admin{
		User: User{
			Name: "Admin", Email: "admin@mail.com", Password: "adminpass", ID: 2,
		},
		Privileges: []string{"read", "write", "delete"},
	}