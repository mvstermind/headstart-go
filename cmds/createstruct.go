package cmds

func MakeDirs(c string) {
	switch {
	case c == "Flat":
		MakeFlatTemp()
	case c == "Layared":
		MakeLayared()
	case c == "Domain-Driven":
		MakeDomainD()
	case c == "Clean Architecture":
		MakeClean()
	case c == "Modular":
		MakeModular()
	}

}
