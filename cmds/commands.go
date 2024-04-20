package cmds

import "fmt"

func Commands(r string) {
	switch {
	case r == "":
		fmt.Println("Type your root directory project name (or -h for help)")
	case r == "-h":
		fmt.Println("i'll add commands in here")
	}
}
