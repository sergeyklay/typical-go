package typast

import (
	"fmt"
	"path/filepath"
	"unicode"
	"unicode/utf8"

	"github.com/typical-go/typical-go/pkg/typgo"
)

// IsFunc return true if annotation is function
func IsFunc(annot *Annot) bool {
	funcDecl, ok := annot.Type.(*FuncDecl)
	return ok && !funcDecl.IsMethod()
}

// IsMethod return true if annotation is method
func IsMethod(annot *Annot) bool {
	funcDecl, ok := annot.Type.(*FuncDecl)
	return ok && funcDecl.IsMethod()
}

// IsStruct return true if annotation is struct
func IsStruct(annot *Annot) bool {
	_, ok := annot.Type.(*StructDecl)
	return ok
}

// IsInterface return true if annotation is struct
func IsInterface(annot *Annot) bool {
	_, ok := annot.Type.(*InterfaceDecl)
	return ok
}

// IsPublic return true if decl is public access
func IsPublic(annot *Annot) bool {
	rune, _ := utf8.DecodeRuneInString(annot.GetName())
	return unicode.IsUpper(rune)
}

// Package of annotation
func Package(annot *Annot) string {
	return fmt.Sprintf("%s/%s", typgo.ProjectPkg, filepath.Dir(annot.Path))
}
