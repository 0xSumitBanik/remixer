/*
Copyright © 2022 Sumit Banik <sumitbanik02@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "remixer",
	Short:   "Remixer generates shareable Remix IDE links for your Solidity Smart Contracts",
	Long:    `Remixer generates shareable Remix IDE links for your Solidity Smart Contracts`,
	Version: "0.0.1",
	
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
