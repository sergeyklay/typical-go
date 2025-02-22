package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func main() {
	// NOTE: ProjectName and ProjectVersion passed from descriptor in "tools/typical-build" when gobuild
	fmt.Println("-------------------------------")
	fmt.Println("typgo.ProjectName: " + typgo.ProjectName)
	fmt.Println("typgo.ProjectVersion: " + typgo.ProjectVersion)
	fmt.Println("-------------------------------")
	fmt.Println()

	fmt.Println("Hello, Typical")
}
