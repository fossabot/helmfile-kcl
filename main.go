package main

import (
	"os"

	"kcl-lang.io/helmfile-kcl/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
