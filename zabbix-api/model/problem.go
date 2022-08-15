package model

type Problem struct {

	/**
	  问题事件的 ID。
	*/
	Eventid string `json:"eventid,omitempty"`

	/**
	  问题事件的类型。可用值：0 - 由触发器创建的事件；3 - 内部事件；4 - 在服务状态更新时创建的事件。
	*/
	Source int64 `json:"source,omitempty"`

	/**
	  与问题事件相关的对象类型。触发事件的可用值：0 - 触发。内部事件的可用值： 0 - 触发器；4 - 项目；5 - LLD 规则。服务事件的可用值：6 - 服务。
	*/
	Object int64 `json:"object,omitempty"`

	/**
	  相关对象的ID。
	*/
	Objectid string `json:"objectid,omitempty"`

	/**
	  创建问题事件的时间。
	*/
	Clock int64 `json:"clock,omitempty"`

	/**
	  创建问题事件时的纳秒。
	*/
	Ns int64 `json:"ns,omitempty"`

	/**
	  恢复事件 ID。
	*/
	REventid string `json:"r_eventid,omitempty"`

	/**
	  创建恢复事件的时间。
	*/
	RClock int64 `json:"r_clock,omitempty"`

	/**
	  创建恢复事件时的纳秒。
	*/
	RNs int64 `json:"r_ns,omitempty"`

	/**
	  如果此事件被全局关联规则恢复，则关联规则 ID。
	*/
	Correlationid string `json:"correlationid,omitempty"`

	/**
	  如果问题是手动关闭的，用户 ID。
	*/
	Userid string `json:"userid,omitempty"`

	/**
	  已解决的问题名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  问题的确认状态。可用值：0 - 未确认；1 - 已确认。
	*/
	Acknowledged int64 `json:"acknowledged,omitempty"`

	/**
	  问题当前严重性。可用值：0 - 未分类；1 - 信息；2 - 警告；3 - 平均； 4 - 高；5 - 灾难。
	*/
	Severity int64 `json:"severity,omitempty"`

	/**
	  问题是否被抑制。可用值：0 - 问题处于正常状态；1 - 问题被抑制。
	*/
	Suppressed int64 `json:"suppressed,omitempty"`

	/**
	  带有扩展宏的操作数据。
	*/
	Opdata string `json:"opdata,omitempty"`

	/**
	  活动媒介类型 URL。
	*/
	Urls []*MediaTypeUrlObject `json:"urls,omitempty"`
}

type ProblemGetParam struct {
	GetBaseObject

	/**
	  仅返回给定 ID 的问题。
	*/
	Eventids []string `json:"eventids,omitempty"`

	/**
	  仅返回属于给定主机组的对象创建的问题。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回属于给定主机的对象创建的问题。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回给定对象创建的问题。
	*/
	Objectids []string `json:"objectids,omitempty"`

	/**
	  仅返回给定类型的问题。有关支持的事件类型列表，请参阅问题事件对象页面。默认值：0 - 由触发器创建的问题。
	*/
	Source int64 `json:"source,omitempty"`

	/**
	  仅返回由给定类型的对象创建的问题。有关支持的对象类型列表，请参阅问题事件对象页面。默认值：0 - 触发器。
	*/
	Object int64 `json:"object,omitempty"`

	/**
	  true - 仅返回已确认的问题；false - 仅未确认。
	*/
	Acknowledged bool `json:"acknowledged,omitempty"`

	/**
	  true - 仅返回被抑制的问题；false - 返回正常状态的问题。
	*/
	Suppressed bool `json:"suppressed,omitempty"`

	/**
	  仅返回给定事件严重性的问题。仅当对象是触发器时才适用。
	*/
	Severities []int64 `json:"severities,omitempty"`

	/**
	  标签搜索规则。可用值：0 - （默认） And/Or；2 - Or。
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  仅返回给定标签的问题。按标签精确匹配，按值和运算符不区分大小写搜索。格式：[{"tag": "<tag>", "value": "<value>", "operator": "<operator> "}, ...]。空数组返回所有问题。可能的运算符类型：0 - （默认）相似；1 - 等于；2 - 不相似；3 - 不等于；4 - 存在；5 - 不存在。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  true - 返回 PROBLEM 和最近解决的问题（取决于 N 秒的 Display OK 触发器）。默认值：false - 仅未解决的问题。
	*/
	Recent bool `json:"recent,omitempty"`

	/**
	  仅返回 ID 大于或等于给定 ID 的问题。
	*/
	EventidFrom string `json:"eventid_from,omitempty"`

	/**
	  仅返回 ID 小于或等于给定 ID 的问题。
	*/
	EventidTill string `json:"eventid_till,omitempty"`

	/**
	  仅返回在给定时间之后或在给定时间创建的问题。
	*/
	TimeFrom int64 `json:"time_from,omitempty"`

	/**
	  仅返回在给定时间之前或在给定时间创建的问题。
	*/
	TimeTill int64 `json:"time_till,omitempty"`

	/**
	  返回带有问题更新的 acknowledges 属性。问题更新按时间倒序排列。问题更新对象具有以下属性：acknowledgeid - (string) 更新的 ID；userid - (string) 更新事件的用户 ID；eventid - (string) 更新事件的 ID；clock - (timestamp) 事件更新的时间； message - (string) 消息文本；action - (integer) 更新操作类型（参见 event.acknowledge）；old_severity - 此更新操作之前的(integer) 事件严重性；new_severity - 此更新操作后的(integer) 事件严重性；支持count。
	*/
	SelectAcknowledges []string `json:"selectAcknowledges,omitempty"`

	/**
	  返回带有问题标签的 tags 属性。输出格式：[{"tag": "<tag>", "value": "<value>"}, ...]。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  返回带有维护列表的 suppression_data 属性：maintenanceid - (string) 维护的 ID；suppress_until - (integer) 出现问题的时间被压制。
	*/
	SelectSuppressionData []string `json:"selectSuppressionData,omitempty"`

	/**
	  按给定的属性对结果进行排序。可用值：eventid。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
