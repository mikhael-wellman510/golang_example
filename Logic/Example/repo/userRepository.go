package repo

import (
	"fmt"
	"latihan-golang/Logic/Example/entity"
)

type UserRepository struct {
	//Private final user
	User []entity.User
}

func NewUserRepository() *UserRepository {

	// Buat instance
	return &UserRepository{
		User: []entity.User{
			{Id: 1, UserName: "Mikhael", Age: 25},
			{Id: 2, UserName: "Deni", Age: 30},
		},
	}
}

func (ur *UserRepository) FindAllUser() []entity.User {

	return ur.User
}

func (ur *UserRepository) AddUserRepo(user entity.User) entity.User {

	idx := len(ur.User) + 1
	saveUser := entity.User{
		Id:       idx,
		UserName: user.UserName,
		Age:      user.Age,
		Address:  user.Address,
	}

	ur.User = append(ur.User, saveUser)

	return saveUser
}

func (ur *UserRepository) FindById(user int) (entity.User, bool) {
	fmt.Println("Ids : ", user)
	fmt.Println("tes ", ur.User)
	for _, val := range ur.User {
		if val.Id == user {
			return val, true
		}
	}

	return entity.User{}, false
}
