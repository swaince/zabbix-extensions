package model

type Host struct {

	/**
	  (只读) 主机的ID。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  主机的正式名称。
	*/
	Host string `json:"host,omitempty"`

	/**
	  主机的描述信息。
	*/
	Description string `json:"description,omitempty"`

	/**
	  (只读) 主机的来源。取值范围:0—普通主机;4—自动发现的主机。
	*/
	Flags int64 `json:"flags,omitempty"`

	/**
	  主机清单模式。取值范围为:-1 - (默认) 禁用的;0 - 手动的;1 - 自动的。
	*/
	InventoryMode int64 `json:"inventory_mode,omitempty"`

	/**
	  IPMI 验证算法。可能的值:-1 - (默认) default;0 - none;1 - MD2;2 - MD54 - straight;5 - OEM;6 - RMCP+.
	*/
	IpmiAuthtype int64 `json:"ipmi_authtype,omitempty"`

	/**
	  IPMI 密码。
	*/
	IpmiPassword string `json:"ipmi_password,omitempty"`

	/**
	  IPMI 权限等级。可能的值:1 - callback;2 - (默认) user;3 - operator;4 - admin;5 - OEM.
	*/
	IpmiPrivilege int64 `json:"ipmi_privilege,omitempty"`

	/**
	  IPMI 用户名。
	*/
	IpmiUsername string `json:"ipmi_username,omitempty"`

	/**
	  (只读) 有效维护启动时间。
	*/
	MaintenanceFrom int64 `json:"maintenance_from,omitempty"`

	/**
	  (只读) 有效的维护状态。取值包括:0 - (默认)不维护;1 -有效维护。
	*/
	MaintenanceStatus int64 `json:"maintenance_status,omitempty"`

	/**
	  (只读) 有效的维护类型。可能的值是:0 - (默认) 数据收集维护;1 - 数据不收集维护。
	*/
	MaintenanceType int64 `json:"maintenance_type,omitempty"`

	/**
	  (只读) 主机上当前生效的维护ID。
	*/
	Maintenanceid string `json:"maintenanceid,omitempty"`

	/**
	  可见的主机名。默认: host 属性值。
	*/
	Name string `json:"name,omitempty"`

	/**
	  用于监控主机的Proxy ID。
	*/
	ProxyHostid string `json:"proxy_hostid,omitempty"`

	/**
	  主机的状态和功能。可能的值是:0 - (默认) 开启监控的主机;1 - 没有开启监控的主机。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  链接到主机。可能的值是:1 - (默认) 没有加密;2 - PSK;4 - 已加密。
	*/
	TlsConnect int64 `json:"tls_connect,omitempty"`

	/**
	  连接主机。可能的位图值如下::1 - (默认) 未加密;2 - PSK;4 - 已认证.
	*/
	TlsAccept int64 `json:"tls_accept,omitempty"`

	/**
	  证书发行者。
	*/
	TlsIssuer string `json:"tls_issuer,omitempty"`

	/**
	  证书问题。
	*/
	TlsSubject string `json:"tls_subject,omitempty"`

	/**
	  (只写) PSK 认证。 如果启用了 tls_connect 或 tls_accept ，则需要。不要将敏感信息放入PSK标识中，该信息通过网络不加密传输，以告知接收端使用哪个PSK。
	*/
	TlsPskIdentity string `json:"tls_psk_identity,omitempty"`

	/**
	  (只写) 预共享密钥，至少32位十六进制数字。如果启用了 tls_connect 或 tls_accept ，则需要。
	*/
	TlsPsk string `json:"tls_psk,omitempty"`
}

type HostCreateParam struct {
	Host

	/**
	  把主机添加到目标组。 主机组必须已定义 groupid 属性。
	*/
	Groups []*HostGroup `json:"groups,omitempty"`

	/**
	  为主机创建的接口。
	*/
	Interfaces []*HostInterface `json:"interfaces,omitempty"`

	/**
	  主机标签。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  主机连接的模板。模板必须已定义templateid属性。
	*/
	Templates []*Template `json:"templates,omitempty"`

	/**
	  为主机创建的用户宏。
	*/
	Macros []*Macro `json:"macros,omitempty"`

	/**
	  主机资产清单属性。
	*/
	Inventory interface{}/* TODO */ `json:"inventory,omitempty"`
}

type HostUpdateParam struct {
	Host

	/**
	  更换主机所属的主机组。主机组必须已定义 groupid 属性。所有未在请求中列出的主机组将被解除链接。
	*/
	Groups []*HostGroup `json:"groups,omitempty"`

	/**
	  主机接口用于替换当前主机接口。 所有未在请求中列出的接口将被删除。
	*/
	Interfaces []*HostInterface `json:"interfaces,omitempty"`

	/**
	  主机标签替换当前主机标签。所有未在请求中列出的标签将被删除。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  主机清单属性。
	*/
	Inventory interface{}/* TODO */ `json:"inventory,omitempty"`

	/**
	  用户宏替换当前用户宏。所有未在请求中列出的宏将被删除。
	*/
	Macros []*Macro `json:"macros,omitempty"`

	/**
	  模板替换当前链接的模板。所有未在请求中列出的模板只会被解除链接。模板必须定义了 templateid 属性。
	*/
	Templates []*Template `json:"templates,omitempty"`

	/**
	  模板取消与主机的链接并清除。模板必须定义了 templateid 属性。
	*/
	TemplatesClear []*Template `json:"templates_clear,omitempty"`
}

type HostGetParam struct {
	GetBaseObject

	/**
	  只返回属于给定组的主机。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  只返回与所发现的给定服务相关的主机。
	*/
	Dserviceids []string `json:"dserviceids,omitempty"`

	/**
	  只返回具有给定图形的主机。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回具有给定主机id的主机。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  只返回具有给定web检查的主机。
	*/
	Httptestids []string `json:"httptestids,omitempty"`

	/**
	  只返回使用给定接口的主机。
	*/
	Interfaceids []string `json:"interfaceids,omitempty"`

	/**
	  只返回具有给定监控项的主机。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  只返回受给定维护影响的主机。
	*/
	Maintenanceids []string `json:"maintenanceids,omitempty"`

	/**
	  只返回开启监控的主机。
	*/
	MonitoredHosts bool `json:"monitored_hosts,omitempty"`

	/**
	  只返回代理的主机。
	*/
	ProxyHosts bool `json:"proxy_hosts,omitempty"`

	/**
	  只返回被给定代理监视的主机。
	*/
	Proxyids []string `json:"proxyids,omitempty"`

	/**
	  返回主机和模板。
	*/
	TemplatedHosts bool `json:"templated_hosts,omitempty"`

	/**
	  只返回链接到给定模板的主机。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  只返回具有给定触发器的主机。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  只返回有监控项的主机。覆盖 with_monitored_items 和 with_simple_graph_items 参数。
	*/
	WithItems bool `json:"with_items,omitempty"`

	/**
	  只返回具有项目原型的主机。覆盖 with_simple_graph_item_prototypes 参数。
	*/
	WithItemPrototypes bool `json:"with_item_prototypes,omitempty"`

	/**
	  只返回具有项目原型的主机，该原型支持创建并具有数字类型的信息。
	*/
	WithSimpleGraphItemPrototypes bool `json:"with_simple_graph_item_prototypes,omitempty"`

	/**
	  只返回有图形的主机。
	*/
	WithGraphs bool `json:"with_graphs,omitempty"`

	/**
	  只返回具有图形原型的主机。
	*/
	WithGraphPrototypes bool `json:"with_graph_prototypes,omitempty"`

	/**
	  只返回有网络检查的主机。覆盖 with_monitored_httptests 参数.
	*/
	WithHttptests bool `json:"with_httptests,omitempty"`

	/**
	  只返回启用web检查的主机。
	*/
	WithMonitoredHttptests bool `json:"with_monitored_httptests,omitempty"`

	/**
	  只返回启用监控项的主机。覆盖 with_simple_graph_items 参数。
	*/
	WithMonitoredItems bool `json:"with_monitored_items,omitempty"`

	/**
	  只返回启用触发器的主机。触发器中使用的所有监控项也必须启用。
	*/
	WithMonitoredTriggers bool `json:"with_monitored_triggers,omitempty"`

	/**
	  只返回具有数字类型信息的条目的主机。
	*/
	WithSimpleGraphItems bool `json:"with_simple_graph_items,omitempty"`

	/**
	  只返回有触发器的主机。覆盖 with_monitored_triggers 参数。
	*/
	WithTriggers bool `json:"with_triggers,omitempty"`

	/**
	  返回已抑制问题的主机。可能的值:null - (默认) 所有主机;true - 仅包含被抑制问题的主机;false - 只允许存在未抑制问题的主机。
	*/
	WithProblemsSuppressed bool `json:"withProblemsSuppressed,omitempty"`

	/**
	  标记搜索规则。可能的值:0 - (默认) 和/或;2 - 或.
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  返回只存在给定严重性问题的主机。仅当问题对象为触发器时适用。
	*/
	Severities []int64 `json:"severities,omitempty"`

	/**
	  只返回带有给定标签的主机。根据标记精确匹配，根据标记值进行大小写敏感或不区分大小写的搜索，具体取决于操作符值。格式: [{"tag": "<tag>", "value": "<value>", "operator": "<operator>"}, ...].空数组返回所有主机。可能的操作符值:0 - (默认) 包含;1 - 登录;2 - 不类似;3 - 不等于4 - 存在;5 - 不存在。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  返回那些在所有链接模板中也给出了 tags 的主机。 默认:可能的值:true - 链接模板也必须有tags;false - (默认) 链接的模板标签将被忽略。
	*/
	InheritedTags bool `json:"inheritedTags,omitempty"`

	/**
	  返回 发现 属性和主机低级发现规则。支持 count.
	*/
	SelectDiscoveries []string `json:"selectDiscoveries,omitempty"`

	/**
	  返回 发现规则 属性和创建主机的底层发现规则(来自VMware监控中的主机原型)。
	*/
	SelectDiscoveryRule []string `json:"selectDiscoveryRule,omitempty"`

	/**
	  返回 图形 属性和主机图。支持 count.
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  返回 主机组 属性和主机所属的主机组数据。
	*/
	SelectGroups []string `json:"selectGroups,omitempty"`

	/**
	  返回 hostDiscovery 属性和主机发现对象数据。主机发现对象将发现的主机链接到主机原型，或将主机原型链接到LLD规则，并具有以下属性:host - (string) 主机原型的主机;hostid - (string) 发现的主机或主机原型ID;parent_hostid - (string) 创建主机的主机原型ID;parent_itemid - (string) 发现主机的LLD规则ID;lastcheck - (timestamp) 最后发现主机的时间;ts_delete - (timestamp) 删除已发现主机的时间。
	*/
	SelectHostDiscovery []string `json:"selectHostDiscovery,omitempty"`

	/**
	  返回 httpTests 属性与主机web场景.支持 count.
	*/
	SelectHttpTests []string `json:"selectHttpTests,omitempty"`

	/**
	  返回 interfaces 属性和主机接口Supports count。
	*/
	SelectInterfaces []string `json:"selectInterfaces,omitempty"`

	/**
	  返回 inventory 属性和主机库存数据。
	*/
	SelectInventory []string `json:"selectInventory,omitempty"`

	/**
	  返回 items 属性和主机项。支持 count.
	*/
	SelectItems []string `json:"selectItems,omitempty"`

	/**
	  返回 macros 属性和主机宏。
	*/
	SelectMacros []string `json:"selectMacros,omitempty"`

	/**
	  返回 parentTemplates 属性和主机链接到的模板。支持 count.
	*/
	SelectParentTemplates []string `json:"selectParentTemplates,omitempty"`

	/**
	  返回 dashboards 属性.支持 count.
	*/
	SelectDashboards []string `json:"selectDashboards,omitempty"`

	/**
	  返回 tags 属性和主机标记。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  返回 inheritedTags 属性的标记，这些标记位于链接到主机的所有模板上。
	*/
	SelectInheritedTags []string `json:"selectInheritedTags,omitempty"`

	/**
	  返回 triggers 属性和主机触发器。支持 count.
	*/
	SelectTriggers []string `json:"selectTriggers,omitempty"`

	/**
	  返回 valuemaps 属性和主机映射值。
	*/
	SelectValueMaps []string `json:"selectValueMaps,omitempty"`

	/**
	  限制子选择返回的记录数量。适用于以下子选择:selectParentTemplates - 结果将按 host 排序;selectInterfaces;selectItems - 结果将按 name 排序;selectDiscoveries - 结果将按 name 排序;selectTriggers - 结果将按 description 排序;selectGraphs - 结果将按 name 排序;selectDashboards - 结果将按 name 排序。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  只返回库存数据与给定通配符搜索匹配的主机。该参数受与 search 相同的附加参数的影响。
	*/
	SearchInventory []string `json:"searchInventory,omitempty"`

	/**
	  根据给定的属性对结果进行排序。可能的值: hostid, host, name, status.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
