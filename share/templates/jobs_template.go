package appy_templates

const JobsFileTemplate = `
package {{ .ProjectName }}_jobs

import (
	appy_jobs "github.com/nfwGytautas/appy-go/jobs"
)

{{ GeneratedNotice }}

func Setup() {
	{{ range .Jobs }}
	appy_jobs.Get().Add(appy_jobs.JobOptions{
		Job: {{ .Name }},
		Tick: {{ .Tick }} * time.Second,
		Type: {{ if .Independent }} appy_jobs.JobTypePersistent {{ else }} appy_jobs.JobTypePooled {{ end }},
	})
	{{ end }}
}

`

const JobImplementationTemplate = `
package {{ .Config.ProjectName }}_jobs

func {{ .Job.Name }}() {
	// Add your job logic here
}

`
