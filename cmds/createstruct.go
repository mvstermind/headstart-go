package cmds

import (
	"github.com/mastermind/headstartgo/cmds/templates"
)

func MakeDirs(c string) {
	switch {
	case c == "Flat":
		templates.FlatTemp()
	case c == "Layared":
		templates.Layared()
	case c == "Domain-Driven":
		templates.DomainDriven()
	case c == "Clean Architecture":
		templates.Clean()
	case c == "Modular":
		templates.Modular()

	}
}
