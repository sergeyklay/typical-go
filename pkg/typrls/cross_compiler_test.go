package typrls_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
	"github.com/urfave/cli/v2"
)

func TestCrossCompile(t *testing.T) {

	testcases := []struct {
		TestName string
		typrls.CrossCompiler
		TagName         string
		RunExpectations []*typgo.MockCommand
		ExpectedErr     string
	}{
		{
			CrossCompiler: typrls.CrossCompiler{
				Targets: []typrls.Target{"darwin/amd64", "linux/amd64"},
			},
			TagName: "v0.0.1",
			RunExpectations: []*typgo.MockCommand{
				{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_darwin_amd64 ./cmd/some-project"},
				{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_linux_amd64 ./cmd/some-project"},
			},
		},
		{
			TestName: "go build error",
			CrossCompiler: typrls.CrossCompiler{
				Targets: []typrls.Target{"darwin/amd64"},
			},
			TagName: "v0.0.1",
			RunExpectations: []*typgo.MockCommand{
				{
					CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=v0.0.1\" -o /some-project_v0.0.1_darwin_amd64 ./cmd/some-project",
					ReturnError: errors.New("some-error"),
				},
			},
			ExpectedErr: "some-error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			c := &typrls.Context{
				TagName: tt.TagName,
				Context: &typgo.Context{
					Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
					Descriptor: &typgo.Descriptor{
						ProjectName:    "some-project",
						ProjectVersion: "0.0.1",
					},
				},
			}
			defer c.PatchBash(tt.RunExpectations)(t)

			err := tt.Release(c)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTarget(t *testing.T) {
	testcases := []struct {
		TestName string
		typrls.Target
		ExpectedOS   string
		ExpectedArch string
	}{
		{Target: "darwin/amd64", ExpectedOS: "darwin", ExpectedArch: "amd64"},
		{Target: "linux/amd64", ExpectedOS: "linux", ExpectedArch: "amd64"},
		{Target: "no-slash", ExpectedOS: "", ExpectedArch: ""},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.ExpectedOS, tt.OS())
			require.Equal(t, tt.ExpectedArch, tt.Arch())
		})
	}
}
