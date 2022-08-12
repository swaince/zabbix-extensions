package model

type HostGroup struct {

	/**
	  (只读) 主机组ID。
	*/
	Groupid string `json:"groupid,omitempty"`

	/**
	  主机组的名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  (只读) 主机组的来源。取值范围:0—普通主机组;4—自动发现的主机组。
	*/
	Flags int64 `json:"flags,omitempty"`

	/**
	  (只读) 组是否被系统内部使用。内部组不能被删除。取值范围:0 - (默认值) 非内部;1 - 内部。
	*/
	Internal int64 `json:"internal,omitempty"`

	/**
	  统一唯一标识符，用于将导入的主机组与已经存在的主机组连接。自动生成，如果没有给出。对于更新操作，此字段为 只读。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type HostGroupCreateParam struct {
	HostGroup
}

type HostGroupUpdateParam struct {
	HostGroup
}

type HostGroupGetParam struct {
	GetBaseObject

	/**
	  只返回包含给定图形的主机或模板的主机组。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回具有给定主机组id的主机组。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  只返回包含给定主机的主机组。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  只返回受给定维护影响的主机组。
	*/
	Maintenanceids []string `json:"maintenanceids,omitempty"`

	/**
	  只返回包含监控主机的主机组。
	*/
	MonitoredHosts bool `json:"monitored_hosts,omitempty"`

	/**
	  只返回包含主机的主机组。
	*/
	RealHosts bool `json:"real_hosts,omitempty"`

	/**
	  只返回包含模板的主机组。
	*/
	TemplatedHosts bool `json:"templated_hosts,omitempty"`

	/**
	  只返回包含给定模板的主机组。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  只返回包含具有给定触发器的主机或模板的主机组。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  只返回包含有图形的主机的主机组。
	*/
	WithGraphs bool `json:"with_graphs,omitempty"`

	/**
	  只返回包含带有图形原型的主机的主机组。
	*/
	WithGraphPrototypes bool `json:"with_graph_prototypes,omitempty"`

	/**
	  只返回包含主机 或 模板的主机组。
	*/
	WithHostsAndTemplates bool `json:"with_hosts_and_templates,omitempty"`

	/**
	  只返回包含web检查主机的主机组。覆盖 with_monitored_httptests 参数。
	*/
	WithHttptests bool `json:"with_httptests,omitempty"`

	/**
	  只返回包含有项的主机或模板的主机组。覆盖 with_monitored_items 和 with_simple_graph_items 参数。
	*/
	WithItems bool `json:"with_items,omitempty"`

	/**
	  只返回包含带有项目原型的主机的主机组。覆盖 with_simple_graph_item_prototypes 参数。
	*/
	WithItemPrototypes bool `json:"with_item_prototypes,omitempty"`

	/**
	  只返回包含具有项目原型的主机的主机组，项目原型支持创建并具有数字类型的信息。
	*/
	WithSimpleGraphItemPrototypes bool `json:"with_simple_graph_item_prototypes,omitempty"`

	/**
	  只返回包含启用web检查的主机的主机组。
	*/
	WithMonitoredHttptests bool `json:"with_monitored_httptests,omitempty"`

	/**
	  只返回包含启用项的主机或模板的主机组。覆盖 with_simple_graph_items 参数。
	*/
	WithMonitoredItems bool `json:"with_monitored_items,omitempty"`

	/**
	  只返回包含启动触发器的主机的主机组。触发器中使用的所有项目也必须启用。
	*/
	WithMonitoredTriggers bool `json:"with_monitored_triggers,omitempty"`

	/**
	  只返回包含带有数字项的主机的主机组。
	*/
	WithSimpleGraphItems bool `json:"with_simple_graph_items,omitempty"`

	/**
	  只返回包含有触发器的主机的主机组。覆盖 with_monitored_triggers 参数。
	*/
	WithTriggers bool `json:"with_triggers,omitempty"`

	/**
	  返回 发现规则 属性和创建主机组的LLD规则。
	*/
	SelectDiscoveryRule map[string][]string `json:"selectDiscoveryRule,omitempty"`

	/**
	  返回 groupDiscovery 属性和主机组发现对象.发现主机组对象将发现的主机组链接到主机组原型，具有以下属性:groupid - (字符串) 发现的主机组ID;lastcheck - (时间戳) 最后发现主机组的时间;name - (字符串) 主机组原型的名称;parent_group_prototypeid - (字符串) 创建主机组的主机组原型ID;ts_delete - (时间戳) 当主机组未被发现时，删除该主机组的时间。
	*/
	SelectGroupDiscovery map[string][]string `json:"selectGroupDiscovery,omitempty"`

	/**
	  返回主机属性和属于主机组的主机。支持 count.
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  返回 模板 属性和属于主机组的模板。支持 count.
	*/
	SelectTemplates map[string][]string `json:"selectTemplates,omitempty"`

	/**
	  限制子选择返回的记录数量。适用于下列子选择项:selectHosts - 结果将按 host 排序;selectTemplates - r结果将按 host 排序。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  根据给定的属性对结果进行排序。可能的值是: groupid, name.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
