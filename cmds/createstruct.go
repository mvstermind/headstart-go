package cmds

import (
	"fmt"
	"os"

	"github.com/mastermind/headstartgo/cmds/templates"
)

func MakeDirs(c string, root string) {
	// add more stuff in here
	if len(root) > 2 {
		err := fmt.Errorf("Pass only one value as your project name")
		fmt.Printf(err.Error())
		os.Exit(1)
	}

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
