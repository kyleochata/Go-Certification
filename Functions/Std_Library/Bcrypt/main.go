package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	login := "passwor123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Logged in")
}
