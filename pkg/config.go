package appy_pkg

import (
	"os"

	"gopkg.in/yaml.v2"
)

type AppyConfig struct {
	ProjectName string `yaml:"project"`
	Language    string `yaml:"language"`

	Migrations string `yaml:"migrations"`
	Quries     string `yaml:"queries"`
	Api        string `yaml:"api"`
}

func WriteConfig(config *AppyConfig) error {
	// Open a file for writing
	file, err := os.Create("appy.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the data to YAML and write it to the file
	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func ReadConfig() (*AppyConfig, error) {
	var config AppyConfig

	// Read YAML data from a file
	yamlData, err := os.ReadFile("appy.yaml")
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML data into the Person struct
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
