package typgo

import (
	"context"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// BuildCli detail
	BuildCli struct {
		*Descriptor
		ASTStore *typast.ASTStore
		Precond  *typtmpl.Precond
	}

	// Context of build tool
	Context struct {
		typlog.Logger
		*cli.Context
		*BuildCli
	}

	// CliFunc is command line function
	CliFunc func(*Context) error
)

func createBuildCli(d *Descriptor) *BuildCli {
	var (
		astStore *typast.ASTStore
		err      error
	)
	appDirs, appFiles := WalkLayout(d.Layouts)

	if astStore, err = typast.CreateASTStore(appFiles...); err != nil {
		// TODO:
		// logger.Warn(err.Error())
	}

	return &BuildCli{
		Descriptor: d,
		ASTStore:   astStore,
		Precond: &typtmpl.Precond{
			Imports: retrImports(appDirs),
			Package: "main",
		},
	}
}

func retrImports(dirs []string) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", typvar.ProjectPkg, dir))
		}
	}
	return imports
}

// Context of build-cli
func (b *BuildCli) Context(name string, c *cli.Context) *Context {
	return &Context{
		Logger: typlog.Logger{
			Name: name,
		},
		Context:  c,
		BuildCli: b,
	}
}

// Commands of build-tool
func (b *BuildCli) Commands() (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(b),
		cmdRun(b),
		cmdPublish(b),
		cmdClean(b),
	}

	if b.Utility != nil {
		for _, cmd := range b.Utility.Commands(b) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

// ActionFn to return related action func
func (b *BuildCli) ActionFn(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		c := b.Context(strings.ToUpper(name), cli)
		return fn(c)
	}
}

//
// Context
//

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}
