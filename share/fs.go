package appy_share

import "os"

// Ensure directory exists, if it doesn't create it
func EnsureDirectory(dir string) error {
	// Check if directory exists
	if _, err := os.Stat(dir); err == nil {
		return nil
	}

	// Create directory
	err := os.Mkdir(dir, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Ensure a file exists, if it doesn't create it
func EnsureFile(file string) error {
	// Check if file exists
	if _, err := os.Stat(file); err == nil {
		return nil
	}

	// Create file
	_, err := os.Create(file)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// Clear a file's contents
func ClearFile(file string) error {
	// Open file for writing
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

// Ensure a file exists and is empty
func EnsureCleanFile(file string) error {
	err := EnsureFile(file)
	if err != nil {
		return err
	}

	err = ClearFile(file)
	if err != nil {
		return err
	}

	return nil
}
