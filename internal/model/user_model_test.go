package model_test

import (
	"integra_backend/internal/db"
	"integra_backend/internal/entity"
	"integra_backend/internal/model"
	"testing"

	cfg "integra_backend/internal/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var dbConn db.DbConnection
var um model.UserModel
var lastId int64

func TestUserModel(t *testing.T) {
	RegisterFailHandler(Fail)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// adjust it
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	// pass it in to RunSpecs
	RunSpecs(t, "Testing User Model")
}

var _ = BeforeSuite(func() {
	cfg.ConfigEnv()
	//Config DB COnnection
	host, port, user, pwd, dbname, _ := cfg.DBCredentials()
	dbConn, _ = db.NewDbConnection(host, port, user, pwd, dbname)
	//Creating the user model
	um = model.NewUserModel(dbConn)
	userFirst := &entity.UserEntity{
		UserName:   "first_user_test",
		FirstName:  "User Test I",
		LastName:   "Ginkgo Gomega I",
		Email:      "usertestgg65@gmail.com",
		UserStatus: "T",
		Department: "Developing",
	}
	um.CreateUser(userFirst)

})

var _ = Describe("Testing User Model", Label("User Model"), func() {

	When("User Model test List Users ", func() {

		Context("When successfully find all users", func() {
			It("should return the correct response", func() {
				u, err := um.ListUsers()
				Expect(err).NotTo(HaveOccurred())
				Expect(u).NotTo(BeNil())
			})
		})

	})

	When("User Model test Create User ", func() {

		Context("When successfully create", func() {
			It("should return the correct response", func() {

				userSecond := &entity.UserEntity{
					UserName:   "user_test",
					FirstName:  "User Test II",
					LastName:   "Ginkgo Gomega II",
					Email:      "usertestgg66@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				uSecond, errSecond := um.CreateUser(userSecond)
				lastId = userSecond.UserId
				Expect(errSecond).NotTo(HaveOccurred())
				Expect(uSecond).NotTo(BeNil())
			})
		})

		Context("When a field required is missing", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "user_test",
					FirstName:  "User Test",
					Email:      "usertestgg65@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.CreateUser(user)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

		Context("When a user name is in use", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "iuereyreir@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.CreateUser(user)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

		Context("When a email is in use", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "new_user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "usertestgg65@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.CreateUser(user)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

		Context("When a user status is longer than one character", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "my_last_create_user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "creatinglast222@gmail.com",
					UserStatus: "TT",
					Department: "Developing",
				}
				u, err := um.CreateUser(user)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

	})

	When("User Model test Update User ", func() {

		Context("When successfully update", func() {
			It("should return the correct response", func() {
				user := &entity.UserEntity{
					UserName:   "user_test",
					FirstName:  "User Test II",
					LastName:   "Ginkgo Gomega II",
					Email:      "usertestgg66@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.UpdateUser(user, lastId)
				lastId = u.UserId
				Expect(err).NotTo(HaveOccurred())
				Expect(u).NotTo(BeNil())
			})
		})

		Context("When a user name is in use", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "first_user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "iuereyreir@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.UpdateUser(user, lastId)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

		Context("When a email is in use", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "new_user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "usertestgg65@gmail.com",
					UserStatus: "T",
					Department: "Developing",
				}
				u, err := um.UpdateUser(user, lastId)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

		Context("When a cannot find the provided ID", func() {
			It("should return the incorrect response", func() {
				user := &entity.UserEntity{
					UserName:   "my_last_updated_user_test",
					FirstName:  "New User",
					LastName:   "Hello World",
					Email:      "creatinglast222@gmail.com",
					UserStatus: "TT",
					Department: "Developing",
				}
				u, err := um.UpdateUser(user, lastId+1)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

	})

	When("User Model test Delete User ", func() {

		Context("When successfully delete", func() {
			It("should return the correct response", func() {
				u, err := um.DeleteUser(lastId)
				Expect(err).NotTo(HaveOccurred())
				Expect(u).NotTo(BeNil())
			})
		})

		Context("When user id does not exist", func() {
			It("should return the incorrect response", func() {
				u, err := um.DeleteUser(lastId + 1)
				Expect(err).To(HaveOccurred())
				Expect(u).To(BeNil())
			})
		})

	})
})
