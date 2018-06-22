package structStudy

func main(){
	u := new(User)
	u.Name = "kook"
	u.Age = 36
	u.Create(u)
}