package helper

import (
	"errors"
	"github.com/jaskaur18/InstaPurity-whitelist-unfollower/types"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

var vp *viper.Viper

func CreateConfigFile(path string) error {
	// Create test double data
	dummyData := types.Config{
		Username:     "username",
		Password:     "***",
		Interval:     20,
		Limit:        500,
		WhiteListTXT: "~/whitelist.txt",
	}

	// Marshal test double data to YAML format
	dummyDataBytes, err := yaml.Marshal(dummyData)
	if err != nil {
		return errors.New("error while marshaling dummy data")
	}

	// Write test double data to the config file
	err = os.WriteFile(path, dummyDataBytes, 0644)
	if err != nil {
		return errors.New("error while writing dummy data to config file")
	}

	return nil
}

var CONFIG *types.Config

func LoadConfig() (*types.Config, error) {
	// Check if the config file exists
	configPath := path.Join(HomePath, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err := CreateConfigFile(configPath)
		if err != nil {
			log.Printf("Error creating config file: %v", err)
			return nil, err
		}
	}

	vp = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(HomePath)

	config := &types.Config{}

	err := vp.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return config, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	CONFIG = config

	return config, nil
}
