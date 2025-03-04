package service

import (
	"fmt"
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/entity"
	"latihan-golang/Logic/Example/repo"
	"time"
)

// kita bikin Object
type UserService struct {
	UserRepository *repo.UserRepository
}

/*
Kenapa pakai userRepository , karena constructor nya
mereturn type data struct UserRepository
*/
// Ini Constructor nya
func NewUserService(userRepository *repo.UserRepository) *UserService {

	// Mereturn UserService
	return &UserService{

		UserRepository: userRepository,
	}
}

// ini method nya

func (us *UserService) FindAllUser() []dto.UserResponse {

	findAll := us.UserRepository.FindAllUser()
	listUser := []dto.UserResponse{}

	for _, val := range findAll {
		mapping := dto.UserResponse{
			Id:        val.Id,
			UserName:  val.UserName,
			Age:       val.Age,
			Address:   val.Address,
			CreatedAt: time.Now(),
		}
		listUser = append(listUser, mapping)
	}
	return listUser
}

func (us *UserService) AddUserService(userReq dto.UserRequest) dto.UserResponse {

	mappingUser := entity.User{
		UserName: userReq.UserName,
		Age:      userReq.Age,
		Address:  userReq.Address,
	}

	saveUser := us.UserRepository.AddUserRepo(mappingUser)

	return dto.UserResponse{
		Id:        saveUser.Id,
		UserName:  saveUser.UserName,
		Age:       saveUser.Age,
		Address:   saveUser.Address,
		CreatedAt: time.Now(),
	}
}

func (us *UserService) FindByIdService(id int) (entity.User, bool) {

	findData, cek := us.UserRepository.FindById(id)
	fmt.Println("Cek hasil : ", findData)
	if cek {
		return entity.User{
			Id:       findData.Id,
			UserName: findData.UserName,
			Age:      findData.Age,
			Address:  findData.Address,
		}, true
	}
	return entity.User{}, false
}
