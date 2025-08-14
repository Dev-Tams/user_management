package user

import (
	"encoding/json"
	"fmt"
	"os"
)

var Users = []User{
	{Name: "James", Email: "james4pf@mail.com", Password: "securepassword", ID: 8, Role: "admin"},
	{Name: "John", Email: "johndoe@mail.com", Password: "securepassword", ID: 6},
	{Name: "Doe", Email: "doe@mail.com", Password: "securepassword", ID: 5},
	{Name: "Mary", Email: "marydoe@mail.com", Password: "securepassword", ID: 9},
	{Name: "Tesla", Email: "teslamail.com", Password: "securepassword", ID: 7},
}

var Admins = []Admin{
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
			Section: []string{"email", "blog", "substack"},
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

var Editors = []Editor{

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

var U_user = User{
	Name: "Tami", Email: "tami@mail.com", Password: "password", ID: 4,
}

var E_editor = Editor{
	User: User{
		Name: "Editor", Email: "editor@mail.com", ID: 2,
	},
	CanEdit: true,
	Section: []string{"Substack", "Blog"},
}

var A_admin = Admin{
	User: User{
		Name: "Admin", Email: "admin@mail.com", Password: "adminpass", ID: 2,
	},
	Privileges: []string{"read", "write", "delete"},
}

func WriteToJson(filename string, data any) (string, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(data)
	encoder.SetIndent("", "  ")
	if err != nil {
		return "", err
	}
	return fmt.Sprintln("JSON written to file."), nil

}

func ReadFromJSon(filename string, data any) (string, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	decode := json.NewDecoder(f)
	err = decode.Decode(&data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("User loaded from JSON: %+v\n, user", data), nil
}

func BasicMarsh(data any) (string, error) {
	_, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintln(string(jsonData)), nil
}


func BasicUnMarsh(data []byte, datas any) error{
	return json.Unmarshal(data, datas)
}