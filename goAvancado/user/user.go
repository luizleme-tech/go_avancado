package user

import "fmt"

type User struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}
