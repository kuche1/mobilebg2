package configuration

import "mobilebg2/define"

func getConfigFile() string {
	// TODO: use the OS-specific folder
	configFilePath := define.CONFIG_FILE_NAME
	return configFilePath
}
