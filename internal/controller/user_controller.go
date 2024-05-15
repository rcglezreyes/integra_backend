package controller

import (
	"fmt"
	"integra_backend/internal/entity"
	"integra_backend/internal/message"
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
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentCreateUserController
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
		customMsg = message.MsgResponseInsertDataPgDatabaseError
		component = message.MsgComponentCreateUserModel
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUser)
}

func (uc *userController) ListUsers(c echo.Context) error {
	customMsg := message.MsgResponseListDataPgDatabaseError
	component := message.MsgComponentListUserModel
	pgUsers, err := uc.userModel.ListUsers()
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUsers)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	user := entity.UserEntity{}
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentUpdateUserController
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
		customMsg = message.MsgResponseUpdateDataPgDatabaseError
		component = message.MsgComponentUpdateUserModel
		return uc.handleError(c, component, customMsg, err)
	}
	return utils.HandleSuccess(c, pgUser)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	//Deleting user
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentDeleteUserController
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return uc.handleError(c, component, customMsg, err)
	}
	user_id := int64(id)
	pgUser, err := uc.userModel.DeleteUser(user_id)
	if err != nil {
		customMsg = message.MsgResponseDeleteDataPgDatabaseError
		component = message.MsgComponentDeleteUserModel
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
