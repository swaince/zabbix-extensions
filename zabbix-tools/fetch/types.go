package fetch

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	KeyRegexp = regexp.MustCompile(`(\w+)\s*\n*([(（].*[)）])?\s*\n*`)
)

type FieldObject struct {
	RawKey  string
	RawType string
	RawDesc string
	NewKey  string
	NewType string
	NewDesc string
}

func (f *FieldObject) Parse() {
	f.NewKey = ParseKey(f.RawKey)
	f.NewType = ParseType(f.RawType)
	f.NewDesc = f.RawDesc
}

func ParseKey(rawKey string) string {
	key := rawKey
	sm := KeyRegexp.FindAllStringSubmatch(key, 1)
	key = KeyRegexp.ReplaceAllString(key, sm[0][1])
	return key
}

// ParseType /**
func ParseType(rawType string) string {
	if strings.Contains(rawType, "array") || strings.Contains(rawType, "数组") || strings.Contains(rawType, "属组") {
		t := strings.ReplaceAll(rawType, "array", "")
		t = strings.ReplaceAll(t, "/", "")
		t = strings.ReplaceAll(t, "数组", "")
		t = strings.ReplaceAll(t, "属组", "")
		return fmt.Sprintf("%s%s", "[]", ParseType(t))
	} else if rawType == "boolean" || rawType == "布尔值" {
		return "bool"
	} else if rawType == "flag" || rawType == "标记" {
		return "bool"
	} else if rawType == "integer" || rawType == "整数" || rawType == "整型" {
		return "int64"
	} else if rawType == "float" || rawType == "浮点数" || rawType == "浮点型" {
		return "float64"
	} else if rawType == "string" || rawType == "字符串" {
		return "string"
	} else if rawType == "text" {
		return "string"
	} else if rawType == "timestamp" {
		return "int64"
	} else if rawType == "array" || rawType == "属组" || rawType == "数组" {
		return "[]interface{}"
	} else if rawType == "object" || rawType == "对象" {
		return "interface{}"
	} else if rawType == "query" || rawType == "查询" {
		return "map[string][]string"
	} else {
		panic(errors.New("type %s don't support"))
	}
}

type ClassObject struct {
	Parent     string
	Package    string
	StructName string
	Fields     []*FieldObject
}
