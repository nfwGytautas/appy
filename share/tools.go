package appy_share

import "os/exec"

// Run go fmt on the file
func RunGoFmt(file string) error {
	// Run go fmt on the file
	err := exec.Command("go", "fmt", file).Run()
	if err != nil {
		return err
	}

	return nil
}

// Run goimports on the file
func RunGoImport(file string) error {
	// Run goimports on the file
	err := exec.Command("goimports", "-w", file).Run()
	if err != nil {
		return err
	}

	return nil
}

// Run the full suite of go tools on a file
func RunGoFileTools(file string) error {
	err := RunGoFmt(file)
	if err != nil {
		return err
	}

	err = RunGoImport(file)
	if err != nil {
		return err
	}

	return nil
}

// Run go mod tidy
func RunGoModTidy() error {
	err := exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		return err
	}

	return nil
}
