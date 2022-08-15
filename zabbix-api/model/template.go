package model

type Template struct {

	/**
	  (只读) 模版的ID。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  - 模板的技术名称。
	*/
	Host string `json:"host,omitempty"`

	/**
	  模版的描述。
	*/
	Description string `json:"description,omitempty"`

	/**
	  模版的可见名称。默认值： host 属性值。
	*/
	Name string `json:"name,omitempty"`

	/**
	  通用唯一标识符，用于将导入的模板链接到现有模板。如果没有给出则自动生成。对于更新操作，此字段为只读。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type TemplateCreateParam struct {
	Template

	/**
	  将模版添加到主机 群组。主机群组必须定义groupid 属性。
	*/
	Groups []*HostGroup `json:"groups,omitempty"`

	/**
	  模版 标签.
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  模版 要链接到模版。模板必须定义templateid属性。
	*/
	Templates []*Template `json:"templates,omitempty"`

	/**
	  要为模版创建的用户宏 。
	*/
	Macros []*Macro `json:"macros,omitempty"`
}

type TemplateUpdateParam struct {
	Template

	/**
	  用于替换模板所属的当前主机组的主机群组。主机组必须定义groupid属性。
	*/
	Groups []*HostGroup `json:"groups,omitempty"`

	/**
	  替换当前模板标记的模板标签。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  用户宏替换给定模板上的当前用户宏。
	*/
	Macros []*Macro `json:"macros,omitempty"`

	/**
	  替换当前链接的模板的模版。未传递的模板仅被取消链接。模板必须定义templateid属性。
	*/
	Templates []*Template `json:"templates,omitempty"`

	/**
	  取消链接并清除给定模板的模版。模板必须定义templateid属性。
	*/
	TemplatesClear []*Template `json:"templates_clear,omitempty"`
}

type TemplateGetParam struct {
	GetBaseObject

	/**
	  仅返回具有给定模板ID的模板。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  仅返回属于给定主机群组的模板。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回作为给定模板子级的模板。
	*/
	ParentTemplateids []string `json:"parentTemplateids,omitempty"`

	/**
	  仅返回链接到给定主机/模板的模板。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回包含给定图形的模板。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  仅返回包含给定监控项的模板。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  仅返回包含给定触发器的模板。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  仅返回包含监控项的模板。
	*/
	WithItems bool `json:"with_items,omitempty"`

	/**
	  仅返回具有触发器的模板。
	*/
	WithTriggers bool `json:"with_triggers,omitempty"`

	/**
	  仅返回包含图形的模板。
	*/
	WithGraphs bool `json:"with_graphs,omitempty"`

	/**
	  仅返回有web场景的模板。
	*/
	WithHttptests bool `json:"with_httptests,omitempty"`

	/**
	  标签搜索规则。可能的值：0 - (默认值) And/Or;2 - Or.
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  仅返回带有给定标签的模板。根据标签进行精确匹配，并根据运算符值按标签值进行区分大小写或不区分大小写的搜索。Format: [{"tag": "<tag>", "value": "<value>", "operator": "<operator>"}, ...].一个空数组返回所有模板。可能的运算符值0 - (默认值) Contains;1 - Equals;2 - Not like;3 - Not equal4 - Exists;5 - Not exists.
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  在群组属性中返回模板所属的主机组。
	*/
	SelectGroups []string `json:"selectGroups,omitempty"`

	/**
	  在标签中返回模板标签。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  在主机 中返回链接到模板的主机。支持 count.
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回 模版 中的子模板。支持 count.
	*/
	SelectTemplates []string `json:"selectTemplates,omitempty"`

	/**
	  返回parentTemplates中的父模板支持 count.
	*/
	SelectParentTemplates []string `json:"selectParentTemplates,omitempty"`

	/**
	  从httpTests 中的模板返回web场景。支持 count.
	*/
	SelectHttpTests []string `json:"selectHttpTests,omitempty"`

	/**
	  从监控项中的模板返回监控项。支持 count.
	*/
	SelectItems []string `json:"selectItems,omitempty"`

	/**
	  从discoveries属性中的模板返回低级别自动发现。支持 count.
	*/
	SelectDiscoveries []string `json:"selectDiscoveries,omitempty"`

	/**
	  从触发器中的模板返回触发器。支持 count.
	*/
	SelectTriggers []string `json:"selectTriggers,omitempty"`

	/**
	  从图形中的模板返回图形。支持 count.
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  从macros属性中的模板返回宏..
	*/
	SelectMacros []string `json:"selectMacros,omitempty"`

	/**
	  从仪表盘属性中的模板返回仪表板。支持 count.
	*/
	SelectDashboards []string `json:"selectDashboards,omitempty"`

	/**
	  返回一个带有模板值映射的值映射属性。
	*/
	SelectValueMaps []string `json:"selectValueMaps,omitempty"`

	/**
	  限制子选择返回的记录数。适用于以下子选项：selectTemplates - 结果将按name排序selectHosts - 按host排序;selectParentTemplates - 按host排序;selectItems - 按name排序;selectDiscoveries - 按name排序;selectTriggers - 按description排序;selectGraphs - 按name排序;selectDashboards - 按name排序.
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  按给定属性对结果进行排序。可能的值： hostid, host, name, status.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
