// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/gorethink/gorethink.v4"
	"gopkg.in/gorethink/gorethink.v4/internal/compare"
)

// These tests test the type of command
func TestDatumTypeofSuite(t *testing.T) {
	suite.Run(t, new(DatumTypeofSuite))
}

type DatumTypeofSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumTypeofSuite) SetupTest() {
	suite.T().Log("Setting up DatumTypeofSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_dtype").Exec(suite.session)
	err = r.DBCreate("db_dtype").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_dtype").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *DatumTypeofSuite) TearDownSuite() {
	suite.T().Log("Tearing down DatumTypeofSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_dtype").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *DatumTypeofSuite) TestCases() {
	suite.T().Log("Running DatumTypeofSuite: These tests test the type of command")

	{
		// datum/typeof.yaml line #5
		/* 'NULL' */
		var expected_ string = "NULL"
		/* r.expr(null).type_of() */

		suite.T().Log("About to run line #5: r.Expr(nil).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr(nil).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #5")
	}

	{
		// datum/typeof.yaml line #9
		/* 'NULL' */
		var expected_ string = "NULL"
		/* r.type_of(null) */

		suite.T().Log("About to run line #9: r.TypeOf(nil)")

		runAndAssert(suite.Suite, expected_, r.TypeOf(nil), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}
}
