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

	// engine

	EngineTypeBlacklist []string
}

func newConfigWithDefaults() *Config {
	return &Config{
		PriceMin: 10_000, //1_852 // bgn 3_600
		PriceMax: 15_000, //3_852 // 4_632 // bgn 9_000

		EngineTypeBlacklist: []string{
			"Дизелов",
			"Електрически",
		},
	}
}

func NewConfig() *Config {
	configFilePath := getConfigFile()

	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("Config file does not exist, creating\n")
		config := newConfigWithDefaults()
		config.save(configFilePath)
	}

	dataAsBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Using defaults: Could not open config file `%v`: %v", configFilePath, err)
		dataAsBytes = []byte("{}")
	}

	dataAsStr := string(dataAsBytes)

	decoder := json.NewDecoder(strings.NewReader(dataAsStr))
	decoder.DisallowUnknownFields()

	config := newConfigWithDefaults()
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	config.save(configFilePath)

	return config
}

func (self *Config) save(file string) {
	fmt.Printf("Saving config file: %v\n", file)

	data, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		panic(err)
	}
}
