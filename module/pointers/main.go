package main

import "fmt"

type User struct {
	Name string
}

func modify(u *User) {
	// *u = User{Name: "Good"}
	(*u).Name = "ldfj"
	// u.Name = "Good"
}

func main() {
	u := &User{Name: "Hello"}
	fmt.Printf("u before change: %#v\n", u)
	modify(u)
	fmt.Printf("u after changed: %#v\n", u)
}
