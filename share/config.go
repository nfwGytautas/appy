package appy_share

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type AppyConfig struct {
	ProjectName string    `yaml:"project"`
	Domain      string    `yaml:"domain"`
	Version     string    `yaml:"version"`
	Storage     []Storage `yaml:"storage"`
	Jobs        []Job     `yaml:"jobs"`
	Http        *Http     `yaml:"http"`
}

type Storage struct {
	Type string `yaml:"type"`
}

type Job struct {
	Name        string `yaml:"name"`
	Tick        uint32 `yaml:"tick"`
	Independent bool   `yaml:"independent"`
}

type Http struct {
	Address   string     `yaml:"address"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Name   string     `yaml:"name"`
	Method string     `yaml:"method"`
	Path   string     `yaml:"path"`
	Args   []Argument `yaml:"args"`
}

type Argument struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func ReadConfig() (*AppyConfig, error) {
	var result *AppyConfig

	// Check if appy.yaml exsits
	if _, err := os.Stat("appy.yaml"); err != nil {
		return nil, fmt.Errorf("appy.yaml doesn't exist")
	}

	// Read appy.yaml file
	yamlFile, err := os.ReadFile("appy.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
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
