package main

import (
	"github.com/swaince/zabbix-tools/fetch"
	"github.com/swaince/zabbix-tools/template"
)

func main() {
	c := fetch.FetchDoc("host/get", 0)
	c.StructName = "HostGetParam"
	c.Parent = "Host"
	c.Package = "test"

	//time.Sleep(5 * time.Second)
	template.Render(c)
}
