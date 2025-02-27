package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	cfgFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(cfgFile)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) SetUser(name string) error {
	// Add current_user_name to Config Struct and write it to config json file
	c.CurrentUserName = name
	// fmt.Printf("%+v\n", c)
	return write(c)

}
