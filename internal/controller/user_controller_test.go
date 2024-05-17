package controller_test

import (
	cfg "integra_backend/internal/config"
	"integra_backend/internal/controller"
	cv "integra_backend/internal/custom_validator"
	"integra_backend/internal/db"
	"integra_backend/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var e *echo.Echo
var dbConn db.DbConnection
var m model.UserModel
var c controller.UserController
var lastId int64

// var application app.App

func TestUserController(t *testing.T) {
	RegisterFailHandler(Fail)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// adjust it
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	// pass it in to RunSpecs
	RunSpecs(t, "Testing User Controller")
}

var _ = BeforeSuite(func() {
	cfg.ConfigEnv()
	e = echo.New()
	//Validator
	e.Validator = cv.NewCustomValidator(validator.New())
	//DB
	host, port, user, pwd, dbname, _ := cfg.DBCredentials()
	dbConn, _ = db.NewDbConnection(host, port, user, pwd, dbname)
	//Model
	m = model.NewUserModel(dbConn)
	//Controller
	c = controller.NewUserController(m)
	//app
	// application = app.NewApp(c)
	// application.ConfigRoutes(e)
})

var _ = Describe("Testing User Controller", Label("User Controller"), func() {

	When("User Controller Creates New User", func() {

		Context("When successfully creating a new user", func() {
			It("should return the correct response", func() {
				requestBody := `{
					"user_name": "user_test_controller",
					"first_name": "User Test",
					"last_name": "Ginkgo Gomega",
					"email": "user_test_controller@ginkgo.go",
					"user_status": "T",
					"department": "Developing"
				}`

				req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(requestBody))
				rec := httptest.NewRecorder()
				req.Header.Add("Accept", "application/json;charset=utf-8")
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.CreateUser(ctx)
				if err != nil {
					Expect(err.Error()).To(Equal("userName in use"))
				} else {
					Expect(err).NotTo(HaveOccurred())
				}
			})
		})

	})

	When("User Controller Lists Users", func() {

		Context("When successfully list users", func() {
			It("should return the correct response", func() {
				req := httptest.NewRequest(http.MethodPost, "/users", nil)
				rec := httptest.NewRecorder()
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.ListUsers(ctx)
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	When("User Controller Updates User", func() {

		Context("When unsuccessfully updates user", func() {
			It("should return the incorrect response (Bad integer value of user ID)", func() {
				requestBody := `{
					"user_name": "user_test_controller",
					"first_name": "User Test",
					"last_name": "Ginkgo Gomega",
					"email": "user_test_controller@ginkgo.go",
					"user_status": "T",
					"department": "Developing"
				}`
				req := httptest.NewRequest(http.MethodPost, "/update_user/bad_id", strings.NewReader(requestBody))
				rec := httptest.NewRecorder()
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.UpdateUser(ctx)
				Expect(err).NotTo(BeNil())
			})

			It("should return the incorrect response (User ID does not exist)", func() {
				requestBody := `{
					"user_name": "user_test_controller",
					"first_name": "User Test",
					"last_name": "Ginkgo Gomega",
					"email": "user_test_controller@ginkgo.go",
					"user_status": "T",
					"department": "Developing"
				}`
				req := httptest.NewRequest(http.MethodPost, "/update_user/999999999", strings.NewReader(requestBody))
				rec := httptest.NewRecorder()
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.UpdateUser(ctx)
				Expect(err).NotTo(BeNil())
			})
		})

	})

	When("User Controller Deletes User", func() {

		Context("When unsuccessfully deletes user", func() {
			It("should return the incorrect response (Bad integer value of user ID)", func() {
				req := httptest.NewRequest(http.MethodPost, "/delete_user/bad_id", nil)
				rec := httptest.NewRecorder()
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.DeleteUser(ctx)
				Expect(err).NotTo(BeNil())
			})

			It("should return the incorrect response (User ID does not exist)", func() {
				req := httptest.NewRequest(http.MethodPost, "/delete_user/999999999", nil)
				rec := httptest.NewRecorder()
				req.Header.Add("Content-Type", "application/json")
				ctx := e.NewContext(req, rec)
				_, err := c.DeleteUser(ctx)
				Expect(err).NotTo(BeNil())
			})
		})

	})

})
