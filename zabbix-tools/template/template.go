package template

import (
	"github.com/swaince/zabbix-tools/fetch"
	"html/template"
	"os"
)

const (
	StructTemplate = `
package {{ .Package }}

type {{ .StructName }} struct {
{{ if .Parent }}
	{{ .Parent }}
{{ end }}
{{range .Fields }}
    /**
    {{ .NewDesc }}
    */
	{{ .NewKey}} {{ .NewType }}
{{end}}
}`
)

func Render(c *fetch.ClassObject) {
	tmpl, _ := template.New(c.StructName).Parse(StructTemplate)
	err := tmpl.Execute(os.Stdout, c)

	if err != nil {
		panic(err)
	}

}
