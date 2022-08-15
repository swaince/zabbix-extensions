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
	Filter map[string][]string `json:"filter,omitempty"`

	/**
	  限制返回记录的数量。
	*/
	Limit int64 `json:"limit,omitempty"`

	/**
	  要返回的对象属性。默认：extend。
	*/
	Output []string `json:"output,omitempty"`

	/**
	  在结果数组中，用ID作为键。
	*/
	Preservekeys bool `json:"preservekeys,omitempty"`

	/**
	  返回给定通配符（不区分大小写）匹配到的结果。接受一个数组，键是属性名，值是要搜索的字符串。如果没有给出其他选项，将会执行LIKE "%…%"搜索。仅适用于string和text字段。
	*/
	Search map[string][]string `json:"search,omitempty"`

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

type HttpHeaderObject struct {

	/**
	  Header 名称.
	*/
	Name string

	/**
	  Header 值.
	*/
	Value string

	/**
	    HTTP代理连接字符串。
	仅供HTTP代理项使用。
	*/
	HttpProxy string

	/**
	    	0 - RAW (default)
	1 - JSON
	H如何处理反应。
	仅供HTTP代理项使用。
	*/
	OutputFormat string

	/**
	    0 - RAW (default)
	2 - JSON
	3 - XML
	发布数据体的类型。
	仅供HTTP代理项使用。
	*/
	PostType string

	/**
	    HTTP(S)请求体数据。
	仅供HTTP代理项使用
	*/
	Posts string
}

type ParameterObject struct {

	/**
	    参数名称。
	仅供脚本项使用。
	*/
	Name string

	/**
	    参数值。
	仅供脚本项使用。
	*/
	Value string
}

type QueryFieldObject struct {

	/**
	  参数名称
	*/
	Name string

	/**
	  参数值
	*/
	Value string

	/**
	    请求方法。

	仅供HTTP代理项使用。
	0 - GET (default)
	1 - POST
	2 - PUT
	3 - HEAD
	*/
	RequestMethod string

	/**
	    应该存储响应的哪一部分。

	仅供HTTP代理项使用。
	0 - BODY (default)
	1 - HEADERS
	2 - BOTH
	*/
	RetrieveMode string

	/**
	    Public SSL密钥文件路径。
	仅供HTTP代理项使用。
	*/
	SslCertFile string

	/**
	    Private SSL Key文件路径。
	仅供HTTP代理项使用。
	*/
	SslKeyFile string

	/**
	    SSL密钥文件的密码。
	仅供HTTP代理项使用。
	*/
	SslKeyPassword string

	/**
	    用逗号分隔的所需HTTP状态码范围。 支持用户宏。
	例如:200,200-{$M}，{$M}，200-400
	仅供HTTP代理监控项使用。
	*/
	StatusCodes string

	/**
	    验证主机证书是否可信。

	仅供HTTP代理项使用。
	0 - NO (default)
	1 - YES
	*/
	VerifyPeer string

	/**
	    项数据轮询请求超时。 支持用户宏。

	HTTP代理和Script监控项使用。
	*/
	Timeout string

	/**
	    验证URL中的主机名是否在通用名称字段或主机证书的主题备用名称字段中。

	仅供HTTP代理项使用。\t0 - NO (default)
	1 - YES
	*/
	VerifyHost string
}
