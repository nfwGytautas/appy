package appy_templates

const EndpointTemplate = `
package {{ .Config.ProjectName }}_http

import (
	{{ .Config.ProjectName }}_handlers "{{ .Config.Domain }}/{{ .Config.ProjectName }}/api/handlers"
)

{{ GeneratedNotice }}

func {{ .Endpoint.Name | Title }}Handler(c *gin.Context) {
	ctx := c.Request.Context()

	// Send to the handler
	args := {{ .Config.ProjectName }}_handlers.{{ .Endpoint.Name | Title }}Args{}

	result := {{ .Config.ProjectName}}_handlers.{{ .Endpoint.Name | Title }}Handler(ctx, args)

	// Format result
	if (result.Body != nil) {
		c.JSON(result.StatusCode, result.Body)
	} else {
		c.Status(result.StatusCode)
	}

	return
}

`
