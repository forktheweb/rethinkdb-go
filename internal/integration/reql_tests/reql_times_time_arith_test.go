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

// Test basic time arithmetic
func TestTimesTimeArithSuite(t *testing.T) {
	suite.Run(t, new(TimesTimeArithSuite))
}

type TimesTimeArithSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TimesTimeArithSuite) SetupTest() {
	suite.T().Log("Setting up TimesTimeArithSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_times_ar").Exec(suite.session)
	err = r.DBCreate("db_times_ar").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_times_ar").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *TimesTimeArithSuite) TearDownSuite() {
	suite.T().Log("Tearing down TimesTimeArithSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_times_ar").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TimesTimeArithSuite) TestCases() {
	suite.T().Log("Running TimesTimeArithSuite: Test basic time arithmetic")

	// times/time_arith.yaml line #3
	// rt1 = 1375147296.681
	suite.T().Log("Possibly executing: var rt1 float64 = 1375147296.681")

	rt1 := 1375147296.681
	_ = rt1 // Prevent any noused variable errors

	// times/time_arith.yaml line #4
	// rt2 = 1375147296.682
	suite.T().Log("Possibly executing: var rt2 float64 = 1375147296.682")

	rt2 := 1375147296.682
	_ = rt2 // Prevent any noused variable errors

	// times/time_arith.yaml line #5
	// rt3 = 1375147297.681
	suite.T().Log("Possibly executing: var rt3 float64 = 1375147297.681")

	rt3 := 1375147297.681
	_ = rt3 // Prevent any noused variable errors

	// times/time_arith.yaml line #6
	// rt4 = 2375147296.681
	suite.T().Log("Possibly executing: var rt4 float64 = 2375147296.681")

	rt4 := 2375147296.681
	_ = rt4 // Prevent any noused variable errors

	// times/time_arith.yaml line #7
	// rts = [rt1, rt2, rt3, rt4]
	suite.T().Log("Possibly executing: var rts []interface{} = []interface{}{rt1, rt2, rt3, rt4}")

	rts := []interface{}{rt1, rt2, rt3, rt4}
	_ = rts // Prevent any noused variable errors

	// times/time_arith.yaml line #9
	// t1 = r.epoch_time(rt1)
	suite.T().Log("Possibly executing: var t1 r.Term = r.EpochTime(rt1)")

	t1 := r.EpochTime(rt1)
	_ = t1 // Prevent any noused variable errors

	// times/time_arith.yaml line #10
	// t2 = r.epoch_time(rt2)
	suite.T().Log("Possibly executing: var t2 r.Term = r.EpochTime(rt2)")

	t2 := r.EpochTime(rt2)
	_ = t2 // Prevent any noused variable errors

	// times/time_arith.yaml line #11
	// t3 = r.epoch_time(rt3)
	suite.T().Log("Possibly executing: var t3 r.Term = r.EpochTime(rt3)")

	t3 := r.EpochTime(rt3)
	_ = t3 // Prevent any noused variable errors

	// times/time_arith.yaml line #12
	// t4 = r.epoch_time(rt4)
	suite.T().Log("Possibly executing: var t4 r.Term = r.EpochTime(rt4)")

	t4 := r.EpochTime(rt4)
	_ = t4 // Prevent any noused variable errors

	// times/time_arith.yaml line #13
	// ts = r.expr([t1, t2, t3, t4])
	suite.T().Log("Possibly executing: var ts r.Term = r.Expr([]interface{}{t1, t2, t3, t4})")

	ts := r.Expr([]interface{}{t1, t2, t3, t4})
	_ = ts // Prevent any noused variable errors

	{
		// times/time_arith.yaml line #17
		/* true */
		var expected_ bool = true
		/* ((t2 - t1) * 1000).do(lambda x:(x > 0.99) & (x < 1.01)) */

		suite.T().Log("About to run line #17: r.Sub(t2, t1).Mul(1000).Do(func(x r.Term) interface{} { return r.Gt(x, 0.99).And(r.Lt(x, 1.01))})")

		runAndAssert(suite.Suite, expected_, r.Sub(t2, t1).Mul(1000).Do(func(x r.Term) interface{} { return r.Gt(x, 0.99).And(r.Lt(x, 1.01)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #17")
	}

	{
		// times/time_arith.yaml line #20
		/* 1 */
		var expected_ int = 1
		/* t3 - t1 */

		suite.T().Log("About to run line #20: r.Sub(t3, t1)")

		runAndAssert(suite.Suite, expected_, r.Sub(t3, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// times/time_arith.yaml line #23
		/* 1000000000 */
		var expected_ int = 1000000000
		/* t4 - t1 */

		suite.T().Log("About to run line #23: r.Sub(t4, t1)")

		runAndAssert(suite.Suite, expected_, r.Sub(t4, t1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// times/time_arith.yaml line #28
		/* true */
		var expected_ bool = true
		/* ((t1 - t2) * 1000).do(lambda x:(x < -0.99) & (x > -1.01)) */

		suite.T().Log("About to run line #28: r.Sub(t1, t2).Mul(1000).Do(func(x r.Term) interface{} { return r.Lt(x, -0.99).And(r.Gt(x, -1.01))})")

		runAndAssert(suite.Suite, expected_, r.Sub(t1, t2).Mul(1000).Do(func(x r.Term) interface{} { return r.Lt(x, -0.99).And(r.Gt(x, -1.01)) }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// times/time_arith.yaml line #31
		/* -1 */
		var expected_ int = -1
		/* t1 - t3 */

		suite.T().Log("About to run line #31: r.Sub(t1, t3)")

		runAndAssert(suite.Suite, expected_, r.Sub(t1, t3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// times/time_arith.yaml line #34
		/* -1000000000 */
		var expected_ int = -1000000000
		/* t1 - t4 */

		suite.T().Log("About to run line #34: r.Sub(t1, t4)")

		runAndAssert(suite.Suite, expected_, r.Sub(t1, t4), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// times/time_arith.yaml line #39
		/* ([rt1, rt2, rt3, rt4]) */
		var expected_ []interface{} = []interface{}{rt1, rt2, rt3, rt4}
		/* ts.map(lambda x:t1 + (x - t1)).map(lambda x:x.to_epoch_time()) */

		suite.T().Log("About to run line #39: ts.Map(func(x r.Term) interface{} { return r.Add(t1, r.Sub(x, t1))}).Map(func(x r.Term) interface{} { return x.ToEpochTime()})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(x r.Term) interface{} { return r.Add(t1, r.Sub(x, t1)) }).Map(func(x r.Term) interface{} { return x.ToEpochTime() }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #39")
	}

	{
		// times/time_arith.yaml line #43
		/* err("ReqlQueryLogicError", "Expected type NUMBER but found PTYPE<TIME>.", []) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type NUMBER but found PTYPE<TIME>.")
		/* ts.map(lambda x:(t1 + x) - t1).map(lambda x:x.to_epoch_time()) */

		suite.T().Log("About to run line #43: ts.Map(func(x r.Term) interface{} { return r.Add(t1, x).Sub(t1)}).Map(func(x r.Term) interface{} { return x.ToEpochTime()})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(x r.Term) interface{} { return r.Add(t1, x).Sub(t1) }).Map(func(x r.Term) interface{} { return x.ToEpochTime() }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #43")
	}

	{
		// times/time_arith.yaml line #47
		/* ([rt1, rt2, rt3, rt4]) */
		var expected_ []interface{} = []interface{}{rt1, rt2, rt3, rt4}
		/* ts.map(lambda x:t1 - (t1 - x)).map(lambda x:x.to_epoch_time()) */

		suite.T().Log("About to run line #47: ts.Map(func(x r.Term) interface{} { return r.Sub(t1, r.Sub(t1, x))}).Map(func(x r.Term) interface{} { return x.ToEpochTime()})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(x r.Term) interface{} { return r.Sub(t1, r.Sub(t1, x)) }).Map(func(x r.Term) interface{} { return x.ToEpochTime() }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #47")
	}

	{
		// times/time_arith.yaml line #52
		/* ([[[false, true,  true,  false, true,  false],
		[true,  true,  false, true,  false, false],
		[true,  true,  false, true,  false, false],
		[true,  true,  false, true,  false, false]],
		[[false, false, false, true,  true,  true],
		[false, true,  true,  false, true,  false],
		[true,  true,  false, true,  false, false],
		[true,  true,  false, true,  false, false]],
		[[false, false, false, true,  true,  true],
		[false, false, false, true,  true,  true],
		[false, true,  true,  false, true,  false],
		[true,  true,  false, true,  false, false]],
		[[false, false, false, true,  true,  true],
		[false, false, false, true,  true,  true],
		[false, false, false, true,  true,  true],
		[false, true,  true,  false, true,  false]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{false, true, true, false, true, false}, []interface{}{true, true, false, true, false, false}, []interface{}{true, true, false, true, false, false}, []interface{}{true, true, false, true, false, false}}, []interface{}{[]interface{}{false, false, false, true, true, true}, []interface{}{false, true, true, false, true, false}, []interface{}{true, true, false, true, false, false}, []interface{}{true, true, false, true, false, false}}, []interface{}{[]interface{}{false, false, false, true, true, true}, []interface{}{false, false, false, true, true, true}, []interface{}{false, true, true, false, true, false}, []interface{}{true, true, false, true, false, false}}, []interface{}{[]interface{}{false, false, false, true, true, true}, []interface{}{false, false, false, true, true, true}, []interface{}{false, false, false, true, true, true}, []interface{}{false, true, true, false, true, false}}}
		/* ts.map(lambda x:ts.map(lambda y:[x < y, x <= y, x == y, x != y, x >= y, x > y])) */

		suite.T().Log("About to run line #52: ts.Map(func(x r.Term) interface{} { return ts.Map(func(y r.Term) interface{} { return []interface{}{r.Lt(x, y), r.Le(x, y), r.Eq(x, y), r.Ne(x, y), r.Ge(x, y), r.Gt(x, y)}})})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(x r.Term) interface{} {
			return ts.Map(func(y r.Term) interface{} {
				return []interface{}{r.Lt(x, y), r.Le(x, y), r.Eq(x, y), r.Ne(x, y), r.Ge(x, y), r.Gt(x, y)}
			})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #52")
	}

	// times/time_arith.yaml line #73
	// datum_types = r.expr([null, true, false, 1, "1", [1], {"1":1}, r.binary(b'')])
	suite.T().Log("Possibly executing: var datum_types r.Term = r.Expr([]interface{}{nil, true, false, 1, '1', []interface{}{1}, map[interface{}]interface{}{'1': 1, }, r.Binary([]byte{})})")

	datum_types := r.Expr([]interface{}{nil, true, false, 1, "1", []interface{}{1}, map[interface{}]interface{}{"1": 1}, r.Binary([]byte{})})
	_ = datum_types // Prevent any noused variable errors

	{
		// times/time_arith.yaml line #79
		/* ([[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[false, false, false, true,  true,  true],
		[true,  true,  false, true,  false, false]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]],
		[[true,  true,  false, true,  false, false],
		[false, false, false, true,  true,  true]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{false, false, false, true, true, true}, []interface{}{true, true, false, true, false, false}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}, []interface{}{[]interface{}{true, true, false, true, false, false}, []interface{}{false, false, false, true, true, true}}}
		/* datum_types.map(lambda x:r.expr([[x, t1], [t1, x]]).map(lambda xy:xy[0].do(lambda x2:xy[1].do(lambda y:[x2 < y, x2 <= y, x2 == y, x2 != y, x2 >= y, x2 > y])))) */

		suite.T().Log("About to run line #79: datum_types.Map(func(x r.Term) interface{} { return r.Expr([]interface{}{[]interface{}{x, t1}, []interface{}{t1, x}}).Map(func(xy r.Term) interface{} { return xy.AtIndex(0).Do(func(x2 r.Term) interface{} { return xy.AtIndex(1).Do(func(y r.Term) interface{} { return []interface{}{r.Lt(x2, y), r.Le(x2, y), r.Eq(x2, y), r.Ne(x2, y), r.Ge(x2, y), r.Gt(x2, y)}})})})})")

		runAndAssert(suite.Suite, expected_, datum_types.Map(func(x r.Term) interface{} {
			return r.Expr([]interface{}{[]interface{}{x, t1}, []interface{}{t1, x}}).Map(func(xy r.Term) interface{} {
				return xy.AtIndex(0).Do(func(x2 r.Term) interface{} {
					return xy.AtIndex(1).Do(func(y r.Term) interface{} {
						return []interface{}{r.Lt(x2, y), r.Le(x2, y), r.Eq(x2, y), r.Ne(x2, y), r.Ge(x2, y), r.Gt(x2, y)}
					})
				})
			})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #79")
	}

	{
		// times/time_arith.yaml line #99
		/* ([[[false, true,  true,  true],
		[false, false, true,  true],
		[false, false, false, true],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, true,  true],
		[false, false, false, true],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, true],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{false, true, true, true}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, true}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}}}
		/* ts.map(lambda a:ts.map(lambda b:ts.map(lambda c:b.during(a, c)))) */

		suite.T().Log("About to run line #99: ts.Map(func(a r.Term) interface{} { return ts.Map(func(b r.Term) interface{} { return ts.Map(func(c r.Term) interface{} { return b.During(a, c)})})})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(a r.Term) interface{} {
			return ts.Map(func(b r.Term) interface{} { return ts.Map(func(c r.Term) interface{} { return b.During(a, c) }) })
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #99")
	}

	{
		// times/time_arith.yaml line #119
		/* ([[[false, false, false, false],
		[false, false, true,  true],
		[false, false, false, true],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, true],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, true}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}}}
		/* ts.map(lambda a:ts.map(lambda b:ts.map(lambda c:b.during(a, c, left_bound='open')))) */

		suite.T().Log("About to run line #119: ts.Map(func(a r.Term) interface{} { return ts.Map(func(b r.Term) interface{} { return ts.Map(func(c r.Term) interface{} { return b.During(a, c).OptArgs(r.DuringOpts{LeftBound: 'open', })})})})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(a r.Term) interface{} {
			return ts.Map(func(b r.Term) interface{} {
				return ts.Map(func(c r.Term) interface{} { return b.During(a, c).OptArgs(r.DuringOpts{LeftBound: "open"}) })
			})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #119")
	}

	{
		// times/time_arith.yaml line #139
		/* ([[[true,  true,  true,  true],
		[false, true,  true,  true],
		[false, false, true,  true],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, true,  true,  true],
		[false, false, true,  true],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, true,  true],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, true]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{true, true, true, true}, []interface{}{false, true, true, true}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, true, true, true}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, true}}}
		/* ts.map(lambda a:ts.map(lambda b:ts.map(lambda c:b.during(a, c, right_bound='closed')))) */

		suite.T().Log("About to run line #139: ts.Map(func(a r.Term) interface{} { return ts.Map(func(b r.Term) interface{} { return ts.Map(func(c r.Term) interface{} { return b.During(a, c).OptArgs(r.DuringOpts{RightBound: 'closed', })})})})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(a r.Term) interface{} {
			return ts.Map(func(b r.Term) interface{} {
				return ts.Map(func(c r.Term) interface{} { return b.During(a, c).OptArgs(r.DuringOpts{RightBound: "closed"}) })
			})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #139")
	}

	{
		// times/time_arith.yaml line #159
		/* ([[[false, false, false, false],
		[false, true,  true,  true],
		[false, false, true,  true],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, true,  true],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, true]],
		[[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false],
		[false, false, false, false]]]) */
		var expected_ []interface{} = []interface{}{[]interface{}{[]interface{}{false, false, false, false}, []interface{}{false, true, true, true}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, true, true}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, true}}, []interface{}{[]interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}, []interface{}{false, false, false, false}}}
		/* ts.map(lambda a:ts.map(lambda b:ts.map(lambda c:b.during(a, c, left_bound='open', right_bound='closed')))) */

		suite.T().Log("About to run line #159: ts.Map(func(a r.Term) interface{} { return ts.Map(func(b r.Term) interface{} { return ts.Map(func(c r.Term) interface{} { return b.During(a, c).OptArgs(r.DuringOpts{LeftBound: 'open', RightBound: 'closed', })})})})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(a r.Term) interface{} {
			return ts.Map(func(b r.Term) interface{} {
				return ts.Map(func(c r.Term) interface{} {
					return b.During(a, c).OptArgs(r.DuringOpts{LeftBound: "open", RightBound: "closed"})
				})
			})
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #159")
	}

	{
		// times/time_arith.yaml line #179
		/* rts */
		var expected_ []interface{} = rts
		/* ts.map(lambda x:x.date() + x.time_of_day()).map(lambda x:x.to_epoch_time()) */

		suite.T().Log("About to run line #179: ts.Map(func(x r.Term) interface{} { return x.Date().Add(x.TimeOfDay())}).Map(func(x r.Term) interface{} { return x.ToEpochTime()})")

		runAndAssert(suite.Suite, expected_, ts.Map(func(x r.Term) interface{} { return x.Date().Add(x.TimeOfDay()) }).Map(func(x r.Term) interface{} { return x.ToEpochTime() }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #179")
	}

	{
		// times/time_arith.yaml line #185
		/* rt1 */
		var expected_ float64 = rt1
		/* r.epoch_time(rt1).do(r.js("(function(data){return data})")).to_epoch_time() */

		suite.T().Log("About to run line #185: r.EpochTime(rt1).Do(r.JS('(function(data){return data})')).ToEpochTime()")

		runAndAssert(suite.Suite, expected_, r.EpochTime(rt1).Do(r.JS("(function(data){return data})")).ToEpochTime(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #185")
	}

	{
		// times/time_arith.yaml line #190
		/* ("2012-08-01T00:00:00+00:00") */
		var expected_ string = "2012-08-01T00:00:00+00:00"
		/* r.do(r.js("new Date('2012-08-01')")).to_iso8601() */

		suite.T().Log("About to run line #190: r.Do(r.JS('new Date('2012-08-01')')).ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.Do(r.JS("new Date('2012-08-01')")).ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #190")
	}

	{
		// times/time_arith.yaml line #195
		/* ("2012-08-01T00:00:00+00:00") */
		var expected_ string = "2012-08-01T00:00:00+00:00"
		/* r.do(r.js("(function(x){doc = new Object(); doc.date = new Date('2012-08-01'); return doc;})"))["date"].to_iso8601() */

		suite.T().Log("About to run line #195: r.Do(r.JS('(function(x){doc = new Object(); doc.date = new Date('2012-08-01'); return doc;})')).AtIndex('date').ToISO8601()")

		runAndAssert(suite.Suite, expected_, r.Do(r.JS("(function(x){doc = new Object(); doc.date = new Date('2012-08-01'); return doc;})")).AtIndex("date").ToISO8601(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #195")
	}
}
