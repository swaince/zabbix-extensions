package main

import (
	"fmt"
	"github.com/swaince/zabbix-tools/common"
	"github.com/swaince/zabbix-tools/fetch"
	"github.com/swaince/zabbix-tools/template"
	"strings"
)

func generateGetBaseObject() {
	base := fetch.FetchDoc("https://www.zabbix.com/documentation/6.0/zh/manual/api/reference_commentary", 1)
	baseObj := &fetch.StructObject{
		StructName: "GetBaseObject",
		Fields:     base,
	}

	obj := CommonObject("Host", "host", 2)
	obj.StructName = "TagObject"

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: "base",
		Structs: []*fetch.StructObject{
			baseObj,
			obj,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateHost() {
	Common("Host", "host")
}

func generateHostGroup() {
	Common("HostGroup", "hostgroup")
}

func generateHostInterface() {
	Common("HostInterface", "hostinterface")

}

func generateDiscoverRole() {
	Common("DiscoverRole", "discoveryrule")
}

func generateItem() {
	Common("Item", "item")
}

func generateItemPrototype() {
	Common("ItemPrototype", "itemprototype")
}

func generateValueMap() {

	target := "ValueMap"
	path := "valuemap"

	obj := CommonObject(target, path, 1)
	obj.StructName = "ValueMapMappingObject"

	Common(target, path, obj)
}

func generateTemplate() {
	Common("Template", "template")
}

func generateMacro() {
	target := "Macro"
	path := "usermacro"

	fs1 := fetch.FetchDoc(common.GetDocUrl(path+"/object"), 1)
	fs := &fetch.FieldObject{
		NewKey:  "Globalmacroid",
		NewType: "string",
		NewDesc: "(readonly) 全局宏的ID.",
		Tag:     "`json:\"globalmacroid,omitempty\"`",
	}

	fs1 = append(fs1, fs)
	obj1 := &fetch.StructObject{
		StructName: target,
		Fields:     fs1,
	}

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateHttpTest() {
	Common("HttpTest", "httptest")
}

func generateEvent() {
	target := "Event"
	path := "event"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateHistory() {
	target := "History"
	path := "history"
	obj1 := CommonObject(target, path, 4)
	obj2 := CommonGet(target, path, 0)

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))

}

func generateAlert() {
	target := "Alert"
	path := "alert"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))

}

func generateGraph() {
	Common("Graph", "graph")
}

func generateGraphPrototype() {
	Common("GraphPrototype", "graphprototype")
}

func generateGraphItem() {
	target := "GraphItem"
	path := "graphitem"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateMediaType() {
	Common("MediaType", "mediatype")
}

func generateAuditLog() {
	target := "AuditLog"
	path := "auditlog"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func Common(target, path string, adds ...*fetch.StructObject) {
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonCreate(target, path, 0)
	obj3 := CommonUpdate(target, path, 0)
	obj4 := CommonGet(target, path, 0)
	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
			obj3,
			obj4,
		},
	}

	if adds != nil {
		packObj.Structs = append(packObj.Structs, adds...)
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateTrigger() {
	Common("Trigger", "trigger")
}

func generateTriggerPrototype() {
	Common("TriggerPrototype", "triggerprototype")
}

func CommonObject(target, path string, index int) *fetch.StructObject {
	fs := fetch.FetchDoc(common.GetDocUrl(path+"/object"), index)
	return &fetch.StructObject{
		StructName: target,
		Fields:     fs,
	}
}

func CommonCreate(target, path string, index int) *fetch.StructObject {
	cfs := fetch.FetchDoc(common.GetDocUrl(path+"/create"), index)
	return &fetch.StructObject{
		StructName: target + "CreateParam",
		Fields:     cfs,
		Parent:     target,
	}
}

func CommonUpdate(target, path string, index int) *fetch.StructObject {
	ufs := fetch.FetchDoc(common.GetDocUrl(path+"/update"), 0)
	return &fetch.StructObject{
		StructName: target + "UpdateParam",
		Fields:     ufs,
		Parent:     target,
	}
}

func CommonGet(target, path string, index int) *fetch.StructObject {
	gfs := fetch.FetchDoc(common.GetDocUrl(path+"/get"), index, fetch.ExcludeFields...)
	return &fetch.StructObject{
		StructName: target + "GetParam",
		Fields:     gfs,
		Parent:     "GetBaseObject",
	}
}
