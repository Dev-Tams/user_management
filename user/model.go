package user

type User struct{
	Name, Email, Password string
	ID int
}

type Admin struct{
	Editor
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

type Todo struct{
	UserId int	`json:"userId"`
	Id int			`json:"id"`
	TItle string 	`json:"title"`
	Completed bool `json:"completed"`
}

type Post struct{
	 Title string `json:"title"`
    Body   string `json:"body"`
    UserID 	int	   `json:"userId"`
}
type CreatePost struct{
	// ID     int    `json:"id"` 
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID int    `json:"userId"`
}