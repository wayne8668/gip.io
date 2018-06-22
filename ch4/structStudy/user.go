package structStudy


type User struct{
	Name string
	Age int
}

func (user User) Create(u User){
	user.Name = u.Name
}