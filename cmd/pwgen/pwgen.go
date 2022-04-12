package main

import (
	"flag"
	"fmt"

	"github.com/mbwk/pwgen"
)

func main() {
	length := flag.Int("length", 12, "length of the generated password")
	nouppercase := flag.Bool("nouppercase", false, "exclude uppercase A-Z characters")
	nolowercase := flag.Bool("nolowercase", false, "exclude lowercase a-z characters")
	number := flag.Int("number", 2, "exact number of numeric 0-9 characters")
	special := flag.Int("special", 2, "exact number of special @%!?*^& characters")

	flag.Parse()

	password, err := pwgen.GeneratePassword(*length, !(*nouppercase), !(*nolowercase), *number, *special)

	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	fmt.Println("Password:", password)
}
