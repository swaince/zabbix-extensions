package model

type TriggerPrototype struct {

	/**
	  (只读) 触发器原型的ID.
	*/
	Triggerid string `json:"triggerid,omitempty"`

	/**
	  触发器原型的名称.
	*/
	Description string `json:"description,omitempty"`

	/**
	  简化的触发器表达式.
	*/
	Expression string `json:"expression,omitempty"`

	/**
	  触发器生成的事件名称.
	*/
	EventName string `json:"event_name,omitempty"`

	/**
	  操作数据.
	*/
	Opdata string `json:"opdata,omitempty"`

	/**
	  触发器原型的附加注释.
	*/
	Comments string `json:"comments,omitempty"`

	/**
	  触发器原型的严重性.有效的值:0 - (默认) 未分类;1 - 信息;2 - 警告;3 - 一般严重;4 - 严重;5 - 灾难.
	*/
	Priority int64 `json:"priority,omitempty"`

	/**
	  触发器原型是启用还是禁用.有效的值:0 - (默认) 已启用;1 - 已禁用.
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  (只读) 触发器原型父模板的ID.
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  触发器原型是否可以产生多个问题事件.有效的值:0 - (默认) 不能生成多个事件;1 - 可以生成多个事件.
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  关联到触发器原型的URL.
	*/
	Url string `json:"url,omitempty"`

	/**
	  正常事件生成模式.有效的值为:0 - (默认) 表达式;1 - 恢复表达式;2 - 无.
	*/
	RecoveryMode int64 `json:"recovery_mode,omitempty"`

	/**
	  简单的触发器恢复表达式.
	*/
	RecoveryExpression string `json:"recovery_expression,omitempty"`

	/**
	  正常事件关闭.有效的值为:0 - (默认) 所有问题;1 - 标签值匹配成功的所有问题.
	*/
	CorrelationMode int64 `json:"correlation_mode,omitempty"`

	/**
	  匹配标签.
	*/
	CorrelationTag string `json:"correlation_tag,omitempty"`

	/**
	  允许手动关闭.有效的值为:0 - (默认) 否;1 - 是.
	*/
	ManualClose int64 `json:"manual_close,omitempty"`

	/**
	  触发器原型发现状态.有效的值:0 - (默认) 将发现新的触发器;1 - 不会发现新触发器，现有触发器将被标记为丢失.
	*/
	Discover int64 `json:"discover,omitempty"`

	/**
	  通用唯一标识符，用于将导入的触发器原型链接到现有的触发器原型。仅用于模板上的触发器原型。如果未给出，则自动生成.对于更新操作，此字段为 只读.
	*/
	Uuid string `json:"uuid,omitempty"`
}

type TriggerPrototypeCreateParam struct {
	TriggerPrototype

	/**
	  触发器原型所依赖的触发器和触发器原型.触发器必须已定义triggerid属性.
	*/
	Dependencies []interface{}/* TODO */ `json:"dependencies,omitempty"`

	/**
	  触发器 标签.
	*/
	Tags []interface{}/* TODO */ `json:"tags,omitempty"`
}

type TriggerPrototypeUpdateParam struct {
	TriggerPrototype

	/**
	  触发器原型所依赖的触发器和触发器原型.触发器必须已定义triggerid属性..
	*/
	Dependencies []interface{}/* TODO */ `json:"dependencies,omitempty"`

	/**
	  触发器原型 标签.
	*/
	Tags []interface{}/* TODO */ `json:"tags,omitempty"`
}

type TriggerPrototypeGetParam struct {
	GetBaseObject

	/**
	  仅返回属于被监控主机的已启用触发器原型.
	*/
	Active bool `json:"active,omitempty"`

	/**
	  仅返回属于指定低级别发现规则的触发器原型.
	*/
	Discoveryids []string `json:"discoveryids,omitempty"`

	/**
	  仅返回使用指定函数的触发器.有关支持的功能列表，请参阅支持的触发功能 页面.
	*/
	Functions []string `json:"functions,omitempty"`

	/**
	  仅返回属于指定名称的主机组中属于主机的触发器原型.
	*/
	Group string `json:"group,omitempty"`

	/**
	  仅返回来自指定主机组ID中属于主机的触发器原型.
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回属于指名称的主机的触发器原型.
	*/
	Host string `json:"host,omitempty"`

	/**
	  仅返回属于指定主机id的触发器原型.
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  如果设置为true，仅返回从模板继承的触发器原型.
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  如果设置为true，仅返回属于在维护中主机的已启用触发器原型，.
	*/
	Maintenance bool `json:"maintenance,omitempty"`

	/**
	  仅返回严重性大于或等于指定严重性的触发器原型.
	*/
	MinSeverity int64 `json:"min_severity,omitempty"`

	/**
	  仅返回属于被监控主机且仅包含已启用监控项的已启用触发器原型.
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  如果设置为 true 仅返回属于模板的触发器原型.
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  仅返回属于指定模板ID的触发器原型.
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  仅返回属于给定ID的触发器原型.
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  展开在触发器原型表达式中的函数和宏.
	*/
	ExpandExpression bool `json:"expandExpression,omitempty"`

	/**
	  在 dependencies 属性中返回触发器原型和触发器原型所依赖的触发器.
	*/
	SelectDependencies map[string][]string `json:"selectDependencies,omitempty"`

	/**
	  返回触发器原型所属的 低级别发现规则 .
	*/
	SelectDiscoveryRule map[string][]string `json:"selectDiscoveryRule,omitempty"`

	/**
	  在 functions 属性中返回触发器原型中使用的函数.函数对象表示触发器表达式中使用的函数，并具有以下属性:functionid - (string) 函数的ID;itemid - (string) 在函数中使用的监控项ID;function - (string) 函数名称;parameter - (string) 传递给函数的参数，查询参数被返回字符串中的 $ 符号替换.
	*/
	SelectFunctions map[string][]string `json:"selectFunctions,omitempty"`

	/**
	  在 组 属性中返回触发器原型所属的主机组.
	*/
	SelectGroups map[string][]string `json:"selectGroups,omitempty"`

	/**
	  在 主机 属性中返回触发器原型所属的主机.
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  返回项目和项目原型使用 监控项 属性中的触发器原型.
	*/
	SelectItems map[string][]string `json:"selectItems,omitempty"`

	/**
	  返回 标签 中的触发器原型标签.
	*/
	SelectTags map[string][]string `json:"selectTags,omitempty"`

	/**
	  限制子查询返回的记录数量.适用于以下子查询:selectHosts - 以host分类结果.
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  根据指定属性对结果分类.有效的值为: triggerid, description, status 和 priority.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
