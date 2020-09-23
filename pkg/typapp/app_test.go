package typapp_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func ExampleStart() {
	// append contructor definition
	typapp.AppendCtor(&typapp.Constructor{
		Fn: func() string {
			return "World"
		},
	})

	// append destructor definition
	typapp.AppendDtor(&typapp.Destructor{
		Fn: func() {
			fmt.Println("clean something")
		},
	})

	// start the application
	err := typapp.Run(func(text string) {
		fmt.Printf("Hello %s\n", text)
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Output: Hello World
	// clean something
}

func TestStart(t *testing.T) {
	defer typapp.ClearCtors()
	defer typapp.ClearDtors()

	var debugger []string
	typapp.AppendCtor(&typapp.Constructor{
		Fn: func() string { return "success" },
	})
	typapp.AppendDtor(&typapp.Destructor{
		Fn: func() { debugger = append(debugger, "clean") },
	})

	require.NoError(t, typapp.Run(func(s string) {
		debugger = append(debugger, s)
	}))
	require.Equal(t, []string{"success", "clean"}, debugger)
}

func TestApp_Run_Error(t *testing.T) {
	var out strings.Builder
	typapp.Stdout = &out
	defer func() { typapp.Stdout = os.Stdout }()

	app := &typapp.App{
		EntryPoint: fnErr("entry-point-err"),
		Dtors: []*typapp.Destructor{
			{Fn: fnErr("dtor-1-err")},
			{Fn: fnErr("dtor-2-err")},
		},
	}
	require.EqualError(t, app.Run(), "entry-point-err")
	require.Equal(t, "WARN: dtor-1-err\nWARN: dtor-2-err\n", out.String())
}
func TestApp_Run_BadConstructor(t *testing.T) {
	app := &typapp.App{
		Ctors: []*typapp.Constructor{
			{Fn: "bad-contructor"},
		},
	}
	require.EqualError(t, app.Run(), "must provide constructor function, got bad-contructor (type string)")
}

func TestApp_Run(t *testing.T) {
	var text string
	app := &typapp.App{
		EntryPoint: func(s string) { text = s },
		Ctors: []*typapp.Constructor{
			{Fn: func() string { return "success" }},
		},
	}
	require.NoError(t, app.Run())
	require.Equal(t, "success", text)
}

func fnErr(errMsg string) func() error {
	return func() error {
		return errors.New(errMsg)
	}
}
