package typgen

import (
	"go/ast"
)

type (

	// TypeDecl type declaration
	TypeDecl struct {
		GenDecl
		Name string
		Docs []string
	}
	// GenDecl generic declaration
	GenDecl struct {
		Docs []string
	}
)

func createGenDecl(genDecl *ast.GenDecl, file *File) []Type {
	var types []Type
	for _, spec := range genDecl.Specs {
		switch spec.(type) {
		case *ast.TypeSpec:
			typeSpec := spec.(*ast.TypeSpec)
			typeDecl := TypeDecl{
				GenDecl: GenDecl{Docs: docs(genDecl.Doc)},
				Name:    typeSpec.Name.Name,
				Docs:    docs(typeSpec.Doc),
			}

			switch typeSpec.Type.(type) {
			case *ast.InterfaceType:
				types = append(types, CreateInterfaceDecl(typeDecl))
			case *ast.StructType:
				types = append(types, CreateStructDecl(typeDecl, typeSpec.Type.(*ast.StructType)))
			}
		}
	}
	return types
}

func docs(group *ast.CommentGroup) []string {
	if group == nil {
		return nil
	}
	var docs []string
	for _, comment := range group.List {
		docs = append(docs, comment.Text)
	}
	return docs
}

//
// TypeDecl
//

var _ Type = (*TypeDecl)(nil)

// GetName get name
func (t *TypeDecl) GetName() string {
	return t.Name
}

// GetDocs get doc
func (t *TypeDecl) GetDocs() []string {
	if t.Docs != nil {
		return t.Docs
	}
	return t.GenDecl.Docs
}
