package cli

import (
	"fmt"

	"github.com/isumitbanik/remixer/pkg/config"
	"github.com/isumitbanik/remixer/pkg/utils"

	"github.com/spf13/cobra"
)

func Generate(cmd *cobra.Command, args []string) {
	filePath, err := cmd.Flags().GetString("file")
	config.CheckError(err)

	base64EncodedData, err := ProcessFileInput(filePath)
	config.CheckError(err)
	dataMap, err := ParseAdditionalFlags(cmd)
	config.CheckError(err)
	urlOutput := GenerateFinalOutput(dataMap, base64EncodedData)

	utils.CopyToClipboard(urlOutput)

	fmt.Println("Shareable URL copied to clipboard. ✅")
}

func ParseAdditionalFlags(cmd *cobra.Command) (map[string]string, error) {

	flagsMap := make(map[string]string)
	theme, err := cmd.Flags().GetString("theme")
	config.CheckError(err)
	plugins, err := cmd.Flags().GetString("plugins")
	config.CheckError(err)
	autoCompile, err := cmd.Flags().GetString("auto-compile")
	config.CheckError(err)
	optimize, err := cmd.Flags().GetString("optimize")
	config.CheckError(err)
	language, err := cmd.Flags().GetString("language")
	config.CheckError(err)

	flagsMap["theme"] = theme
	flagsMap["plugins"] = plugins
	flagsMap["autoCompile"] = autoCompile
	flagsMap["optimize"] = optimize
	flagsMap["language"] = language

	return flagsMap, nil
}

func GenerateFinalOutput(dataMap map[string]string, base64EncodedData string) string {
	url := config.REMIX_URL + config.CODE_ARG +
		base64EncodedData + config.AND + config.THEME_ARG + dataMap["theme"] +
		config.AND + config.AUTOCOMPILE_ARG + dataMap["autoCompile"] + config.AND +
		config.OPTIMIZE_ARG + dataMap["optimize"]
	return url
}
