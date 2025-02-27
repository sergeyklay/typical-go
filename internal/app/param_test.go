package app_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGetParam(t *testing.T) {
	c := &typgo.Context{
		Context: cliContext([]string{
			"-typical-build=1",
			"-typical-tmp=2",
			"-project-pkg=github.com/user/project",
		}),
	}
	param, err := app.GetParam(c)

	require.NoError(t, err)
	require.Equal(t, &app.Param{
		TypicalBuild: "1",
		TypicalTmp:   "2",
		ProjectPkg:   "github.com/user/project",
		ProjectName:  "project",
		SetupTarget:  "project",
	}, param)
}

func TestGetParam_Default(t *testing.T) {
	c := &typgo.Context{
		Context: cliContext([]string{}),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go list -m", OutputBytes: []byte("some-package")},
	})(t)

	param, err := app.GetParam(c)

	require.NoError(t, err)
	require.Equal(t, &app.Param{
		TypicalBuild: "tools/typical-build",
		TypicalTmp:   ".typical-tmp",
		ProjectPkg:   "some-package",
		ProjectName:  "some-package",
		SetupTarget:  ".",
	}, param)
}

func TestGetParam_Default_FailedRetrivePackage(t *testing.T) {
	c := &typgo.Context{
		Context: cliContext([]string{}),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{
			CommandLine: "go list -m",
			ErrorBytes:  []byte("error-message"),
			ReturnError: errors.New("some-error"),
		},
	})(t)
	_, err := app.GetParam(c)
	require.EqualError(t, err, "some-error: error-message")
}

func cliContext(args []string) *cli.Context {
	flagSet := &flag.FlagSet{}
	flagSet.String(app.TypicalTmpParam, app.DefaultTypicalTmp, "")
	flagSet.String(app.TypicalBuildParam, app.DefaultTypicalBuild, "")
	flagSet.String(app.ProjectPkgParam, "", "")
	flagSet.Bool("go-mod", false, "")
	flagSet.Bool("new", false, "")

	flagSet.Parse(args)

	return cli.NewContext(nil, flagSet, nil)
}
