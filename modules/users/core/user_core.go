package user_core

import (
	user_model "main/modules/users/model"
	user_service "main/modules/users/service"
	base_controller "main/pkg/base/controller"
	base_service "main/pkg/base/service"
)

var Service = base_service.Init[user_model.Struct](&user_service.UserService{})
var Controller = base_controller.NewController(*Service)
