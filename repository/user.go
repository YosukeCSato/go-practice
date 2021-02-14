package repository

import (
	"database/sql"
	"fmt"

	"github.com/YosukeCSato/go-practice/entity"
	"github.com/YosukeCSato/go-practice/repository/mapper"
)

type UserRepository interface {
	GetUser(uint64) uint64
}

type userRepository struct {
	ID uint64
}

type User struct {
	ID    uint64
	Name  string
	Email string
}

func NewUserRepository(ID uint64) *userRepository {
	return &userRepository{ID}
}

func (ur userRepository) GetUser() *entity.User {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/users")
	if err != nil {
		return nil
	}

	user := mapper.NewUser()

	const query = `
	  SELECT
		  id,
			name,
			email
		FROM
		  users
	  WHERE
		  id = ?
		LIMIT 1
	`

	if err := db.QueryRow(query, ur.ID).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		panic(err)
	}

	fmt.Printf("%vです!!!!!(ID: %v, email: %v)\n", user.Name, user.ID, user.Email)

	defer db.Close()
	return mapper.ToEntity(user)
}
