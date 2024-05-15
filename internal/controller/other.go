package controller

// import (
// 	"fmt"
// 	"integra_backend/internal/entity"
// 	"integra_backend/internal/model"
// 	"integra_backend/internal/utils"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

type otherController struct {
}

// type UserController interface {
// 	CreateUser(echo.Context) error
// 	ListUsers(echo.Context) error
// 	UpdateUser(c echo.Context) error
// }

// func NewUserController(userModel model.UserModel) UserController {
// 	return &userController{
// 		userModel: userModel,
// 	}
// }

// func (uc *userController) CreateUser(c echo.Context) error {
// 	newUser := entity.UserEntity{}
// 	customMsg := "error bind data"
// 	component := "[UserController.CreateUser]"
// 	//Binding and validating
// 	user, err := uc.validateUser(c, &newUser, component, customMsg)
// 	if err != nil {
// 		return err
// 	}
// 	//Creating user
// 	pgUser, err := uc.userModel.CreateUser(user)
// 	if err != nil {
// 		customMsg = "error to insert data in Postgres DB"
// 		component = "[UserModel.CreateUser]"
// 		return uc.handleError(c, component, customMsg, err)
// 	}
// 	return utils.HandleSuccess(c, pgUser)
// }

// func (uc *userController) ListUsers(c echo.Context) error {
// 	customMsg := "error to list data in Postgres DB"
// 	component := "[UserModel.ListUsers]"
// 	pgUsers, err := uc.userModel.ListUsers()
// 	if err != nil {
// 		return uc.handleError(c, component, customMsg, err)
// 	}
// 	return utils.HandleSuccess(c, pgUsers)
// }

// func (uc *userController) UpdateUser(c echo.Context) error {
// 	newUser := entity.UserEntity{}
// 	customMsg := "error bind data"
// 	component := "[UserController.UpdateUser]"
// 	//Binding and validating
// 	user, err := uc.validateUser(c, &newUser, component, customMsg)
// 	if err != nil {
// 		return err
// 	}
// 	//Updating user
// 	id, err := strconv.Atoi(c.Param("user_id"))
// 	if err != nil {
// 		return uc.handleError(c, component, customMsg, err)
// 	}
// 	user_id := int64(id)
// 	pgUser, err := uc.userModel.UpdateUser(user, user_id)
// 	if err != nil {
// 		customMsg = "error to update data in Postgres DB"
// 		component = "[UserModel.UpdateUser]"
// 		return uc.handleError(c, component, customMsg, err)
// 	}
// 	return utils.HandleSuccess(c, pgUser)
// }

// //Private Functions

// func (uc *userController) handleError(c echo.Context, component string, customMsg string, err error) error {
// 	msg := fmt.Sprintf("%v %v %v \n", component, customMsg, err.Error())
// 	fmt.Println(msg)
// 	return utils.HandleError(c, http.StatusInternalServerError, msg)
// }

// func (uc *userController) validateUser(c echo.Context, newUser *entity.UserEntity, component string, customMsg string) (*entity.UserEntity, error) {
// 	//Binding new user
// 	err := c.Bind(&newUser)
// 	if err != nil {
// 		return nil, uc.handleError(c, component, customMsg, err)
// 	}
// 	//Validating fields
// 	err = c.Validate(&newUser)
// 	if err != nil {
// 		return nil, uc.handleError(c, component, customMsg, err)
// 	}
// 	return newUser, nil
// }
