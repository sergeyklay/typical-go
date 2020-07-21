package typgo_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestRunCompile_Command(t *testing.T) {
	runCmd := &typgo.RunCmd{
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := runCmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")
}

func TestRunCompile_Precmd(t *testing.T) {
	sys := &typgo.BuildSys{
		Commands: []*cli.Command{
			{Name: "pre", Action: func(*cli.Context) error { return errors.New("pre-error") }},
		},
	}

	runCmd := &typgo.RunCmd{
		Precmds: []string{"pre"},
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := runCmd.Command(sys)
	require.EqualError(t, command.Action(&cli.Context{}), "pre-error")
}

func TestStdRun_Execute(t *testing.T) {
	stdRun := &typgo.StdRun{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Name: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"bin/some-name"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}

func TestStdRun_Execute_Predefined(t *testing.T) {
	stdRun := &typgo.StdRun{
		Binary: "some-binary",
	}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Name: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"some-binary"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}
