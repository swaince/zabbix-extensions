package template

import (
	"github.com/swaince/zabbix-tools/fetch"
	"html/template"
	"io"
)

const (
	PackageTemplate = `
package {{ .Package }}

{{ range .Structs }}
type {{ .StructName }} struct {
    {{ if .Parent }}
	{{ .Parent }}
    {{ end }}
    {{- range .Fields }}
    /**
    {{ .NewDesc | unescaped }}
    */
	{{ .NewKey}} {{ .NewType }} {{ .Tag | unescaped }}
    {{ end }}
}
{{ end }}`
)

func Render(c *fetch.PackageObject, wr io.Writer) {
	tmpl, _ := template.New(c.Package).Funcs(template.FuncMap{
		"unescaped": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).Parse(PackageTemplate)
	err := tmpl.Execute(wr, c)

	if err != nil {
		panic(err)
	}

}
