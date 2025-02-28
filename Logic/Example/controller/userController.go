package controller

import (
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/service"
)

// Class dan atribut
type UserController struct {
	UserService service.UserService
}

// Constructor
func NewUserController(userService service.UserService) UserController {

	// Mereturn Object Usercontroller , sehingga bisa panggil method
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) FindAllUserController() []dto.UserResponse {

	return uc.UserService.FindAllUser()
}

func (uc *UserController) AddUserController(userReq dto.UserRequest) dto.UserResponse {
	return uc.UserService.AddUserService(userReq)
}
