package repository

type UserRepository interface {
	GetUser(uint64) uint64
}

type userRepository struct {
	uID uint64
}

func NewUserRepository(uID uint64) UserRepository {
	return userRepository{uID}
}

func (ur userRepository) GetUser(uID uint64) uint64 {
	return uID + ur.uID
}
