package main

import (
	"integra_backend/internal/app"
	"integra_backend/internal/controller"
	"integra_backend/internal/db"
	"integra_backend/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	//Validator
	e.Validator = &CustomValidator{validator: validator.New()}
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	//DB
	dbConn, _ := db.NewDbConnection()
	defer dbConn.CloseConnection()
	//Model
	model := model.NewUserModel(dbConn)
	//Controller
	controller := controller.NewUserController(model)
	//app
	application := app.NewApp(controller)
	application.ConfigRoutes(e)
	//Starting server
	e.Logger.Fatal(e.Start(":1323"))
}
