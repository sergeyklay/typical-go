package typdocker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typdocker"
)

func TestRecipe_DockerCompose(t *testing.T) {
	expected := &typdocker.Recipe{}
	recipe, err := expected.ComposeV3()
	require.Equal(t, expected, recipe)
	require.NoError(t, err)
}
