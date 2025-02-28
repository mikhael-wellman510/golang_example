package api

import (
	"fmt"
	"latihan-golang/Logic/Example/controller"
	"latihan-golang/Logic/Example/dto"
)

func HitUser(userController controller.UserController) {

	// Find All product
	findAllUser1 := userController.FindAllUserController()
	fmt.Println("data user awal :", findAllUser1)

	// Add Product
	addUser := dto.UserRequest{UserName: "Alfredo", Age: 30, Address: "Bogor"}
	userController.AddUserController(addUser)
	findAllUser2 := userController.FindAllUserController()
	fmt.Println("data user new : ", findAllUser2)

	// Find By Id
}
