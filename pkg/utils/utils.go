package utils

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/isumitbanik/remixer/pkg/config"
)

func Base64Encode(data []byte) string {
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(data))
	return base64Encoded
}

func ReadFileContent(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	config.CheckError(err)
	return content, nil
}
func CopyToClipboard(content string) {
	err := clipboard.WriteAll(content)
	config.CheckError(err)
}
