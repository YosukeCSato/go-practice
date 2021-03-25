package main

import (
	"fmt"
	f "fmt"

	"github.com/YosukeCSato/go-practice/repository"
	_ "github.com/go-sql-driver/mysql"
)

func returnFunc() func() {
	return func() {
		f.Println("yes!!!")
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
	f.Println(user.ID)
	fu := later()
	f.Println(fu("Golang"))
	f.Println(fu("is"))
	f.Println(fu("awesome!"))
	f.Println(X)
	f.Println(Y)
	f.Println(Z)
	f.Println(S1)
	f.Println(S2)
	ger := doSomething(12)
	f.Println(ger)

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println("continue")
			continue // これ、上の"continue"はprintされないんだね…
		}
		fmt.Println(i)
		i++
	}
	var x interface{} = 3
	sss, isString := x.(string)
	fmt.Println(isString)
	fmt.Println(sss)

LOOP:
	for {
		for {
			for {
				fmt.Println("start")
				break LOOP
			}
		}
	}
	fmt.Println("完了")

}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func doSomething(a int) (b string) {
	{
		var a int
		const b = "string!!"
		f.Println(a)
		f.Println(b)
	}
	b = "jojojojooo"
	return b
}
