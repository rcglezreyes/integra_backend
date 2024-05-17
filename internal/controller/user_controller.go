package controller

import (
	"fmt"
	"integra_backend/internal/entity"
	"integra_backend/internal/message"
	"integra_backend/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userModel model.UserModel
}

type UserController interface {
	CreateUser(echo.Context) (*entity.UserEntity, error)
	ListUsers(echo.Context) ([]*entity.UserEntity, error)
	UpdateUser(c echo.Context) (*entity.UserEntity, error)
	DeleteUser(c echo.Context) (*entity.UserEntity, error)
}

func NewUserController(userModel model.UserModel) UserController {
	return &userController{
		userModel: userModel,
	}
}

func (uc *userController) CreateUser(c echo.Context) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentCreateUserController
	//Binding new user
	err := c.Bind(&user)
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	//Validating fields
	err = c.Validate(&user)
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	//Creating user
	pgUser, err := uc.userModel.CreateUser(&user)
	if err != nil {
		customMsg = message.MsgResponseInsertDataPgDatabaseError
		component = message.MsgComponentCreateUserModel
		return nil, uc.handleError(component, customMsg, err)
	}
	return pgUser, nil
}

func (uc *userController) ListUsers(c echo.Context) ([]*entity.UserEntity, error) {
	customMsg := message.MsgResponseListDataPgDatabaseError
	component := message.MsgComponentListUserModel
	pgUsers, err := uc.userModel.ListUsers()
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	return pgUsers, nil
}

func (uc *userController) UpdateUser(c echo.Context) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentUpdateUserController
	//Binding new user
	err := c.Bind(&user)
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	//Validating fields
	err = c.Validate(&user)
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	//Updating user
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	user_id := int64(id)
	pgUser, err := uc.userModel.UpdateUser(&user, user_id)
	if err != nil {
		customMsg = message.MsgResponseUpdateDataPgDatabaseError
		component = message.MsgComponentUpdateUserModel
		return nil, uc.handleError(component, customMsg, err)
	}
	return pgUser, nil
}

func (uc *userController) DeleteUser(c echo.Context) (*entity.UserEntity, error) {
	//Deleting user
	customMsg := message.MsgResponseBindDataError
	component := message.MsgComponentDeleteUserController
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return nil, uc.handleError(component, customMsg, err)
	}
	user_id := int64(id)
	pgUser, err := uc.userModel.DeleteUser(user_id)
	if err != nil {
		customMsg = message.MsgResponseDeleteDataPgDatabaseError
		component = message.MsgComponentDeleteUserModel
		return nil, uc.handleError(component, customMsg, err)
	}
	return pgUser, nil
}

//Private Functions

func (uc *userController) handleError(component string, customMsg string, err error) error {
	fmt.Println()
	msg := fmt.Sprintf("%v %v %v \n", component, customMsg, err.Error())
	fmt.Println(msg)
	return err
}
