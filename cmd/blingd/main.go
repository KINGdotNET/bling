package main

import (
	"os"

	"github.com/kingdotnet/bling/cmd/blingd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
