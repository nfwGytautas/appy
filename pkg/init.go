package appy_pkg

import (
	"fmt"
	"os"
)

func Init(name, language string) error {
	// Check if not already initialized
	if _, err := os.Stat("appy.yaml"); err == nil {
		return fmt.Errorf("appy project already initialized in the current directory")
	}

	if !isLanguageSupported(language) {
		return fmt.Errorf("language %v not supported", language)
	}

	fmt.Printf("Initializing project: %v using %v\n", name, language)

	config := AppyConfig{
		ProjectName: name,
		Language:    language,
	}

	return WriteConfig(&config)
}

func isLanguageSupported(language string) bool {
	return language == "go"
}
