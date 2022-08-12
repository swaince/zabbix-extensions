package model

type Alert struct {

	/**
	  告警 ID。
	*/
	Alertid string `json:"alertid,omitempty"`

	/**
	  生成告警的动作 ID。
	*/
	Actionid string `json:"actionid,omitempty"`

	/**
	  告警类型。可用值：0 - 消息；1 - 远程命令。
	*/
	Alerttype int64 `json:"alerttype,omitempty"`

	/**
	  告警生成的时间。
	*/
	Clock int64 `json:"clock,omitempty"`

	/**
	  发送消息或运行命令时出现问题时的错误文本。
	*/
	Error string `json:"error,omitempty"`

	/**
	  告警生成期间的动作升级步骤。
	*/
	EscStep int64 `json:"esc_step,omitempty"`

	/**
	  触发动作的事件 ID。
	*/
	Eventid string `json:"eventid,omitempty"`

	/**
	  用于发送消息的媒体类型 ID。
	*/
	Mediatypeid string `json:"mediatypeid,omitempty"`

	/**
	  消息文本。用于消息告警。
	*/
	Message string `json:"message,omitempty"`

	/**
	  Zabbix 尝试发送消息的次数。
	*/
	Retries int64 `json:"retries,omitempty"`

	/**
	  地址，用户名或接收者的其他标识符。用于消息 alert。
	*/
	Sendto string `json:"sendto,omitempty"`

	/**
	  指示动作操作是否已成功执行的状态。消息告警的可用值：0 - 消息未发送。1 - 消息已发送。2 - 经多次重试后失败。3 - 告警管理员尚未处理的新告警。命令告警可用值：0 - 命令没有运行。1 - 命令运行成功。2 - 尝试在 Zabbix agent 上运行命令但不可用。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  消息主题。用于消息告警。
	*/
	Subject string `json:"subject,omitempty"`

	/**
	  消息发送到的用户的 ID。
	*/
	Userid string `json:"userid,omitempty"`

	/**
	  生成告警的问题事件的 ID。
	*/
	PEventid string `json:"p_eventid,omitempty"`

	/**
	  生成告警的确认 ID。
	*/
	Acknowledgeid string `json:"acknowledgeid,omitempty"`
}

type AlertGetParam struct {
	GetBaseObject

	/**
	  仅返回给定 ID 的告警。
	*/
	Alertids []string `json:"alertids,omitempty"`

	/**
	  仅返回给定动作生成的告警。
	*/
	Actionids []string `json:"actionids,omitempty"`

	/**
	  仅返回给定事件生成的告警。
	*/
	Eventids []string `json:"eventids,omitempty"`

	/**
	  仅返回来自指定主机组的对象生成的告警。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回来自指定主机的对象生成的告警。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回使用给定媒介类型的消息告警。
	*/
	Mediatypeids []string `json:"mediatypeids,omitempty"`

	/**
	  仅返回指定对象生成的告警。
	*/
	Objectids []string `json:"objectids,omitempty"`

	/**
	  仅返回发送给指定用户的消息告警。
	*/
	Userids []string `json:"userids,omitempty"`

	/**
	  仅返回由与给定类型的对象相关的事件生成的告警。参考 event "object" 获取受支持的对象类型列表。默认： 0 - 触发器。
	*/
	Eventobject int64 `json:"eventobject,omitempty"`

	/**
	  仅返回由给定类型的事件生成的告警。参考 event "source" 获取受支持的事件类型列表。默认： 0 - 触发器事件。
	*/
	Eventsource int64 `json:"eventsource,omitempty"`

	/**
	  仅返回在给定时间之后生成的告警。
	*/
	TimeFrom int64 `json:"time_from,omitempty"`

	/**
	  仅返回在给定时间之前生成的告警。
	*/
	TimeTill int64 `json:"time_till,omitempty"`

	/**
	  在 hosts 属性中返回触发动作操作的主机。
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  在 mediatypes 属性中以数组形式返回消息告警的媒介类型。
	*/
	SelectMediatypes map[string][]string `json:"selectMediatypes,omitempty"`

	/**
	  在 users 属性中以数组形式返回消息发送到的用户。
	*/
	SelectUsers map[string][]string `json:"selectUsers,omitempty"`

	/**
	  根据给定的属性排序结果。可用值： alertid, clock, eventid, mediatypeid, sendto 和 status。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
