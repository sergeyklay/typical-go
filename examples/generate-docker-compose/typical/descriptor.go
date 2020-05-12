package typical

import (
	"github.com/typical-go/typical-go/examples/generate-docker-compose/pinger"
	"github.com/typical-go/typical-go/pkg/typdocker"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-docker-compose",
	Version: "1.0.0",

	EntryPoint: pinger.Main,

	BuildSequences: []interface{}{
		&typgo.StdBuild{}, // standard build module
	},

	Utility: typdocker.Compose(
		redisRecipe,
	),
}
