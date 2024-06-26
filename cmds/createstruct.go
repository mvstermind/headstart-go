package cmds

import (
	"os"
	"strings"

	"github.com/mastermind/headstartgo/cmds/templates"
)

func MakeDirs(c string, root string) {

	i := strings.Count(root, " ")
	if i > 0 {
		os.Exit(1)
	}
	switch {
	case c == "Flat":
		templates.FlatTemp(root)
	case c == "Layared":
		templates.Layared(root)
	case c == "Domain-Driven":
		templates.DomainDriven(root)
	case c == "Clean Architecture":
		templates.Clean(root)
	case c == "Modular":
		templates.Modular(root)
	}
}
