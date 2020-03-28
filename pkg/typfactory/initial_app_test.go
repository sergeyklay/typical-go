package typfactory_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

func TestInitialApp(t *testing.T) {
	testcases := []testcase{
		{
			Writer: &typfactory.InitialApp{
				Imports:      []string{"import1", "import2"},
				Constructors: []string{"constructor1", "constructor2"},
			},
			expected: `package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"import1"
	"import2"
)

func init() {
	typapp.AppendConstructor(
		typapp.NewConstructor(constructor1),
		typapp.NewConstructor(constructor2),
	)
}
`,
		},
	}
	for _, tt := range testcases {
		var debugger strings.Builder
		require.NoError(t, tt.Write(&debugger))
		require.Equal(t, tt.expected, debugger.String())
	}
}
