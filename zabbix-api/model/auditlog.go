package model

type AuditLog struct {

	/**
	  (只读) 审计日志项的 ID。使用 CUID 算法生成。
	*/
	Auditid string `json:"auditid,omitempty"`

	/**
	  审计日志项的用户 ID。
	*/
	Userid string `json:"userid,omitempty"`

	/**
	  审计日志项的用户名。
	*/
	Username string `json:"username,omitempty"`

	/**
	  审计日志项的创建时间戳。
	*/
	Clock int64 `json:"clock,omitempty"`

	/**
	  审计日志项的使用者IP地址。
	*/
	Ip string `json:"ip,omitempty"`

	/**
	  审计日志项的动作。可用值：0 - Add；1 - Update；2 - Delete；4 - Logout；7 - Execute；8 - Login；9 - Failed login；10 - History clear。
	*/
	Action int64 `json:"action,omitempty"`

	/**
	  审计日志项的资源类型。可用值：0 - 用户；3 - 媒介类型；4 - 主机；5 - 动作；6 - 图表；11 - 用户组；13 - 触发器；14 - 主机组；15 - 监控项；16 - Image；17 - Value map；18 - Service；19 - Map；22 - Web 场景；23 - 发现规则；25 - Script；26 - Proxy；27 - Maintenance；28 - 正则表达式；29 - Macro；30 - 模板；31 - Trigger prototype；32 - Icon mapping；33 - 仪表盘；34 - Event correlation；35 - Graph prototype；36 - Item prototype；37 - Host prototype；38 - Autoregistration；39 - Module；40 - Settings；41 - Housekeeping；42 - Authentication；43 - Template dashboard；44 - User role；45 - Auth token；46 - Scheduled report。
	*/
	Resourcetype int64 `json:"resourcetype,omitempty"`

	/**
	  审计日志项资源标识符。
	*/
	Resourceid string `json:"resourceid,omitempty"`

	/**
	  审计日志项可读名称。
	*/
	Resourcename string `json:"resourcename,omitempty"`

	/**
	  审计日志项记录集的 ID。在同一操作期间创建的审计日志记录将具有相同的 记录集 ID。使用 CUID 算法生成。
	*/
	Recordsetid string `json:"recordsetid,omitempty"`

	/**
	  审计日志项详情。详细信息存储为 JSON 对象，其中每个属性名称都是发生更改的属性或嵌套对象的路径，并且每个值都包含有关此属性更改的数据。格式为数组格式。可用值格式：["add"] - 被添加的嵌套对象；["add", "<value>"] - 被添加的对象属性中包含了 <value>；["update"] - 被更新的嵌套对象；["update", "<new value>", "<old value>"] - 被更新对象的属性值从 <old value> 更改为 <new value>；["delete"] - 被删除的嵌套对象。
	*/
	Details string `json:"details,omitempty"`
}

type AuditLogGetParam struct {
	GetBaseObject

	/**
	  仅返回具有给定 ID 的 audit log。
	*/
	Auditids []string `json:"auditids,omitempty"`

	/**
	  仅返回由给定用户创建的 audit log。
	*/
	Userids []string `json:"userids,omitempty"`

	/**
	  仅返回在给定时间之后或在给定时间创建的 audit log。
	*/
	TimeFrom int64 `json:"time_from,omitempty"`

	/**
	  仅返回在给定时间之前或指定时间创建的 audit log。
	*/
	TimeTill int64 `json:"time_till,omitempty"`

	/**
	  根据给定的属性排序结果。可能的值：auditid， userid， clock。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
