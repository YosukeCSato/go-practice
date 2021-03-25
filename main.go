package main

import (
	"fmt"

	"github.com/YosukeCSato/go-practice/repository"
	_ "github.com/go-sql-driver/mysql"
)

func returnFunc() func() {
	return func() {
		fmt.Println("yes!!!")
		fmt.Println("yes")
	}
}

const (
	X = 1 + iota
	Y
	Z
	S1 = "あ"
	S2
)

func main() {
	ur := repository.NewUserRepository(1)
	user := ur.GetUser()
	fmt.Println(user.ID)
	f := later()
	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome!"))
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
	fmt.Println(S1)
	fmt.Println(S2)
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
