package main

import "fmt"

type User struct {
	ID int
	Name string
	Email string
}

func (u *User) Rename (name string) {
	u.Name = name
}

func (u User) Label() string {
	return fmt.Sprintf("#%d %s", u.ID, u.Name)
}

func main() {
	u := User{ID: 1, Name: "Reza"}
	u.Rename("Reza Wijaya")
	fmt.Println(u.Label())
}
