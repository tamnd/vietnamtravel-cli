// Package cli assembles the vietnamtravel command tree from the vietnamtravel
// domain on top of the any-cli/kit framework.
package cli

import (
	"github.com/tamnd/any-cli/kit"
	"github.com/tamnd/vietnamtravel-cli/vietnamtravel"
)

// Build metadata, set via -ldflags at release time.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// NewApp assembles the kit application from the vietnamtravel domain. The
// domain's Register installs the client factory and every operation, so the
// binary and a host (ant, which blank-imports the package) share one source of
// truth. kit.Run turns the App into the CLI, plus the serve and mcp surfaces and
// the typed-error-to-exit-code mapping.
//
// To add a command, declare it in vietnamtravel/domain.go with kit.Handle and it
// appears here automatically. Reach for app.AddCommand only for a verb that does
// not fit the emit-records shape, the way version does below.
func NewApp() *kit.App {
	id := vietnamtravel.Domain{}.Info().Identity
	id.Version = Version

	app := kit.New(id)
	(vietnamtravel.Domain{}).Register(app)
	app.AddCommand(newVersionCmd())
	return app
}
