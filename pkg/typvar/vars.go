// Package typvar contains typical variable
package typvar

import (
	"fmt"
	"time"
)

var (
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string

	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string

	// ReleaseFolder location
	ReleaseFolder = "release"

	// BinFolder location
	BinFolder = "bin"

	// CmdFolder location
	CmdFolder = "cmd"

	// ConfigFile location
	ConfigFile = ".env"

	// TestTimeout duration
	TestTimeout = 25 * time.Second

	// TestCoverProfile location
	TestCoverProfile = "cover.out"
)

// Precond path
func Precond(name string) string {
	return fmt.Sprintf("%s/%s/precond_DO_NOT_EDIT.go", CmdFolder, name)
}
