package main

type User struct{
	Name, Email, Password string
	ID int
}

type Admin struct{
 	User
	Privileges [] string
}

type Editor struct{
	User
	CanEdit bool
	Section [] string
}

type Viewer struct{
	User
	AccessLevel int
}