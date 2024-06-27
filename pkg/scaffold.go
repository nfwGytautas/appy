package appy_pkg

import (
	"fmt"

	appy_share "github.com/nfwGytautas/appy/share"
	appy_templates "github.com/nfwGytautas/appy/share/templates"
)

func Scaffold() error {
	config, err := appy_share.ReadConfig()
	if err != nil {
		return err
	}

	fmt.Printf("Scaffolding project: %v\n", config.ProjectName)

	// Check that directories exist
	err = checkBaseDirs(config)
	if err != nil {
		return err
	}

	err = checkHttp(config)
	if err != nil {
		return err
	}

	err = checkGoMod(config)
	if err != nil {
		return err
	}

	err = checkApiConnector(config)
	if err != nil {
		return err
	}

	err = checkHooksFile(config)
	if err != nil {
		return err
	}

	err = checkMain(config)
	if err != nil {
		return err
	}

	return nil
}

// Check that base directories exist if not create them
func checkBaseDirs(config *appy_share.AppyConfig) error {
	fmt.Println("Checking directories...")

	err := appy_share.EnsureDirectory("api")
	if err != nil {
		return err
	}

	err = appy_share.EnsureDirectory("api/handlers")
	if err != nil {
		return err
	}

	if config.Http != nil {
		err := appy_share.EnsureDirectory("api/http")
		if err != nil {
			return err
		}
	}

	if config.Storage != nil {
		err := appy_share.EnsureDirectory("driver")
		if err != nil {
			return err
		}

		err = appy_share.EnsureDirectory("migrations")
		if err != nil {
			return err
		}

		err = appy_share.EnsureDirectory("queries")
		if err != nil {
			return err
		}
	}

	if config.Jobs != nil {
		err := appy_share.EnsureDirectory("jobs")
		if err != nil {
			return err
		}
	}

	err = appy_share.EnsureDirectory("models")
	if err != nil {
		return err
	}

	return nil
}

// Check that handlers and endpoints exist
func checkHttp(config *appy_share.AppyConfig) error {
	if config.Http == nil {
		// Nothing to do
		return nil
	}

	const wiringFile = "api/http/http.go"

	fmt.Println("Checking HTTP endpoints...")

	err := generateHandlerInputsOutputs(config)
	if err != nil {
		return err
	}

	// Create endpoint files
	for _, httpEndpoint := range config.Http.Endpoints {
		err := generateHandler(config, httpEndpoint)
		if err != nil {
			return err
		}

		// Create file
		fileName := fmt.Sprintf("api/http/%v.go", httpEndpoint.Name)

		fmt.Printf("  - %v \n", fileName)

		err = appy_share.EnsureFile(fileName)
		if err != nil {
			return err
		}

		// Clear contents and write boilerpate
		data := struct {
			Config   *appy_share.AppyConfig
			Http     *appy_share.Http
			Endpoint *appy_share.Endpoint
		}{
			Config:   config,
			Http:     config.Http,
			Endpoint: &httpEndpoint,
		}

		err = appy_share.ClearFile(fileName)
		if err != nil {
			return err
		}

		err = appy_templates.WriteTemplateToFile(fileName, appy_templates.EndpointTemplate, data)
		if err != nil {
			return err
		}

		err = appy_share.RunGoFileTools(fileName)
		if err != nil {
			return err
		}
	}

	// Create base wiring file
	err = appy_share.EnsureFile(wiringFile)
	if err != nil {
		return err
	}

	err = appy_share.ClearFile(wiringFile)
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile(wiringFile, appy_templates.HttpWiringTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools(wiringFile)
	if err != nil {
		return err
	}

	return nil
}

// Check that main.go exists
func checkMain(config *appy_share.AppyConfig) error {
	fmt.Println("Checking main.go...")
	err := appy_share.EnsureFile("main.go")
	if err != nil {
		return err
	}

	// Clear contents
	err = appy_share.ClearFile("main.go")
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile("main.go", appy_templates.MainTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools("main.go")
	if err != nil {
		return err
	}

	return nil
}

// Check that go.mod exist
func checkGoMod(config *appy_share.AppyConfig) error {
	fmt.Println("Checking go.mod...")
	err := appy_share.EnsureFile("go.mod")
	if err != nil {
		return err
	}

	// Clear contents
	err = appy_share.ClearFile("go.mod")
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile("go.mod", appy_templates.GoModTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoModTidy()
	if err != nil {
		return err
	}

	return nil
}

func generateHandler(config *appy_share.AppyConfig, endpoint appy_share.Endpoint) error {
	handlerFile := fmt.Sprintf("api/handlers/http_%v.go", endpoint.Name)

	// Don't do anything if it already exists
	if appy_share.FileExists(handlerFile) {
		return nil
	}

	err := appy_share.EnsureFile(handlerFile)
	if err != nil {
		return err
	}

	err = appy_share.ClearFile(handlerFile)
	if err != nil {
		return err
	}

	data := struct {
		Config   *appy_share.AppyConfig
		Endpoint *appy_share.Endpoint
	}{
		Config:   config,
		Endpoint: &endpoint,
	}

	err = appy_templates.WriteTemplateToFile(handlerFile, appy_templates.HandlerTemplate, data)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools(handlerFile)
	if err != nil {
		return err
	}

	return nil
}

// Generates the handler inputs and outputs file
func generateHandlerInputsOutputs(config *appy_share.AppyConfig) error {
	const handlerFile = "api/handlers/appy_generated.go"

	err := appy_share.EnsureFile(handlerFile)
	if err != nil {
		return err
	}

	err = appy_share.ClearFile(handlerFile)
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile(handlerFile, appy_templates.HandlerImplementationTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools(handlerFile)
	if err != nil {
		return err
	}

	return nil
}

// Check that the api connector exists
func checkApiConnector(config *appy_share.AppyConfig) error {
	fmt.Println("Checking API connector...")

	const connectorFile = "api/connector.go"

	err := appy_share.EnsureFile(connectorFile)
	if err != nil {
		return err
	}

	err = appy_share.ClearFile(connectorFile)
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile(connectorFile, appy_templates.ConnectorTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools(connectorFile)
	if err != nil {
		return err
	}

	return nil
}

func checkHooksFile(config *appy_share.AppyConfig) error {
	fmt.Println("Checking hooks.go...")

	const hooksFile = "hooks.go"

	// Don't do anything if it already exists
	if appy_share.FileExists(hooksFile) {
		return nil
	}

	err := appy_share.EnsureFile(hooksFile)
	if err != nil {
		return err
	}

	err = appy_share.ClearFile(hooksFile)
	if err != nil {
		return err
	}

	err = appy_templates.WriteTemplateToFile(hooksFile, appy_templates.HooksTemplate, config)
	if err != nil {
		return err
	}

	err = appy_share.RunGoFileTools(hooksFile)
	if err != nil {
		return err
	}

	return nil
}
