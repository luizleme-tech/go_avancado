package main

import (
	"encoding/json"
	"fmt"
	"goAvancado/user"
)

//func main() {
//	user := user.User{"Indiana Jones", 10}
//	fmt.Println(user.Name)
//	user.UpdateName("Henry Jones")
//	user.PrintName()
//	UpdateName(&user, "Professor Jones")
//	user.PrintName()
//}
//
//func UpdateName(u *user.User, newName string) {
//	u.Name = newName
//}

func mainStruct() {
	user := user.User{"Henry Jones Jr", 10}
	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
