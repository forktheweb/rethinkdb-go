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

// Tests meta operations in composite queries
func TestMetaCompositeSuite(t *testing.T) {
	suite.Run(t, new(MetaCompositeSuite))
}

type MetaCompositeSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MetaCompositeSuite) SetupTest() {
	suite.T().Log("Setting up MetaCompositeSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_composite").Exec(suite.session)
	err = r.DBCreate("db_composite").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_composite").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *MetaCompositeSuite) TearDownSuite() {
	suite.T().Log("Tearing down MetaCompositeSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_composite").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MetaCompositeSuite) TestCases() {
	suite.T().Log("Running MetaCompositeSuite: Tests meta operations in composite queries")

	{
		// meta/composite.py.yaml line #4
		/* ({'dbs_created':3,'config_changes':arrlen(3)}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"dbs_created": 3, "config_changes": arrlen(3)}
		/* r.expr([1,2,3]).for_each(r.db_create('db_' + r.row.coerce_to('string'))) */

		suite.T().Log("About to run line #4: r.Expr([]interface{}{1, 2, 3}).ForEach(r.DBCreate(r.Add('db_', r.Row.CoerceTo('string'))))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(r.DBCreate(r.Add("db_", r.Row.CoerceTo("string")))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #4")
	}

	{
		// meta/composite.py.yaml line #8
		/* partial({'tables_created':9}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"tables_created": 9})
		/* r.db_list().set_difference(["rethinkdb", "test"]).for_each(lambda db_name:
		r.expr([1,2,3]).for_each(lambda i:
		r.db(db_name).table_create('tbl_' + i.coerce_to('string')))) */

		suite.T().Log("About to run line #8: r.DBList().SetDifference([]interface{}{'rethinkdb', 'test'}).ForEach(func(db_name r.Term) interface{} { return r.Expr([]interface{}{1, 2, 3}).ForEach(func(i r.Term) interface{} { return r.DB(db_name).TableCreate(r.Add('tbl_', i.CoerceTo('string')))})})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(
			func (i r.Term) interface{}{
				return r.Expr([]interface{}{1, 2, 3}).ForEach(func(j r.Term) interface{} {
					return r.DB(r.Add("db_", i.CoerceTo("string"))).TableCreate(r.Add("tbl_", j.CoerceTo("string")))
				})
			}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #8")
	}

	{
		// meta/composite.py.yaml line #13
		/* partial({'dbs_dropped':3,'tables_dropped':9}) */
		var expected_ compare.Expected = compare.PartialMatch(map[interface{}]interface{}{"dbs_dropped": 3, "tables_dropped": 9})
		/* r.db_list().set_difference(["rethinkdb", "test"]).for_each(r.db_drop(r.row)) */

		suite.T().Log("About to run line #13: r.DBList().SetDifference([]interface{}{'rethinkdb', 'test'}).ForEach(r.DBDrop(r.Row))")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{1, 2, 3}).ForEach(r.DBDrop(r.Add("db_", r.Row.CoerceTo("string")))), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #13")
	}
}
