package controllers

import "models"

type UserController struct {
	model *models.User
}

var userController UserController = UserController{}

func (UserController) GetUser(req *Request, res *Response) {
	c.JSON(200, "")
}
