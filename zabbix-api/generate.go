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
	obj1 := &fetch.StructObject{
		StructName: "GetBaseObject",
		Fields:     base,
	}

	m1 := GetDefaultMappings()
	obj1.ParseStructType(m1)

	obj2 := CommonObject("Host", "host", 2)
	obj2.StructName = "TagObject"

	f1 := &fetch.FieldObject{
		RawKey:    "name",
		NewKey:    "Name",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "Header 名称.",
		FinalType: "string",
	}

	f2 := &fetch.FieldObject{
		RawKey:    "value",
		NewKey:    "Value",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "Header 值.",
		FinalType: "string",
	}

	f3 := &fetch.FieldObject{
		RawKey:    "http_proxy",
		NewKey:    "HttpProxy",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "HTTP代理连接字符串。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f4 := &fetch.FieldObject{
		RawKey:    "output_format",
		NewKey:    "OutputFormat",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "\t0 - RAW (default)\n1 - JSON\nH如何处理反应。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f5 := &fetch.FieldObject{
		RawKey:    "post_type",
		NewKey:    "PostType",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "0 - RAW (default)\n2 - JSON\n3 - XML\n发布数据体的类型。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f6 := &fetch.FieldObject{
		RawKey:    "posts",
		NewKey:    "Posts",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "HTTP(S)请求体数据。\n仅供HTTP代理项使用",
		FinalType: "string",
	}

	obj3 := &fetch.StructObject{
		StructName: "HttpHeaderObject",
		Fields: []*fetch.FieldObject{
			f1, f2, f3, f4, f5, f6,
		},
	}

	/*
		parameters	-	用户定义参数的根元素。

		仅供脚本项使用。
		name	x	string	参数名称。

		仅供脚本项使用。
		value	-	string	参数值。

		仅供脚本项使用。
	*/
	f7 := &fetch.FieldObject{
		RawKey:    "name",
		NewKey:    "Name",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "参数名称。\n仅供脚本项使用。",
		FinalType: "string",
	}
	f8 := &fetch.FieldObject{
		RawKey:    "value",
		NewKey:    "Value",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "参数值。\n仅供脚本项使用。",
		FinalType: "string",
	}

	obj4 := &fetch.StructObject{
		StructName: "ParameterObject",
		Fields: []*fetch.FieldObject{
			f7, f8,
		},
	}

	f9 := &fetch.FieldObject{
		RawKey:    "name",
		NewKey:    "Name",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "参数名称",
		FinalType: "string",
	}

	f10 := &fetch.FieldObject{
		RawKey:    "value",
		NewKey:    "Value",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "参数值",
		FinalType: "string",
	}

	f11 := &fetch.FieldObject{
		RawKey:    "request_method",
		NewKey:    "RequestMethod",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "请求方法。\n\n仅供HTTP代理项使用。\n0 - GET (default)\n1 - POST\n2 - PUT\n3 - HEAD",
		FinalType: "string",
	}

	f12 := &fetch.FieldObject{
		RawKey:    "retrieve_mode",
		NewKey:    "RetrieveMode",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "应该存储响应的哪一部分。\n\n仅供HTTP代理项使用。\n0 - BODY (default)\n1 - HEADERS\n2 - BOTH",
		FinalType: "string",
	}

	f13 := &fetch.FieldObject{
		RawKey:    "ssl_cert_file",
		NewKey:    "SslCertFile",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "Public SSL密钥文件路径。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f14 := &fetch.FieldObject{
		RawKey:    "ssl_key_file",
		NewKey:    "SslKeyFile",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "Private SSL Key文件路径。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f15 := &fetch.FieldObject{
		RawKey:    "ssl_key_password",
		NewKey:    "SslKeyPassword",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "SSL密钥文件的密码。\n仅供HTTP代理项使用。",
		FinalType: "string",
	}

	f16 := &fetch.FieldObject{
		RawKey:    "status_codes",
		NewKey:    "StatusCodes",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "用逗号分隔的所需HTTP状态码范围。 支持用户宏。\n例如:200,200-{$M}，{$M}，200-400\n仅供HTTP代理监控项使用。",
		FinalType: "string",
	}

	f18 := &fetch.FieldObject{
		RawKey:    "timeout",
		NewKey:    "Timeout",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "项数据轮询请求超时。 支持用户宏。\n\nHTTP代理和Script监控项使用。",
		FinalType: "string",
	}

	f19 := &fetch.FieldObject{
		RawKey:    "verify_host",
		NewKey:    "VerifyHost",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "验证URL中的主机名是否在通用名称字段或主机证书的主题备用名称字段中。\n\n仅供HTTP代理项使用。\\t0 - NO (default)\n1 - YES",
		FinalType: "string",
	}

	f17 := &fetch.FieldObject{
		RawKey:    "verify_peer",
		NewKey:    "VerifyPeer",
		RawType:   "string",
		NewType:   "string",
		NewDesc:   "验证主机证书是否可信。\n\n仅供HTTP代理项使用。\n0 - NO (default)\n1 - YES",
		FinalType: "string",
	}

	obj5 := &fetch.StructObject{
		StructName: "QueryFieldObject",
		Fields: []*fetch.FieldObject{
			f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19,
		},
	}

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: "base",
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
			obj3,
			obj4,
			obj5,
		},
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
}

func generateHost() {
	target := "Host"
	path := "host"

	m := GetDefaultMappings()
	Common(target, path, m)
}

func generateHostGroup() {
	target := "HostGroup"
	path := "hostgroup"
	Common(target, path, GetDefaultMappings())
}

func generateHostInterface() {
	target := "HostInterface"
	path := "hostinterface"
	obj1 := CommonObject(target, path, 1)
	dType := "InterfaceDetail"
	obj1.StructName = dType

	Common(target, path, GetDefaultMappings(), obj1)

}

func generateDiscoverRole() {
	Common("DiscoverRole", "discoveryrule", GetDefaultMappings())
}

func generateItem() {
	m := GetDefaultMappings()
	target := "Item"
	path := "item"
	obj1 := CommonObject(target, path, 2)
	obj1.StructName = "PreprocessingObject"
	m["Params"] = "string"
	Common(target, path, m, obj1)
}

func generateItemPrototype() {
	Common("ItemPrototype", "itemprototype", GetDefaultMappings())
}

func generateValueMap() {

	target := "ValueMap"
	path := "valuemap"

	obj := CommonObject(target, path, 1)
	obj.StructName = "ValueMapMappingObject"

	Common(target, path, GetDefaultMappings(), obj)
}

func generateTemplate() {
	Common("Template", "template", GetDefaultMappings())
}

func generateMacro() {
	target := "Macro"
	path := "usermacro"

	fs1 := fetch.FetchDoc(common.GetDocUrl(path+"/object"), 1)
	fs := &fetch.FieldObject{
		NewKey:    "Globalmacroid",
		NewType:   "string",
		FinalType: "string",
		NewDesc:   "(readonly) 全局宏的ID.",
		Tag:       "`json:\"globalmacroid,omitempty\"`",
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
	target := "HttpTest"
	path := "httptest"
	obj1 := CommonObject(target, path, 2)
	obj1.StructName = "HttpWebStepObject"
	Common(target, path, GetDefaultMappings(), obj1)
}

func generateEvent() {
	target := "Event"
	path := "event"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	obj3 := CommonObject(target, path, 2)
	obj3.StructName = "MediaTypeUrlObject"

	packObj := &fetch.PackageObject{
		Package:  "model",
		FileName: strings.ToLower(target),
		Structs: []*fetch.StructObject{
			obj1,
			obj2,
			obj3,
		},
	}
	m := GetDefaultMappings()
	obj1.ParseStructType(m)
	obj2.ParseStructType(m)
	obj3.ParseStructType(m)

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
	Common("Graph", "graph", GetDefaultMappings())
}

func generateGraphPrototype() {
	Common("GraphPrototype", "graphprototype", GetDefaultMappings())
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
	m := GetDefaultMappings()
	target := "MediaType"
	path := "mediatype"
	obj1 := CommonObject(target, path, 2)
	obj1.StructName = "MessageTemplateObject"
	m["MessageTemplates"] = "MessageTemplateObject"
	Common(target, path, m, obj1)
}

func generateAuditLog() {
	target := "AuditLog"
	path := "auditlog"
	obj1 := CommonObject(target, path, 0)
	obj2 := CommonGet(target, path, 0)

	m := GetDefaultMappings()
	obj1.ParseStructType(m)
	obj2.ParseStructType(m)

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

func generateTrigger() {
	Common("Trigger", "trigger", GetDefaultMappings())
}

func generateTriggerPrototype() {
	Common("TriggerPrototype", "triggerprototype", GetDefaultMappings())
}

func generateProblem() {
	target := "Problem"
	path := "problem"

	m := GetDefaultMappings()
	obj1 := CommonObject(target, path, 0)
	obj1.ParseStructType(m)
	obj2 := CommonGet(target, path, 0)
	obj2.ParseStructType(m)

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

func Common(target, path string, mappings map[string]string, adds ...*fetch.StructObject) {
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

	for _, o := range packObj.Structs {
		o.ParseStructType(mappings)
	}

	template.Write(packObj, fmt.Sprintf("%s/%s.go", packObj.Package, packObj.FileName))
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

func GetDefaultMappings() map[string]string {
	m := make(map[string]string)
	m["Filter"] = "map[string][]string"
	m["Search"] = "map[string][]string"
	m["Output"] = "[]string"

	m["Groups"] = "[]*HostGroup"
	m["Interfaces"] = "[]*HostInterface"
	m["Tags"] = "[]*TagObject"
	m["Templates"] = "[]*Template"
	m["TemplatesClear"] = "[]*Template"
	m["Macros"] = "[]*Macro"
	m["SearchInventory"] = "[]string"
	// todo inventory

	m["Details"] = "InterfaceDetail"

	m["Headers"] = "HttpHeaderObject"
	m["Parameters"] = "[]*ParameterObject"
	m["QueryFields"] = "QueryFieldObject"

	m["Preprocessing"] = "[]*PreprocessingObject"

	m["Mappings"] = "[]*ValueMapMappingObject"

	m["Templates"] = "[]*Template"
	m["TemplatesClear"] = "[]*Template"

	m["Variables"] = "[]*HttpHeaderObject"

	m["Steps"] = "[]*HttpWebStepObject"

	m["Urls"] = "[]*MediaTypeUrlObject"
	m["Gitems"] = "[]*GraphItem"
	m["Dependencies"] = "[]*Trigger"

	return m
}
