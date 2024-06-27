package appy_templates

import (
	"os"
	"text/template"
)

func WriteTemplateToFile(file string, templateContent string, arg any) error {
	// Open a file for writing
	f, err := os.OpenFile(file, os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	// Write the contents
	t := template.Must(template.New("appy_template").Funcs(funcMap).Parse(templateContent))

	// Execute the template with data
	err = t.Execute(f, arg)
	if err != nil {
		return err
	}

	return nil
}
