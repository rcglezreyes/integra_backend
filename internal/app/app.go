package app

import (
	"integra_backend/internal/controller"
	"integra_backend/internal/message"
	"net/http"

	_ "integra_backend/internal/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type app struct {
	controller controller.UserController
}

type App interface {
	ConfigRoutes(*echo.Echo)
	CreateUser(echo.Context) error
	ListUsers(c echo.Context) error
}

func NewApp(controller controller.UserController) App {
	return &app{
		controller: controller,
	}
}

// @title Swagger Example CRUD Users API
// @version 1.0
// @description This is a sample CRUD Users API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (a *app) ConfigRoutes(e *echo.Echo) {

	// Routes
	e.POST("/user", a.CreateUser)
	e.GET("/users", a.ListUsers)
	e.PUT("/update_user/:user_id", a.UpdateUser)
	e.DELETE("/delete_user/:user_id", a.DeleteUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param parameters body entity.UserEntity true "ENTRY PAYLOAD"
// @Success 200 {object} map[string]interface{}
// @Router /user [post]
func (a *app) CreateUser(c echo.Context) error {
	err := a.controller.CreateUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": message.MsgResponseUserCreatedSuccess,
	})
}

// ListUsers func for lists all users.
// @Description List users.
// @Summary List users
// @Tags List Users
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func (a *app) ListUsers(c echo.Context) error {
	err := a.controller.ListUsers(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": message.MsgResponseUserListedSuccess,
	})
}

// UpdateUser func for updates an existing user.
// @Description Updates an existing user.
// @Summary updates an existing user
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path int true "USER ID"
// @Param parameters body entity.UserEntity true "ENTRY PAYLOAD"
// @Success 200 {object} map[string]interface{}
// @Router /update_user/{user_id} [put]
func (a *app) UpdateUser(c echo.Context) error {
	err := a.controller.UpdateUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": message.MsgResponseUserUpdatedSuccess,
	})
}

// DeleteUser func for an existing user.
// @Description Deletes an existing user.
// @Summary deletes an existing user
// @Tags User
// @Produce json
// @Param user_id path int true "USER ID"
// @Success 200 {object} map[string]interface{}
// @Router /delete_user/{user_id} [delete]
func (a *app) DeleteUser(c echo.Context) error {
	err := a.controller.DeleteUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": message.MsgResponseUserDeletedSuccess,
	})
}
