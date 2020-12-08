package main

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.10.17",

	Tasks: []typgo.Tasker{
		// compile
		&typgo.GoBuild{MainPackage: "."},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/**", "pkg/**"},
		},
		// run
		&typgo.RunBinary{
			Before: typgo.TaskNames{"build"},
		},
		// examples
		&typgo.Task{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				c.Execute(&typgo.Bash{
					Name:   "go",
					Args:   []string{"test", "./examples/..."},
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				})
				return nil
			}),
		},
		// release
		&typrls.ReleaseTool{
			Before:    typgo.TaskNames{"test", "examples"},
			Publisher: &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
