package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typapp-sample",
	ProjectVersion: "1.0.0",
	ProjectLayouts: []string{"internal"},

	Tasks: []typgo.Tasker{
		// annotate
		&typast.Annotators{
			&typapp.CtorAnnotation{},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.BuildCmdRuns{"annotate", "build"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
