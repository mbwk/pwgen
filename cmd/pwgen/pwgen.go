package main

import (
	"fmt"

	"github.com/mbwk/pwgen"
)

func main() {
	password, err := pwgen.GeneratePassword(12, true, true, 2, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Password:", password)
}
