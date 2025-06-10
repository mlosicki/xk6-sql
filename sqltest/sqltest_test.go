package sqltest_test

import (
	_ "embed"
	"testing"

	"github.com/mlosicki/xk6-sql-ext/sql"
	"github.com/mlosicki/xk6-sql-ext/sqltest"
	_ "github.com/proullon/ramsql/driver"
)

//go:embed testdata/script.js
var script string

func TestRunScript(t *testing.T) {
	t.Parallel()

	sql.RegisterModule("ramsql")

	sqltest.RunScript(t, "ramsql", "testdb", script)
}
