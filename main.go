package main

import (
	"fmt"
	"integra_backend/internal/app"
	"integra_backend/internal/config"
	"integra_backend/internal/controller"
	cv "integra_backend/internal/custom_validator"
	"integra_backend/internal/db"
	"integra_backend/internal/model"
	"os"

	cfg "integra_backend/internal/config"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	if err, isConfigurable := config.ConfigEnv(); !isConfigurable {
		log.Error(err.Error())
		os.Exit(1)
	} else {
		//DB
		host, port, user, pwd, dbname, err := cfg.DBCredentials()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		e := echo.New()
		//Validator
		e.Validator = cv.NewCustomValidator(validator.New())
		// Middleware
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))
		e.Use(middleware.Recover())
		e.Use(middleware.CORS())

		dbConn, _ := db.NewDbConnection(host, port, user, pwd, dbname)
		defer dbConn.CloseConnection()
		//Model
		model := model.NewUserModel(dbConn)
		//Controller
		controller := controller.NewUserController(model)
		//app
		application := app.NewApp(controller)
		application.ConfigRoutes(e)
		//Starting server
		server := fmt.Sprintf(":%v", viper.GetString(cfg.APP_PORT))
		e.Logger.Fatal(e.Start(server))
	}
}
