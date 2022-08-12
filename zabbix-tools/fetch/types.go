package fetch

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
)

var (
	/*

		默认：false。
			boolean	返回与在search参数中给定条件不匹配的结果。
			object	仅返回与给定过滤条件完全匹配的结果。

		接受一个数组，键是属性名，值是单个值或者要匹配值的数组。

		不适用于text 字段。
			integer	限制返回记录的数量。
			query	要返回的对象属性。

		默认：extend。
			boolean	在结果数组中，用ID作为键。
			object	返回给定通配符（不区分大小写）匹配到的结果。

		接受一个数组，键是属性名，值是要搜索的字符串。如果没有给出其他选项，将会执行LIKE "%…%"搜索。

		仅适用于string和text字段。
			boolean	如果设置为true，则返回与filter or search参数中给定的任何条件匹配的结果，而不是所有条件。

		默认：false。
			boolean	如果设置为true，则可以在search参数中使用"*"作为通配符。

		默认：false。
			string/array	按照给定的属性对结果进行排序。可用于排序的属性列表，请参考特定API get方法描述。宏在排序前不会被展开。

		如果没有给出特定的值，数据会无序返回。
			string/array	排序顺序。如果传递数组，则每个值都将与sortfield参数中给定的对应属性匹配。

		可用值：
		ASC - (默认) 升序；
		DESC - 降序。

	*/
	KeyRegexp     = regexp.MustCompile(`(\w+)\s*\n*([(（].*[)）])?\s*\n*`)
	ExcludeFields = []string{
		"countOutput",
		"editable",
		"excludeSearch",
		"filter",
		"limit",
		"output",
		"preservekeys",
		"search",
		"searchByAny",
		"searchWildcardsEnabled",
		"sortorder",
		"startSearch",
	}
)

type FieldObject struct {
	RawKey   string
	RawType  string
	RawDesc  string
	NewKey   string
	NewType  string
	NewDesc  string
	Tag      string
	ParseKey string
}

func (f *FieldObject) Parse() {
	parseKey := ParseKey(f.RawKey)
	f.NewKey = UnderscoreToUpperCamelCase(parseKey)
	f.NewType = ParseType(f.RawType)
	f.NewDesc = f.RawDesc
	f.Tag = fmt.Sprintf("`json:\"%s,omitempty\"`", parseKey)
	f.ParseKey = parseKey
}

func ParseKey(rawKey string) string {
	key := rawKey
	sm := KeyRegexp.FindAllStringSubmatch(key, 1)
	if len(sm) != 1 {
		return fmt.Sprintf("键[%s]包含中文无效", rawKey)
	}
	key = KeyRegexp.ReplaceAllString(key, sm[0][1])
	return key
}

// ParseType /**
func ParseType(rawType string) string {
	if rawType == "boolean" || rawType == "布尔值" || rawType == "布尔" {
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
	} else if rawType == "timestamp" || rawType == "时间戳" {
		return "int64"
	} else if rawType == "array" || rawType == "属组" || rawType == "数组" {
		return "[]interface{} /* TODO */"
	} else if rawType == "object" || rawType == "对象" || rawType == "HTTP fields" || rawType == "Media type URLs" {
		return "interface{} /* TODO */"
	} else if rawType == "query" || rawType == "查询" {
		return "map[string][]string"
	} else if rawType == "id" {
		return "string"
	} else if strings.Contains(rawType, "array") || strings.Contains(rawType, "数组") || strings.Contains(rawType, "属组") {
		t := strings.ReplaceAll(rawType, "array", "")
		t = strings.ReplaceAll(t, "/", "")
		t = strings.ReplaceAll(t, "数组", "")
		t = strings.ReplaceAll(t, "属组", "")
		t = strings.ReplaceAll(t, "of", "")
		t = strings.ReplaceAll(t, "objects", "object")
		t = strings.TrimSpace(t)
		return fmt.Sprintf("%s%s", "[]", ParseType(t))
	} else {
		panic(errors.New(fmt.Sprintf("type %s don't support", rawType)))
	}
}

// UnderscoreToUpperCamelCase
//下划线单词转为大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = cases.Title(language.English, cases.NoLower).String(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase
// 下划线单词转为小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

type StructObject struct {
	Parent     string
	StructName string
	Fields     []*FieldObject
}

type PackageObject struct {
	Package  string
	FileName string
	Structs  []*StructObject
}
