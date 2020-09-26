package typgo

var (
	// ProjectName of application. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=PROJECT-NAME`
	ProjectName string
	// ProjectVersion of applicatoin. Injected from gobuild ldflags
	// `-X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=PROJECT-NAME`
	ProjectVersion string
	// ProjectPkg only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	ProjectPkg string
	// TypicalTmp only available in BuildTool scope. The BuildTool must be compiled by wrapper.
	TypicalTmp string

	appHelpTemplate = `Typical Build

Usage:

{{"\t"}}./typicalw <command> [argument]

The commands are:
{{range .Commands}}
{{if not .HideHelp}}{{ "\t"}}{{join .Names ", "}}{{ "\t"}}{{.Usage}}{{end}}{{end}}

Use "./typicalw help <topic>" for more information about that topic
`
	subcommandHelpTemplate = `{{.Usage}}

Usage:

	{{.Name}} [command]
	
Commands:{{range .VisibleCategories}}
{{if .Name}}{{.Name}}:{{range .VisibleCommands}}
		{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
	{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
	
{{if .VisibleFlags}} 
Options:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}
`
)
