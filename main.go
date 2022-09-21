package main

import (
	"github.com/Joffref/wifi-cli/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
