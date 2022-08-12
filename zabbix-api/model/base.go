package model

type GetBaseObject struct {

	/**
	  返回结果中的记录数，而不是实际的数据。
	*/
	CountOutput bool `json:"countOutput,omitempty"`

	/**
	  如果设置为true，则只返回用户具有写权限的对象。默认：false。
	*/
	Editable bool `json:"editable,omitempty"`

	/**
	  返回与在search参数中给定条件不匹配的结果。
	*/
	ExcludeSearch bool `json:"excludeSearch,omitempty"`

	/**
	  仅返回与给定过滤条件完全匹配的结果。接受一个数组，键是属性名，值是单个值或者要匹配值的数组。不适用于text 字段。
	*/
	Filter interface{}/* TODO */ `json:"filter,omitempty"`

	/**
	  限制返回记录的数量。
	*/
	Limit int64 `json:"limit,omitempty"`

	/**
	  要返回的对象属性。默认：extend。
	*/
	Output map[string][]string `json:"output,omitempty"`

	/**
	  在结果数组中，用ID作为键。
	*/
	Preservekeys bool `json:"preservekeys,omitempty"`

	/**
	  返回给定通配符（不区分大小写）匹配到的结果。接受一个数组，键是属性名，值是要搜索的字符串。如果没有给出其他选项，将会执行LIKE "%…%"搜索。仅适用于string和text字段。
	*/
	Search interface{}/* TODO */ `json:"search,omitempty"`

	/**
	  如果设置为true，则返回与filter or search参数中给定的任何条件匹配的结果，而不是所有条件。默认：false。
	*/
	SearchByAny bool `json:"searchByAny,omitempty"`

	/**
	  如果设置为true，则可以在search参数中使用"*"作为通配符。默认：false。
	*/
	SearchWildcardsEnabled bool `json:"searchWildcardsEnabled,omitempty"`

	/**
	  按照给定的属性对结果进行排序。可用于排序的属性列表，请参考特定API get方法描述。宏在排序前不会被展开。如果没有给出特定的值，数据会无序返回。
	*/
	Sortfield []string `json:"sortfield,omitempty"`

	/**
	  排序顺序。如果传递数组，则每个值都将与sortfield参数中给定的对应属性匹配。可用值：ASC - (默认) 升序；DESC - 降序。
	*/
	Sortorder []string `json:"sortorder,omitempty"`

	/**
	  search参数比较字段的开始，即执行LIKE "…%"搜索。如果searchWildcardsEnabled设置为true，则忽略。
	*/
	StartSearch bool `json:"startSearch,omitempty"`
}

type TagObject struct {

	/**
	  主机标记名称。
	*/
	Tag string `json:"tag,omitempty"`

	/**
	  主机标记的值。
	*/
	Value string `json:"value,omitempty"`
}
