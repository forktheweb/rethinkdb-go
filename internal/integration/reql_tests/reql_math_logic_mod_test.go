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

// Tests for the basic usage of the mod operation
func TestMathLogicModSuite(t *testing.T) {
	suite.Run(t, new(MathLogicModSuite))
}

type MathLogicModSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MathLogicModSuite) SetupTest() {
	suite.T().Log("Setting up MathLogicModSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("test").Exec(suite.session)
	err = r.DBCreate("test").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("test").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MathLogicModSuite) TearDownSuite() {
	suite.T().Log("Tearing down MathLogicModSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("test").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MathLogicModSuite) TestCases() {
	suite.T().Log("Running MathLogicModSuite: Tests for the basic usage of the mod operation")

	{
		// math_logic/mod.yaml line #6
		/* 1 */
		var expected_ int = 1
		/* r.expr(10) % 3 */

		suite.T().Log("About to run line #6: r.Expr(10).Mod(3)")

		runAndAssert(suite.Suite, expected_, r.Expr(10).Mod(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// math_logic/mod.yaml line #7
		/* 1 */
		var expected_ int = 1
		/* 10 % r.expr(3) */

		suite.T().Log("About to run line #7: r.Mod(10, r.Expr(3))")

		runAndAssert(suite.Suite, expected_, r.Mod(10, r.Expr(3)), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// math_logic/mod.yaml line #8
		/* 1 */
		var expected_ int = 1
		/* r.expr(10).mod(3) */

		suite.T().Log("About to run line #8: r.Expr(10).Mod(3)")

		runAndAssert(suite.Suite, expected_, r.Expr(10).Mod(3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// math_logic/mod.yaml line #16
		/* -1 */
		var expected_ int = -1
		/* r.expr(-10) % -3 */

		suite.T().Log("About to run line #16: r.Expr(-10).Mod(-3)")

		runAndAssert(suite.Suite, expected_, r.Expr(-10).Mod(-3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// math_logic/mod.yaml line #22
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [1]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr(4) % 'a' */

		suite.T().Log("About to run line #22: r.Expr(4).Mod('a')")

		runAndAssert(suite.Suite, expected_, r.Expr(4).Mod("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// math_logic/mod.yaml line #27
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('a') % 1 */

		suite.T().Log("About to run line #27: r.Expr('a').Mod(1)")

		runAndAssert(suite.Suite, expected_, r.Expr("a").Mod(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #27")
	}

	{
		// math_logic/mod.yaml line #32
		/* err('ReqlQueryLogicError', 'Expected type NUMBER but found STRING.', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found STRING.")
		/* r.expr('a') % 'b' */

		suite.T().Log("About to run line #32: r.Expr('a').Mod('b')")

		runAndAssert(suite.Suite, expected_, r.Expr("a").Mod("b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #32")
	}
}
