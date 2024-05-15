package db_test

import (
	"integra_backend/internal/db"
	"testing"

	cfg "integra_backend/internal/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var dbConn db.DbConnection

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// adjust it
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	// pass it in to RunSpecs
	RunSpecs(t, "Testing Db")
}

var _ = BeforeSuite(func() {
	cfg.ConfigEnv()
	//Config DB COnnection
	host, port, user, pwd, dbname, _ := cfg.DBCredentials()
	dbConn, _ = db.NewDbConnection(host, port, user, pwd, dbname)

})

var _ = Describe("Testing Db", Label("DB"), func() {

	When("New Connection", func() {

		Context("When successfully connect", func() {
			It("should return the correct response", func() {
				Expect(dbConn.GetConnection()).NotTo(BeNil())
			})
		})

	})

})
