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
		fmt.Println(" Welcome, user", r.Name)
	case Admin:
		fmt.Printf("Welcome, Admin %v! ", r.Name)
	case Editor:
		fmt.Println("Editor access granted to ", r.Name)
	case Viewer:
		fmt.Printf(" Hello %v , you're a viewer", r.Name)
	default:
		fmt.Println("Unknown role")
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


func  IsValidEmail(email string) (string, error) {
	x := len(email)

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return "", fmt.Errorf("email must contain '@' and '.'")
	}
	if x < 5 {
		return "", fmt.Errorf(" Email is not valid")
	} else {
		return email, nil
	}
}

//Checking all users and logging them in if cred matches
// Check for empty fields.
// Loop through all users to find a matching email.
func Login(email, password string) (string, error){

	email = strings.ToLower(email)
	if email == "" || password == "" {
		return "", LoginError{Reason: "Fields cannot be empty, try again!"}
	}

		//proceed to check valid email
	_, err := IsValidEmail(email)
		if err != nil {
			return "", LoginError{Reason: "Email is wrong"}
		}

	for _, user := range users{
		//check if cred matches in slice
		if email == user.Email{
			if password == user.Password{
				RoleChecker(user)
				return "login sucess", nil
			}else{
				return "", LoginError{Reason: "wrong email or password"}
			}
		}
	}
	return "",  LoginError{Reason: "Email not found"}

}