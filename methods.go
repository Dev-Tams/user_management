package main

import (
	// "errors"
	"fmt"
	"slices"
	"strings"
)

func (u User) FullContact() string {
	return fmt.Sprintf("%v: %v, :%v ", u.ID, u.Name, u.Email)
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello %v!", u.Name)
}

func (u User) LoginStatus() {
	//return a login status later
}

func (a Admin) HasPrivilege(privi string) bool {
	return slices.Contains(a.Privileges, privi)
}

func (e Editor) CanEditSection(edit string) (bool, error) {
	for _, ed := range e.Section {
		if ed == edit {
			return true, nil
		}
	}
	return false, fmt.Errorf(": Unauthorized section")
}

func (v Viewer) ViewAccess() string {
	return fmt.Sprintf("Access granted for %v", v.Name)
}


// Uses type switches or type assertions to determine and print what role a user has

func RoleChecker(role any) {

	switch r := role.(type) {
	case User:
		fmt.Println("User", r)
	case Admin:
		fmt.Println("Admin", r)
	case Editor:
		fmt.Println("Editor", r)
	case Viewer:
		fmt.Println("Viewer", r)
	}
}

//basic email validation
func (u User) IsValidEmail(e string) (string, error) {
	x := len(e)

	if !strings.Contains(e, "@") || !strings.Contains(e, ".") {
		return "", fmt.Errorf("email must contain '@' and '.'")
	}
	if x < 5 {
		return "", fmt.Errorf(" Email is not valid")
	} else {
		return e, nil
	}
}


func  IsValidEmail(e User) (string, error) {
	x := len(e.Email)

	if !strings.Contains(e.Email, "@") || !strings.Contains(e.Email, ".") {
		return "", fmt.Errorf("email must contain '@' and '.'")
	}
	if x < 5 {
		return "", fmt.Errorf(" Email is not valid")
	} else {
		return e.Email, nil
	}
}

//Checking all users and logging them in if cred matches
// Check for empty fields.
// Loop through all users to find a matching email.

func Login(email, password string) (string, error){

	if email == "" || password == "" {
		return "", LoginError{Reason: "Fields cannot be empty, try again!"}
	}

	for _, user := range users{
		if email == user.Email{
			if password == user.Password{
				return "login sucess", nil
			}else{
				return "", LoginError{Reason: "wrong email or password"}
			}
		}
		
	}
		return "",  LoginError{Reason: "Email not found"}
}