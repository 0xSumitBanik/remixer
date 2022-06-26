/*
Copyright © 2022 Sumit Banik <sumitbanik02@gmail.com>

*/
package cmd

import (
	"github.com/isumitbanik/remixer/pkg/cli"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Remixer generate generates the shareable Remix IDE URL for sharing the smart contract",
	Long:    "\nShare Solidity Code Snippets as a shareable Remix IDE URL",
	Run:     cli.Generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	var File string
	var Theme string
	var Plugins string
	var AutoCompile string
	var EVMVersion string
	var Optimize string
	var Language string
	
	generateCmd.Flags().StringVarP(&File, "file", "f", "", "Smart Contract File for generating the shareable URL")
	generateCmd.Flags().StringVarP(&Theme, "theme", "t", "Dark", "Theme for the Remix IDE")
	generateCmd.Flags().StringVarP(&Plugins, "plugins", "p", "", "List of Plugins to be activated for the Remix IDE (seperated by commas)")
	generateCmd.Flags().StringVarP(&AutoCompile, "auto-compile", "", "true", "Toggle flag for turning on/off the auto compiler")
	generateCmd.Flags().StringVarP(&Optimize, "optimize", "", "false", "Toggle flag for turning on/off Solidity Optimizer")
	generateCmd.Flags().StringVarP(&EVMVersion, "evm-version", "", "default", "Version of the Ethereum Virtual Machine")
	generateCmd.Flags().StringVarP(&Language, "language", "l", "Solidity", "Supported Values: Solidity or Yul")
}
