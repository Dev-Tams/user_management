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