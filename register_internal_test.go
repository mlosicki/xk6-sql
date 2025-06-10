package sql

import (
	"testing"

	"github.com/mlosicki/xk6-sql-ext/sql"
	"github.com/stretchr/testify/require"
	"go.k6.io/k6/ext"
)

func Test_register(t *testing.T) {
	t.Parallel()

	require.Contains(t, ext.Get(ext.JSExtension), sql.ImportPath)
}
