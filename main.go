package main

import (
	"database/sql"
	"fmt"

	"entity"
	"repository"
	"service"
	animals "zoo/animals"

	_ "github.com/go-sql-driver/mysql"
)

const X = 1

func plus(x, y int) int {
	return x + y
}

type Animal struct {
	name string
	age  uint64
}

type AnimalInterface interface {
	gero()
	introduction(name string)
}

func (animal Animal) gero() {
	fmt.Println("gero")
}

func (animal Animal) introduction(name string) {
	fmt.Println("私の名前は")
	fmt.Println(name)
}

func one() uint64 {
	return 1
}

func returnFunc() func() {
	return func() {
		fmt.Println("yes!!!")
	}
}

func getUser(ID uint64) {

	db := sql.Open("mysql", "root:password@tcp(127:0.0.1:3306)/users")

	var name string

	if err := db.QueryRow("SELECT name FROM user WHERE id=? LIMIT 1", ID).Scan(&name); err != nil {
		panic(err)
	}

	defer db.Close()
}

func main() {
	returnFunc()()
	f := func(x, y int) int { return x + y }
	yes := f(2, 103)
	fmt.Println(yes)

	fmt.Println(plus(1, 2))
	a := one()
	fmt.Printf("%T\n", a)
	fmt.Printf("ahan=%T\n", "ahan")
	r := "\u65e5本\U00008a9e"
	fmt.Printf("%v", r)
	// fmt.Println(AppName())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

	takagi := Animal{"バーミヤン", 64}
	person := NewAnimal("高木", 60)

	fusigi(takagi)
	fusigi(person)

	entity.Ahan()

	ur := repository.NewUserRepository(300)
	us := service.NewUserService(ur)
	fmt.Println(us.GetUser(1))
	// fmt.Println(us.ur.GetUser())

}

func NewAnimal(name string, age uint64) Animal {
	return Animal{name, age}
}

func fusigi(animal Animal) {
	animal.introduction(animal.name)
	animal.gero()
}
