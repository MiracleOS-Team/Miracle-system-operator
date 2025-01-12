package cmd

import (
	"embed"

	"github.com/MiracleOS-Team/abg/cmdr"
)

var abg *cmdr.App

func New(version string, fs embed.FS) *cmdr.App {
	abg = cmdr.NewApp("abg", version, fs)
	return abg
}

func NewRootCommand(version string) *cmdr.Command {
	root := cmdr.NewCommand(
		"abg",
		abg.Trans("abg.description"),
		abg.Trans("abg.description"),
		nil,
	)
	root.Version = version

	return root
}
