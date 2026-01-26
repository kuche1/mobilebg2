package configuration

import (
	"mobilebg2/define"
	"os"
	"path/filepath"
)

func getConfigFile() string {
	path, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	path = filepath.Join(path, define.CONFIG_FILE_NAME)
	return path
}
