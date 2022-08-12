package model

type MediaType struct {

	/**
	  （只读） 媒介类型ID。
	*/
	Mediatypeid string `json:"mediatypeid,omitempty"`

	/**
	  媒介类型使用的传输。可用值：0 - 电子邮件；1 - 脚本；2 - SMS；4 - Webhook。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  对于脚本媒介类型，exec_path 包含执行脚本的名称。对于脚本媒介类型是必需的。
	*/
	ExecPath string `json:"exec_path,omitempty"`

	/**
	  GSM 调制解调器的串行设备名称。对于 SMS 媒介类型是必需的。
	*/
	GsmModem string `json:"gsm_modem,omitempty"`

	/**
	  身份验证密码。用于电子邮件媒介类型。
	*/
	Passwd string `json:"passwd,omitempty"`

	/**
	  用于发送通知的电子邮件地址。电子邮件媒介类型所必需的。
	*/
	SmtpEmail string `json:"smtp_email,omitempty"`

	/**
	  SMTP HELO。对于电子邮件媒介类型是必需的。
	*/
	SmtpHelo string `json:"smtp_helo,omitempty"`

	/**
	  SMTP 服务器。对于电子邮件媒介类型是必需的。
	*/
	SmtpServer string `json:"smtp_server,omitempty"`

	/**
	  要连接的 SMTP 服务器端口。
	*/
	SmtpPort int64 `json:"smtp_port,omitempty"`

	/**
	  要使用的 SMTP 连接安全级别。可用值：0 - 无；1 - STARTTLS；2 - SSL/TLS。
	*/
	SmtpSecurity int64 `json:"smtp_security,omitempty"`

	/**
	  用于 SMTP 的 SSL 验证主机。可用值：0 - 否；1 - 是。
	*/
	SmtpVerifyHost int64 `json:"smtp_verify_host,omitempty"`

	/**
	  用于 SMTP 的 SSL 验证对等点。可用值：0 - 否；1 - 是。
	*/
	SmtpVerifyPeer int64 `json:"smtp_verify_peer,omitempty"`

	/**
	  要使用的 SMTP 身份验证方法。可用值：0 - 无；1 - 普通密码。
	*/
	SmtpAuthentication int64 `json:"smtp_authentication,omitempty"`

	/**
	  媒介类型是否启用。可用值：0 - （默认）启用；1 - 禁用。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  用户名。用于电子邮件媒介类型。
	*/
	Username string `json:"username,omitempty"`

	/**
	  脚本参数。每个参数都以新的换行符结尾。
	*/
	ExecParams string `json:"exec_params,omitempty"`

	/**
	  可以并行处理的最大警报数。SMS 的可用值：1 - （默认）其他媒介类型的可用值：0-100
	*/
	Maxsessions int64 `json:"maxsessions,omitempty"`

	/**
	  尝试发送警报的最大次数。可用值：1-100默认值：3
	*/
	Maxattempts int64 `json:"maxattempts,omitempty"`

	/**
	  重试尝试之间的间隔。接受带后缀的秒和时间单位。<br可用值：0-1h默认值：10s
	*/
	AttemptInterval string `json:"attempt_interval,omitempty"`

	/**
	  消息格式。可用值：0 - 纯文本；1 - （默认） html。
	*/
	ContentType int64 `json:"content_type,omitempty"`

	/**
	  媒介类型 webhook 脚本 javascript 正文。
	*/
	Script string `json:"script,omitempty"`

	/**
	  媒介类型 webhook 脚本超时。接受带后缀的秒和时间单位。可用值：1-60s默认值：30s
	*/
	Timeout string `json:"timeout,omitempty"`

	/**
	  定义 webhook 脚本响应是否应被解释为标签，并且这些标签应添加到相关事件中。可用值：0 - （默认） 忽略 webhook 脚本response.1 - 将 webhook 脚本响应作为标签处理。
	*/
	ProcessTags int64 `json:"process_tags,omitempty"`

	/**
	  在 problem.get 和 event.get 属性 urls 中显示媒介类型条目。可用值：0 - （默认） 做不添加 urls 条目。1 - 将媒介类型添加到 urls 属性。
	*/
	ShowEventMenu int64 `json:"show_event_menu,omitempty"`

	/**
	  在 problem.get 和 event.get 的 urls 属性中定义媒介类型条目的 url 属性。
	*/
	EventMenuUrl string `json:"event_menu_url,omitempty"`

	/**
	  在 problem.get 和 event.get 的 urls 属性中定义媒介类型条目的 name 属性。
	*/
	EventMenuName string `json:"event_menu_name,omitempty"`

	/**
	  webhook 输入参数的数组。
	*/
	Parameters []interface{}/* TODO */ `json:"parameters,omitempty"`

	/**
	  媒介类型描述。
	*/
	Description string `json:"description,omitempty"`
}

type MediaTypeCreateParam struct {
	MediaType

	/**
	  要为媒体类型创建的 Webhook 参数 。
	*/
	Parameters []interface{}/* TODO */ `json:"parameters,omitempty"`

	/**
	  为媒体类型创建的消息模板 。
	*/
	MessageTemplates []interface{}/* TODO */ `json:"message_templates,omitempty"`
}

type MediaTypeUpdateParam struct {
	MediaType

	/**
	  Webhook 参数 用于替换当前的 webhook 参数。
	*/
	Parameters []interface{}/* TODO */ `json:"parameters,omitempty"`

	/**
	  消息模板 用于替换当前的消息模板。
	*/
	MessageTemplates []interface{}/* TODO */ `json:"message_templates,omitempty"`
}

type MediaTypeGetParam struct {
	GetBaseObject

	/**
	  只返回带有给定ID列表的媒介类型。
	*/
	Mediatypeids []string `json:"mediatypeids,omitempty"`

	/**
	  仅返回给定媒介使用的媒介类型。
	*/
	Mediaids []string `json:"mediaids,omitempty"`

	/**
	  仅返回给定用户使用的媒介类型。
	*/
	Userids []string `json:"userids,omitempty"`

	/**
	  返回一个包含消息模板消息数组的属性。
	*/
	SelectMessageTemplates map[string][]string `json:"selectMessageTemplates,omitempty"`

	/**
	  返回使用媒介类型的用户属性。
	*/
	SelectUsers map[string][]string `json:"selectUsers,omitempty"`

	/**
	  按给定属性对结果进行排序。可用值：mediatypeid。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
