package typenv

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// Layout of project
	Layout = struct {
		App     string
		Bin     string
		Cmd     string
		Temp    string
		Mock    string
		Release string
	}{
		App:     "app",
		Cmd:     "cmd",
		Bin:     "bin",
		Temp:    ".typical-tmp",
		Mock:    "mock",
		Release: "release",
	}

	Readme      = "README.md"
	ProjectName = projectName()

	AppBin      = fmt.Sprintf("%s/%s", Layout.Bin, ProjectName)
	AppMainPath = fmt.Sprintf("%s/%s", Layout.Cmd, ProjectName)

	BuildTool         = "buildtool"
	BuildToolBin      = fmt.Sprintf("%s/%s-%s", Layout.Bin, ProjectName, BuildTool)
	BuildToolMainPath = fmt.Sprintf("%s/%s-%s", Layout.Cmd, ProjectName, BuildTool)

	DescriptorFile = "typical/descriptor.go"
	ChecksumFile   = Layout.Temp + "/checksum"
)

func projectName() (s string) {
	var err error
	if s, err = os.Getwd(); err != nil {
		return "noname"
	}
	return filepath.Base(s)
}
