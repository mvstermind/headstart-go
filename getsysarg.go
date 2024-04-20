package main

import (
	"fmt"
	"os"
)

func GetSysarg() string {

	if len(os.Args) < 2 {
		fmt.Print("Your project needs a name\n")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Print("you canot have spaces in your project name\n")
		os.Exit(1)
	}
	sysarg := os.Args[1]
	if sysarg == "" {
		fmt.Print("Project name cannot be empty, use -h for help\n")
		os.Exit(1)
	}
	return sysarg
}
