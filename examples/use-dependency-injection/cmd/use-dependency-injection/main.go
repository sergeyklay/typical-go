package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	_ "github.com/typical-go/typical-go/examples/use-dependency-injection/internal/generated"
	"github.com/typical-go/typical-go/examples/use-dependency-injection/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	typapp.Start(helloworld.Main2)
}
