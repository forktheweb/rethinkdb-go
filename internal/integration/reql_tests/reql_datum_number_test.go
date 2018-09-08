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

// Tests of conversion to and from the RQL number type
func TestDatumNumberSuite(t *testing.T) {
	suite.Run(t, new(DatumNumberSuite))
}

type DatumNumberSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumNumberSuite) SetupTest() {
	suite.T().Log("Setting up DatumNumberSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_dnum").Exec(suite.session)
	err = r.DBCreate("db_dnum").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_dnum").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *DatumNumberSuite) TearDownSuite() {
	suite.T().Log("Tearing down DatumNumberSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_dnum").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *DatumNumberSuite) TestCases() {
	suite.T().Log("Running DatumNumberSuite: Tests of conversion to and from the RQL number type")

	{
		// datum/number.yaml line #6
		/* 1 */
		var expected_ int = 1
		/* r.expr(1) */

		suite.T().Log("About to run line #6: r.Expr(1)")

		runAndAssert(suite.Suite, expected_, r.Expr(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// datum/number.yaml line #15
		/* -1 */
		var expected_ int = -1
		/* r.expr(-1) */

		suite.T().Log("About to run line #15: r.Expr(-1)")

		runAndAssert(suite.Suite, expected_, r.Expr(-1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// datum/number.yaml line #24
		/* 0 */
		var expected_ int = 0
		/* r.expr(0) */

		suite.T().Log("About to run line #24: r.Expr(0)")

		runAndAssert(suite.Suite, expected_, r.Expr(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// datum/number.yaml line #35
		/* 1.0 */
		var expected_ float64 = 1.0
		/* r.expr(1.0) */

		suite.T().Log("About to run line #35: r.Expr(1.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #35")
	}

	{
		// datum/number.yaml line #44
		/* 1.5 */
		var expected_ float64 = 1.5
		/* r.expr(1.5) */

		suite.T().Log("About to run line #44: r.Expr(1.5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #44")
	}

	{
		// datum/number.yaml line #53
		/* -0.5 */
		var expected_ float64 = -0.5
		/* r.expr(-0.5) */

		suite.T().Log("About to run line #53: r.Expr(-0.5)")

		runAndAssert(suite.Suite, expected_, r.Expr(-0.5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #53")
	}

	{
		// datum/number.yaml line #62
		/* 67498.89278 */
		var expected_ float64 = 67498.89278
		/* r.expr(67498.89278) */

		suite.T().Log("About to run line #62: r.Expr(67498.89278)")

		runAndAssert(suite.Suite, expected_, r.Expr(67498.89278), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #62")
	}

	{
		// datum/number.yaml line #73
		/* 1234567890 */
		var expected_ int = 1234567890
		/* r.expr(1234567890) */

		suite.T().Log("About to run line #73: r.Expr(1234567890)")

		runAndAssert(suite.Suite, expected_, r.Expr(1234567890), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #73")
	}

	{
		// datum/number.yaml line #83
		/* -73850380122423 */
		var expected_ int = -73850380122423
		/* r.expr(-73850380122423) */

		suite.T().Log("About to run line #83: r.Expr(-73850380122423)")

		runAndAssert(suite.Suite, expected_, r.Expr(-73850380122423), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #83")
	}

	{
		// datum/number.yaml line #95
		/* float(1234567890123456789012345678901234567890) */
		var expected_ float64 = float64(1234567890123456789012345678901234567890.0)
		/* r.expr(1234567890123456789012345678901234567890) */

		suite.T().Log("About to run line #95: r.Expr(1234567890123456789012345678901234567890.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(1234567890123456789012345678901234567890.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// datum/number.yaml line #100
		/* 123.4567890123456789012345678901234567890 */
		var expected_ float64 = 123.45678901234568
		/* r.expr(123.4567890123456789012345678901234567890) */

		suite.T().Log("About to run line #100: r.Expr(123.45678901234568)")

		runAndAssert(suite.Suite, expected_, r.Expr(123.45678901234568), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #100")
	}

	{
		// datum/number.yaml line #103
		/* 'NUMBER' */
		var expected_ string = "NUMBER"
		/* r.expr(1).type_of() */

		suite.T().Log("About to run line #103: r.Expr(1).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr(1).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #103")
	}

	{
		// datum/number.yaml line #107
		/* '1' */
		var expected_ string = "1"
		/* r.expr(1).coerce_to('string') */

		suite.T().Log("About to run line #107: r.Expr(1).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr(1).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #107")
	}

	{
		// datum/number.yaml line #110
		/* 1 */
		var expected_ int = 1
		/* r.expr(1).coerce_to('number') */

		suite.T().Log("About to run line #110: r.Expr(1).CoerceTo('number')")

		runAndAssert(suite.Suite, expected_, r.Expr(1).CoerceTo("number"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #110")
	}

	{
		// datum/number.yaml line #115
		/* int_cmp(1) */
		var expected_ int = int_cmp(1)
		/* r.expr(1.0) */

		suite.T().Log("About to run line #115: r.Expr(1.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #115")
	}

	{
		// datum/number.yaml line #119
		/* int_cmp(45) */
		var expected_ int = int_cmp(45)
		/* r.expr(45) */

		suite.T().Log("About to run line #119: r.Expr(45)")

		runAndAssert(suite.Suite, expected_, r.Expr(45), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #119")
	}

	{
		// datum/number.yaml line #123
		/* float_cmp(1.2) */
		var expected_ float64 = float_cmp(1.2)
		/* r.expr(1.2) */

		suite.T().Log("About to run line #123: r.Expr(1.2)")

		runAndAssert(suite.Suite, expected_, r.Expr(1.2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #123")
	}
}
