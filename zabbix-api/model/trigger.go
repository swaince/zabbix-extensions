package model

type Trigger struct {

	/**
	  (只读) 触发器的ID.
	*/
	Triggerid string `json:"triggerid,omitempty"`

	/**
	  触发器的名称.
	*/
	Description string `json:"description,omitempty"`

	/**
	  简化的触发器表达式.
	*/
	Expression string `json:"expression,omitempty"`

	/**
	  生成触发器的事件名称.
	*/
	EventName string `json:"event_name,omitempty"`

	/**
	  当前的数据.
	*/
	Opdata string `json:"opdata,omitempty"`

	/**
	  触发器的附加说明.
	*/
	Comments string `json:"comments,omitempty"`

	/**
	  (只读) 更新触发器状态时出现的任何问题的错误文本.
	*/
	Error string `json:"error,omitempty"`

	/**
	  (只读) 原始触发器.有效的值为:0 - (默认) 普通触发器;4 - 自动发现的触发器.
	*/
	Flags int64 `json:"flags,omitempty"`

	/**
	  (只读) 触发器最后更改其状态的时间.
	*/
	Lastchange int64 `json:"lastchange,omitempty"`

	/**
	  触发器的严重性级别.有效的值为:0 - (默认) 未分类;1 - 信息;2 - 警告;3 - 一般严重;4 - 严重;5 - 灾难.
	*/
	Priority int64 `json:"priority,omitempty"`

	/**
	  (只读) 触发器的状态.有效的值为:0 - (默认) 触发器状态是最新的;1 - 当前的触发器状态是未知的.
	*/
	State int64 `json:"state,omitempty"`

	/**
	  触发器是否处于启用状态或禁用状态.有效的值为:0 - (默认) 已启用;1 - 已禁用.
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  (只读) 父触发器模板ID.
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  触发器是否能够生成多个故障事件.有效的值为:0 - (默认) 不能生成多个事件;1 - 可以生成多个事件.
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  与触发器相关联的URL.
	*/
	Url string `json:"url,omitempty"`

	/**
	  (只读) 触发器是否处于正常或故障状态.有效的值为:0 - (默认) 正常状态;1 - 故障状态.
	*/
	Value int64 `json:"value,omitempty"`

	/**
	  事件恢复生成模式.有效的值为:0 - (默认) 表达式;1 - 恢复表达式;2 - 无.
	*/
	RecoveryMode int64 `json:"recovery_mode,omitempty"`

	/**
	  生成的触发恢复表达式.
	*/
	RecoveryExpression string `json:"recovery_expression,omitempty"`

	/**
	  事件恢复关联的模式.有效的值为:0 - (默认) 所有故障;1 - 与标签值匹配的所有故障.
	*/
	CorrelationMode int64 `json:"correlation_mode,omitempty"`

	/**
	  用于匹配的标签.
	*/
	CorrelationTag string `json:"correlation_tag,omitempty"`

	/**
	  允许手动关闭.有效的值为:0 - (默认) 不允许;1 - 允许.
	*/
	ManualClose int64 `json:"manual_close,omitempty"`

	/**
	  通用唯一标识符，用于将导入的触发器链接到已经存在的触发器，仅用于模板上的触发器，如果没有提前给出则标识符会自动生成.此字段是 只读的，不能通过更新操作修改已有标识符.
	*/
	Uuid string `json:"uuid,omitempty"`
}

type TriggerCreateParam struct {
	Trigger

	/**
	  目的触发器.目的触发器必须存在且已定义triggerid属性.
	*/
	Dependencies []*Trigger `json:"dependencies,omitempty"`

	/**
	  触发器标签 标签.
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type TriggerUpdateParam struct {
	Trigger

	/**
	  依赖触发的触发器.触发器必须已定义triggerid属性.
	*/
	Dependencies []*Trigger `json:"dependencies,omitempty"`

	/**
	  触发器 标签.
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type TriggerGetParam struct {
	GetBaseObject

	/**
	  仅返回属于指定ID的触发器.
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  仅返回属于指定主机组ID的触发器.
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回属于指定模板ID的触发器.
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  仅返回属于指定主机ID的触发器.
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回包含指定监控项ID的触发器.
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  仅返回使用指定函数的触发器.请参阅支持的函数 页面查看有关支持的功能列表.
	*/
	Functions []string `json:"functions,omitempty"`

	/**
	  仅返回属于指定名称的主机组的触发器.
	*/
	Group string `json:"group,omitempty"`

	/**
	  仅返回属于指定名称的主机的触发器.
	*/
	Host string `json:"host,omitempty"`

	/**
	  如果设置为 true 则仅返回从模板继承的触发器.
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  如果设置为 true 仅返回属于模板的触发器.
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  如果设置为 true，则仅返回具有依赖关系的触发器，如果设置为 false 只返回没有依赖关系的触发器.
	*/
	Dependent bool `json:"dependent,omitempty"`

	/**
	  仅返回已启用的触发器且属于已被监控的主机，并且仅包含已启用的监控项.
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  仅返回属于已被监控主机的已启用触发器.
	*/
	Active bool `json:"active,omitempty"`

	/**
	  如果设置为 true，则仅返回属于维护中主机的启用触发器.
	*/
	Maintenance bool `json:"maintenance,omitempty"`

	/**
	  仅返回具有未确认事件的触发器.
	*/
	WithUnacknowledgedEvents bool `json:"withUnacknowledgedEvents,omitempty"`

	/**
	  仅返回已被确认的所有事件的触发器.
	*/
	WithAcknowledgedEvents bool `json:"withAcknowledgedEvents,omitempty"`

	/**
	  仅返回最后一个未被确认事件的触发器.
	*/
	WithLastEventUnacknowledged bool `json:"withLastEventUnacknowledged,omitempty"`

	/**
	  跳过处于问题状态且依赖于其他触发器的触发器，请注意，如果禁用触发器、禁用监控项或禁用主机监控项，则其他触发器将会被忽略.
	*/
	SkipDependent bool `json:"skipDependent,omitempty"`

	/**
	  仅返回在指定时间后更改了状态的触发器.
	*/
	LastChangeSince int64 `json:"lastChangeSince,omitempty"`

	/**
	  仅返回在给指时间之前更改了状态的触发器.
	*/
	LastChangeTill int64 `json:"lastChangeTill,omitempty"`

	/**
	  仅返回最近处于问题状态的触发器.
	*/
	OnlyTrue bool `json:"only_true,omitempty"`

	/**
	  仅返回严重级别大于或等于指定严重级别的触发器.
	*/
	MinSeverity int64 `json:"min_severity,omitempty"`

	/**
	  标签搜索规则.有效的值为:0 - (默认) And/Or;2 - Or.
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  只返回给定标签的触发器。按标签精确匹配，根据运算符值按标签值进行区分大小写或不区分大小写的搜索.格式: [{"tag": "<tag>", "value": "<value>", "operator": "<operator>"}, ...].空数组返回所有触发器.可能的运算符类型:0 - (默认) Like;1 - Equal;2 - Not like;3 - Not equal4 - Exists;5 - Not exists.
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  在触发器描述中展开宏.
	*/
	ExpandComment bool `json:"expandComment,omitempty"`

	/**
	  以触发器的名称展开宏.
	*/
	ExpandDescription bool `json:"expandDescription,omitempty"`

	/**
	  在触发器表达式中展开函数和宏.
	*/
	ExpandExpression bool `json:"expandExpression,omitempty"`

	/**
	  在组 属性中返回触发器所属的主机组.
	*/
	SelectGroups []string `json:"selectGroups,omitempty"`

	/**
	  返回触发器所属的 主机.
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回触发器包含的 监控项.
	*/
	SelectItems []string `json:"selectItems,omitempty"`

	/**
	  在 functions 属性中返回触发器中使用的函数.函数对象表示触发器表达式中使用的函数，并具有以下属性:functionid - (string) 函数的ID;itemid - (string) 函数中使用的监控项 ID;function - (string) 函数的名称;parameter - (string) 传递给函数的参数。查询参数被返回字符串中的$符号替换.
	*/
	SelectFunctions []string `json:"selectFunctions,omitempty"`

	/**
	  返回在dependencies 属性中依赖触发的触发器.
	*/
	SelectDependencies []string `json:"selectDependencies,omitempty"`

	/**
	  返回创建触发器的 低级发现规则.
	*/
	SelectDiscoveryRule []string `json:"selectDiscoveryRule,omitempty"`

	/**
	  在 最新事件 属性中返回最后一个重要的触发事件.
	*/
	SelectLastEvent []string `json:"selectLastEvent,omitempty"`

	/**
	  在标签 属性中返回触发标签.
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  在 triggerDiscovery 属性中返回触发器发现对象. 触发器发现对象将触发器链接到创建它的触发器原型.它具有以下属性:parent_triggerid - (string) 创建触发器的触发器原型的 ID.
	*/
	SelectTriggerDiscovery []string `json:"selectTriggerDiscovery,omitempty"`

	/**
	  限制子查询返回的记录数量.适用于以下子查询:selectHosts - 结果将按host排序.
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  按指定属性对结果进行排序 .有效的值为: triggerid, description, status, priority, lastchange 和 hostname.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
