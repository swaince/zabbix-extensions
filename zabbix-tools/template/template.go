package template

import (
	"github.com/swaince/zabbix-tools/fetch"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	PackageTemplate = `package {{ .Package }}

{{ range .Structs }}
type {{ .StructName }} struct {
    {{ if .Parent }}
	{{ .Parent }}
    {{ end }}
    {{- range .Fields }}
    /**
    {{ .NewDesc | unescaped }}
    */
	{{ .NewKey}} {{ .FinalType }} {{ .Tag | unescaped }}
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

func Write(packObj *fetch.PackageObject, fileName string) {

	dir := filepath.Dir(fileName)
	os.MkdirAll(dir, fs.ModeDir)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	Render(packObj, f)
}
