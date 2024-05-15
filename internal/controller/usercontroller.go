package controller

import (
	"fmt"
	"integra_backend/internal/entity"
	"integra_backend/internal/model"
	"integra_backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userModel model.UserModel
}

type UserController interface {
	CreateUser(echo.Context) error
	ListUsers(echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

func NewUserController(userModel model.UserModel) UserController {
	return &userController{
		userModel: userModel,
	}
}

func (uc *userController) CreateUser(c echo.Context) error {
	user := entity.UserEntity{}
	customMsg := "error bind data"
	component := "[UserController.CreateUser]"
	//Binding new user
	err := c.Bind(&user)
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	//Validating fields
	err = c.Validate(&user)
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	//Creating user
	pgUser, err := uc.userModel.CreateUser(&user)
	if err != nil {
		customMsg = "error to insert data in Postgres DB"
		component = "[UserModel.CreateUser]"
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUser)
}

func (uc *userController) ListUsers(c echo.Context) error {
	customMsg := "error to list data in Postgres DB"
	component := "[UserModel.ListUsers]"
	pgUsers, err := uc.userModel.ListUsers()
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUsers)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	user := entity.UserEntity{}
	customMsg := "error bind data"
	component := "[UserController.UpdateUser]"
	//Binding new user
	err := c.Bind(&user)
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	//Validating fields
	err = c.Validate(&user)
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	//Updating user
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	user_id := int64(id)
	pgUser, err := uc.userModel.UpdateUser(&user, user_id)
	if err != nil {
		customMsg = "error to update data in Postgres DB"
		component = "[UserModel.UpdateUser]"
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUser)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	//Deleting user
	customMsg := "error bind data"
	component := "[UserController.DeleteUser]"
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	user_id := int64(id)
	pgUser, err := uc.userModel.DeleteUser(user_id)
	if err != nil {
		customMsg = "error to delete data in Postgres DB"
		component = "[UserModel.DeleteUser]"
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUser)
}

//Private Functions

func (uc *userController) handleError(c echo.Context, component string, customMsg string, err error) error {
	fmt.Println()
	msg := fmt.Sprintf("%v %v %v \n", component, customMsg, err.Error())
	fmt.Println(msg)
	return utils.HandleError(c, http.StatusInternalServerError, msg)
}
