package typapp_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestDtorAnnotation_Annotate(t *testing.T) {
	os.MkdirAll("folder3/pkg3", 0777)
	defer os.RemoveAll("folder3/pkg3")

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	dtorAnnot := &typapp.DtorAnnotation{}
	ctx := &typannot.Context{
		Destination: "folder3/pkg3",
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
		Summary: &typannot.Summary{
			Annots: []*typannot.Annot{
				{TagName: "@dtor", Decl: &typannot.Decl{Name: "Clean", Package: "pkg", Type: &typannot.FuncType{}}},
			},
		},
	}

	require.NoError(t, dtorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile("folder3/pkg3/dtor_annotated.go")
	require.Equal(t, `package pkg3

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() { 
	typapp.AppendDtor(
		&typapp.Destructor{Fn: pkg.Clean},
	)
}`, string(b))

}

func TestDtorAnnotation_Annotate_Predefined(t *testing.T) {
	os.MkdirAll("folder4/pkg4", 0777)
	defer os.RemoveAll("folder4/pkg4")

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	dtorAnnot := &typapp.DtorAnnotation{
		Target:   "some-target",
		TagName:  "@some-tag",
		Template: "some-template",
	}
	ctx := &typannot.Context{
		Destination: "folder4/pkg4",
		Context: &typgo.Context{
			BuildSys: &typgo.BuildSys{
				Descriptor: &typgo.Descriptor{ProjectName: "some-project"},
			},
		},
		Summary: &typannot.Summary{
			Annots: []*typannot.Annot{
				{TagName: "@some-tag", Decl: &typannot.Decl{Name: "Clean", Package: "pkg", Type: &typannot.FuncType{}}},
			},
		},
	}

	require.NoError(t, dtorAnnot.Annotate(ctx))

	b, _ := ioutil.ReadFile("folder4/pkg4/some-target")
	require.Equal(t, `some-template`, string(b))
}

func TestDtorAnnotation_Annotate_RemoveTargetWhenNoAnnotation(t *testing.T) {
	os.MkdirAll("folder4/pkg4", 0777)
	defer os.RemoveAll("folder4/pkg4")

	dtorAnnot := &typapp.DtorAnnotation{Target: "some-target"}
	ctx := &typannot.Context{
		Destination: "folder4/pkg4",
		Context:     &typgo.Context{},
		Summary:     &typannot.Summary{},
	}

	ioutil.WriteFile("folder4/pkg4/some-target", []byte("some-content"), 0777)
	require.NoError(t, dtorAnnot.Annotate(ctx))
	_, err := os.Stat("folder4/pkg4/some-target")
	require.True(t, os.IsNotExist(err))
}
