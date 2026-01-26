package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	// price (EUR)

	PriceMin int
	PriceMax int
}

func newConfigWithDefaults() *Config {
	return &Config{
		PriceMin: 10_000, //1_852 // bgn 3_600
		PriceMax: 15_000, //3_852 // 4_632 // bgn 9_000
	}
}

func NewConfig() *Config {
	configFilePath := getConfigFile()

	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("Config file does not exist, creating: %v\n", configFilePath)
		config := newConfigWithDefaults()
		config.save(configFilePath)
	}

	dataAsBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Using defaults: Could not open config file `%v`: %v", configFilePath, err)
		dataAsBytes = []byte("{}")
	} else {
		fmt.Printf("Using config file: %v\n", configFilePath)
	}

	dataAsStr := string(dataAsBytes)

	decoder := json.NewDecoder(strings.NewReader(dataAsStr))
	decoder.DisallowUnknownFields()

	config := newConfigWithDefaults()
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	return config
}

func (self *Config) save(file string) {
	data, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		panic(err)
	}
}
