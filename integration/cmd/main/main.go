package main

import (
	"github.com/bhavik402/remidners-api-go/integration/pkg"
	"github.com/pterm/pterm"
)

func main() {
	printTitle()

	// todo: implement this in an CI run, but in an isolated container
	logger := pterm.DefaultLogger.
		WithLevel(pterm.LogLevelTrace)

	err := pkg.RunAllPostTests(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = pkg.RunAllGetTests(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	result, err := pkg.RunAllPutTests()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(*result)

	//delete records after completion no need if this will isolated in a container
}

func printTitle() {
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightYellow)).WithFullWidth().Println("Running Integration Tests")
	pterm.Println()
}
