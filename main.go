package main

/*	License: GPLv3
	Copyright: 2025
	Description: Abg is a wrapper around multiple package managers to install packages and run commands inside a managed packets.
*/

import (
	"embed"
	"os"

	"github.com/MiracleOS-Team/abg/v2/cmd"
	"github.com/MiracleOS-Team/abg/v2/core"
	"github.com/MiracleOS-Team/abg/v2/cmdr"
)

var Version = "development"

//go:embed locales/*.yml
var fs embed.FS
var abg *cmdr.App

func main() {
	core.NewStandardApx()

	abg = cmd.New(Version, fs)

	// check if root, exit if so
	if core.RootCheck(false) {
		cmdr.Error.Println(abg.Trans("abg.errors.noRoot"))
		os.Exit(1)
	}

	// root command
	root := cmd.NewRootCommand(Version)
	abg.CreateRootCommand(root, abg.Trans("abg.msg.help"), abg.Trans("abg.msg.version"))

	msgs := cmdr.UsageStrings{
		Usage:                abg.Trans("abg.msg.usage"),
		Aliases:              abg.Trans("abg.msg.aliases"),
		Examples:             abg.Trans("abg.msg.examples"),
		AvailableCommands:    abg.Trans("abg.msg.availableCommands"),
		AdditionalCommands:   abg.Trans("abg.msg.additionalCommands"),
		Flags:                abg.Trans("abg.msg.flags"),
		GlobalFlags:          abg.Trans("abg.msg.globalFlags"),
		AdditionalHelpTopics: abg.Trans("abg.msg.additionalHelpTopics"),
		MoreInfo:             abg.Trans("abg.msg.moreInfo"),
	}
	abg.SetUsageStrings(msgs)

	// commands
	stacks := cmd.NewStacksCommand()
	root.AddCommand(stacks)

	subsystems := cmd.NewSubSystemsCommand()
	root.AddCommand(subsystems)

	pkgManagers := cmd.NewPkgManagersCommand()
	root.AddCommand(pkgManagers)

	runtimeCmds := cmd.NewRuntimeCommands()
	root.AddCommand(runtimeCmds...)

	// run the app
	err := abg.Run()
	if err != nil {
		cmdr.Error.Println(err)
		os.Exit(1)
	}
}
