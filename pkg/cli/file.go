package cli

import (
	"errors"

	"github.com/isumitbanik/remixer/pkg/config"
	"github.com/isumitbanik/remixer/pkg/utils"
)

func ProcessFileInput(filePath string) (string, error) {
	if filePath != "" {
		fileContent, err := utils.ReadFileContent(filePath)
		config.CheckError(err)

		base64Encoded := utils.Base64Encode(fileContent)
		return base64Encoded, nil
	} else {
		err := errors.New("specify a filename to generate shareable link")
		return "", err

	}
}
