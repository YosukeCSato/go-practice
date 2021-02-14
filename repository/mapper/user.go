package mapper

import "github.com/YosukeCSato/go-practice/entity"

type User struct {
	ID    uint64 `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func NewUser() *User {
	return &User{}
}

func ToEntity(u *User) *entity.User {
	return &entity.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
