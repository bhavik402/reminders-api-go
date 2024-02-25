package main

import (
	"fmt"

	"github.com/bhavik402/remidners-api-go/integration/pkg"
)

func main() {
	fmt.Println("Integration tests")

	// todo: implement this in an CI run, but in an isolated container

	err := pkg.PostAllRecords()
	if err != nil {
		panic(err)
	}

	err = pkg.GetAllRecords()
	if err != nil {
		panic(err)
	}

	//delete records after completion no need if this will isolated in a container
}
