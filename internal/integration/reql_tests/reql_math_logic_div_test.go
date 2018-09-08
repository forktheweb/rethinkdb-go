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

// Tests for the basic usage of the division operation
func TestMathLogicDivSuite(t *testing.T) {
	suite.Run(t, new(MathLogicDivSuite))
}

type MathLogicDivSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicDivSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicDivSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_logic_div").Exec(suite.session)
	err = r.DBCreate("db_logic_div").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_logic_div").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicDivSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicDivSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_logic_div").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicDivSuite) TestCases() {
	suite.T().Log("Running MathLogicDivSuite: Tests for the basic usage of the division operation")

	{
		// math_logic/div.yaml line #6
		/* 2 */
		var expected_ int = 2
		/* r.expr(4) / 2 */

		suite.T().Log("About to run line #6: r.Expr(4).Div(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(4).Div(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// math_logic/div.yaml line #7
		/* 2 */
		var expected_ int = 2
		/* 4 / r.expr(2) */

		suite.T().Log("About to run line #7: r.Div(4, r.Expr(2))")

		runAndAssert(suite.Suite, expected_, r.Div(4, r.Expr(2)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// math_logic/div.yaml line #8
		/* 2 */
		var expected_ int = 2
		/* r.expr(4).div(2) */

		suite.T().Log("About to run line #8: r.Expr(4).Div(2)")

		runAndAssert(suite.Suite, expected_, r.Expr(4).Div(2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/div.yaml line #15
		/* 0.5 */
		var expected_ float64 = 0.5
		/* r.expr(-1) / -2 */

		suite.T().Log("About to run line #15: r.Expr(-1).Div(-2)")

		runAndAssert(suite.Suite, expected_, r.Expr(-1).Div(-2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #15")
	}

	{
		// math_logic/div.yaml line #20
		/* 4.9 / 0.7 */
		var expected_ float64 = 4.9 / 0.7
		/* r.expr(4.9) / 0.7 */

		suite.T().Log("About to run line #20: r.Expr(4.9).Div(0.7)")

		runAndAssert(suite.Suite, expected_, r.Expr(4.9).Div(0.7), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// math_logic/div.yaml line #25
		/* 1.0/120 */
		var expected_ float64 = 1.0 / 120
		/* r.expr(1).div(2,3,4,5) */

		suite.T().Log("About to run line #25: r.Expr(1).Div(2, 3, 4, 5)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Div(2, 3, 4, 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #25")
	}

	{
		// math_logic/div.yaml line #37
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(1) / 0 */

		suite.T().Log("About to run line #37: r.Expr(1).Div(0)")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Div(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// math_logic/div.yaml line #38
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(2.0) / 0 */

		suite.T().Log("About to run line #38: r.Expr(2.0).Div(0)")

		runAndAssert(suite.Suite, expected_, r.Expr(2.0).Div(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// math_logic/div.yaml line #39
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(3) / 0.0 */

		suite.T().Log("About to run line #39: r.Expr(3).Div(0.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(3).Div(0.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// math_logic/div.yaml line #40
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(4.0) / 0.0 */

		suite.T().Log("About to run line #40: r.Expr(4.0).Div(0.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(4.0).Div(0.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// math_logic/div.yaml line #41
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(0) / 0 */

		suite.T().Log("About to run line #41: r.Expr(0).Div(0)")

		runAndAssert(suite.Suite, expected_, r.Expr(0).Div(0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #41")
	}

	{
		// math_logic/div.yaml line #42
		/* err('ReqlQueryLogicError', 'Cannot divide by zero.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Cannot divide by zero.")
		/* r.expr(0.0) / 0.0 */

		suite.T().Log("About to run line #42: r.Expr(0.0).Div(0.0)")

		runAndAssert(suite.Suite, expected_, r.Expr(0.0).Div(0.0), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// math_logic/div.yaml line #46
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('a') / 0.8 */

		suite.T().Log("About to run line #46: r.Expr('a').Div(0.8)")

		runAndAssert(suite.Suite, expected_, r.Expr("a").Div(0.8), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// math_logic/div.yaml line #50
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr(1) / 'a' */

		suite.T().Log("About to run line #50: r.Expr(1).Div('a')")

		runAndAssert(suite.Suite, expected_, r.Expr(1).Div("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}
}
