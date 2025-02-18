package main

import (
	"os"

	"github.com/athul/shelby/mods"
	"github.com/talal/go-bits/color"
)

func main() {
	if len(os.Args) > 1 {
		color.Fprint(os.Stderr, color.Red, "Error.....\n")
		os.Exit(1)
	}
	os.Stdout.Write([]byte("\n" + mods.Info() + "\n"))
}
