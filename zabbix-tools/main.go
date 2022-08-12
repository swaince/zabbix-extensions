package main

import (
	"github.com/swaince/zabbix-tools/common"
	"github.com/swaince/zabbix-tools/fetch"
	"github.com/swaince/zabbix-tools/template"
)

func main() {

	base := fetch.FetchDoc("https://www.zabbix.com/documentation/6.0/zh/manual/api/reference_commentary", 1)
	getBase := &fetch.StructObject{
		StructName: "GetBaseObject",
		Fields:     base,
	}

	host := fetch.FetchDoc(common.GetDocUrl("host/object"), 0, fetch.ExcludeFields...)
	hostObj := &fetch.StructObject{
		StructName: "Host",
		Fields:     host,
	}

	hostGet := fetch.FetchDoc(common.GetDocUrl("host/get"), 0, fetch.ExcludeFields...)

	hostGetObj := &fetch.StructObject{
		Fields:     hostGet,
		StructName: "HostGetParam",
		Parent:     "GetBaseObject",
	}

	p := &fetch.PackageObject{
		Package: "test",
		Structs: []*fetch.StructObject{getBase, hostObj, hostGetObj},
	}

	template.Write(p, "model/host.go")
}
