package main

import "fmt"

type User struct {
	email string
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func main() {
	var user = User{email: "muhirwa232@gmail.com"}
	user.updateEmail("testing@gmail.com")
	fmt.Println(user.Email())
}
