package model

type Event struct {

	/**
	  事件 ID。
	*/
	Eventid string `json:"eventid,omitempty"`

	/**
	  事件类型。可用值：0 - 由触发器创建的事件；1 - 由发现规则创建的事件；2 - 主动式 agent 自动注册创建的事件；3 - 内部事件；4 - 在服务状态更新时创建的事件。
	*/
	Source int64 `json:"source,omitempty"`

	/**
	  与事件相关的对象类型。触发事件的可用值：0 - 触发。发现事件的可用值：1 - 发现的主机；2 - 发现服务。自动注册事件的可用值：3 - 自动注册的主机。内部事件的可能值：0 - 触发器；4 - 监控项；5 - LLD 规则。服务事件的可用值：6 - 服务。
	*/
	Object int64 `json:"object,omitempty"`

	/**
	  相关对象的 ID。
	*/
	Objectid string `json:"objectid,omitempty"`

	/**
	  事件是否已被确认。
	*/
	Acknowledged int64 `json:"acknowledged,omitempty"`

	/**
	  创建事件的时间。
	*/
	Clock int64 `json:"clock,omitempty"`

	/**
	  创建事件时的纳秒。
	*/
	Ns int64 `json:"ns,omitempty"`

	/**
	  已解决的事件名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  相关对象的状态。触发和服务事件的可用值：0 - OK;1 - 问题。发现事件的可用值：0 - 主机或服务启动；1 - 主机或服务宕机；2 - 发现主机或服务；3 - 主机或服务丢失。内部事件的可用值：0 - “正常”状态；1 - “未知”或“不支持”状态。此参数不适用于主动式 agent 自动注册事件。
	*/
	Value int64 `json:"value,omitempty"`

	/**
	  事件当前严重性。可用值：0 - not classified（未分类）；1 - information（信息）；2 - warning（警告）；3 - average（一般严重）；4 - high（严重）；5 - disaster（灾难）。
	*/
	Severity int64 `json:"severity,omitempty"`

	/**
	  已恢复事件 ID
	*/
	REventid string `json:"r_eventid,omitempty"`

	/**
	  用于在全局关联规则下覆盖（关闭）当前事件的事件 ID。请参阅 correlationid 以识别确切的关联规则。该参数仅在事件被全局关联规则关闭时定义。
	*/
	CEventid string `json:"c_eventid,omitempty"`

	/**
	  生成关闭问题的关联规则的 ID 。该参数仅在事件被全局关联规则关闭时定义。
	*/
	Correlationid string `json:"correlationid,omitempty"`

	/**
	  手动关闭事件时的用户ID。
	*/
	Userid string `json:"userid,omitempty"`

	/**
	  事件是否被抑制。可用值：0 - 事件处于正常状态；1 - 事件被抑制。
	*/
	Suppressed int64 `json:"suppressed,omitempty"`

	/**
	  具有扩展宏的操作数据。
	*/
	Opdata string `json:"opdata,omitempty"`

	/**
	  活动媒体类型 URL。
	*/
	Urls []interface{}/* TODO */ `json:"urls,omitempty"`
}

type EventGetParam struct {
	GetBaseObject

	/**
	  仅返回具有给定 ID 的事件。
	*/
	Eventids []string `json:"eventids,omitempty"`

	/**
	  仅返回由属于给定主机组的对象创建的事件。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回由属于给定主机的对象创建的事件。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回由给定对象创建的事件。
	*/
	Objectids []string `json:"objectids,omitempty"`

	/**
	  仅返回给定类型的事件。参阅 事件对象页面 获取支持的事件类型列表。默认值：0 - 触发器事件。
	*/
	Source int64 `json:"source,omitempty"`

	/**
	  仅返回由给定类型的对象创建的事件。参阅 事件对象页面 获取支持的事件类型列表。默认值：0 - 触发器。
	*/
	Object int64 `json:"object,omitempty"`

	/**
	  如果设置为 true 仅返回已确认的事件。
	*/
	Acknowledged bool `json:"acknowledged,omitempty"`

	/**
	  true - 只返回被抑制的事件；false - 返回正常状态的事件。
	*/
	Suppressed bool `json:"suppressed,omitempty"`

	/**
	  仅返回具有给定严重性的事件。仅当对象是触发器时才适用。
	*/
	Severities []int64 `json:"severities,omitempty"`

	/**
	  标签搜索规则。可用值：0 - （默认）和/或；2 - 或。
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  仅返回具有给定标签的事件。通过标签进行精确匹配，通过值和操作符进行不区分大小写的匹配。格式：[{"tag"："<tag>"， "value"："<value>"，"operator"："<operator>"}，...]。空数组返回所有事件。可能的运算符类型：0 - （默认） Like（类似于……）；1 - Equal（等于）；2 - Not like（不类似于……）；3 - Not equal（不等于）；4 - Exists（存在）；5 - Not exists（不存在）。
	*/
	Tags []interface{}/* TODO */ `json:"tags,omitempty"`

	/**
	  仅返回 ID 大于或等于给定 ID 的事件。
	*/
	EventidFrom string `json:"eventid_from,omitempty"`

	/**
	  仅返回 ID 小于或等于给定 ID 的事件。
	*/
	EventidTill string `json:"eventid_till,omitempty"`

	/**
	  仅返回在给定时间或在给定时间之后创建的事件。
	*/
	TimeFrom int64 `json:"time_from,omitempty"`

	/**
	  仅返回在给定时间或在给定时间之前创建的事件。
	*/
	TimeTill int64 `json:"time_till,omitempty"`

	/**
	  仅返回从problem_time_from设定的时间开始处于问题状态的事件。仅当事件源是触发器事件且对象为触发时适用。如果指定了 problem_time_till ，则为强制性。
	*/
	ProblemTimeFrom int64 `json:"problem_time_from,omitempty"`

	/**
	  仅返回在problem_time_still设定的时间之前处于问题状态的事件。仅当事件源是触发器事件且对象为触发时适用。如果指定了 problem_time_from ，则为强制性。
	*/
	ProblemTimeTill int64 `json:"problem_time_till,omitempty"`

	/**
	  仅返回具有给定值的事件。
	*/
	Value []int64 `json:"value,omitempty"`

	/**
	  返回 主机 属性，其中 hosts 包含创建事件的对象。仅支持由触发器、监控项或 LLD 规则生成的事件。
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  返回带有创建事件的对象的 'relatedObject ' 属性。返回的对象类型取决于事件类型。
	*/
	SelectRelatedObject map[string][]string `json:"selectRelatedObject,omitempty"`

	/**
	  返回由事件生成的告警的 告警 属性。警报按时间倒序排列。
	*/
	SelectAlerts map[string][]string `json:"select_alerts,omitempty"`

	/**
	  返回带有事件更新的 acknowledges 属性。 事件更新按时间倒序排列。事件更新对象具有以下属性：acknowledgeid - (string) 确认事件的ID；userid - (string) 更新事件的用户ID；eventid - (string) 更新事件的ID；clock - (timestamp) 事件被更新的时间；message - (string) 消息文本；action - (integer) 更新执行的操作，参见 event.acknowledge；old_severity - (integer) 此更新操作之前的事件严重性；new_severity - (integer) 此更新操作后的事件严重性；username - (string) 更新该事件的用户的username（用户名）；name - (string) 更新该事件的用户的name（可见名）；surname - (string) 更新事件的用户的surname（姓）。支持 count。
	*/
	SelectAcknowledges map[string][]string `json:"select_acknowledges,omitempty"`

	/**
	  返回事件标签的 标签 属性。
	*/
	SelectTags map[string][]string `json:"selectTags,omitempty"`

	/**
	  返回维护列表的suppression_data 属性：maintenanceid - (string) 维护期ID；suppress_until - (integer) 事件被抑制的时间。
	*/
	SelectSuppressionData map[string][]string `json:"selectSuppressionData,omitempty"`

	/**
	  按给定属性对结果进行排序。可用值：eventid，objectid 和 clock。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
