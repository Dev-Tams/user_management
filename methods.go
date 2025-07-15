package main

import (
	"fmt"
	"slices"
)



func (u User) FullContact() string{
	return fmt.Sprintf("%v: %v, :%v ", u.ID, u.Name, u.Email)
} 

func (u User) Greet() string{
	return fmt.Sprintf("Hello %v!", u.Name)
}

func (u User) LoginStatus(){
	//return a login status later
}



func (a Admin) HasPrivilege(privi string) bool{
	return slices.Contains(a.Privileges, privi)
}

func (e Editor) CanEditSection(edit string) bool{
	return e.CanEdit
}


func (v Viewer) ViewAccess() string{
	return fmt.Sprintf("Access granted for %v", v.Name)
}

// . Create a Role Checker Function
// A function that accepts any type (interface{} or any)

// Uses type switches or type assertions to determine and print what role a user has


func RoleChecker(role any){
	switch r := role.(type){
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