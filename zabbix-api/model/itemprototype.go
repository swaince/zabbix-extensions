package model

type ItemPrototype struct {

	/**
	  (只读) 监控项原型的ID
	*/
	Itemid string `json:"itemid,omitempty"`

	/**
	  监控项原型的更新间隔。接受秒或带有后缀的时间单位(30s,1m,2h,1d)。可灵活的指定一个或多个时间间隔自定义间隔或计划， 多个间隔用分号分割。可以使用用户宏和LLD宏. 一个宏必须填充整个字段. 不支持字段中的多个宏或与文本混合的宏.灵活的间隔可以写成两个用正斜杠分隔的宏 (例如 {$FLEX_INTERVAL}/{$FLEX_PERIOD}).Zabbix trapper, 依赖项和for Zabbix agent (active)可选 mqtt.get key。
	*/
	Delay string `json:"delay,omitempty"`

	/**
	  监控项原型的主机ID.对于更新操作，此字段为只读。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  监控项所属的LLD规则的ID.对于更新操作，此字段为只读。
	*/
	Ruleid string `json:"ruleid,omitempty"`

	/**
	  监控项原型的主机接口的ID。 仅用于监控项原型。对于Zabbix agent (active), Zabbix internal, Zabbix trapper,依赖项, 数据库监控项和可计算监控项原型是可选的。
	*/
	Interfaceid string `json:"interfaceid,omitempty"`

	/**
	  监控项原型key.
	*/
	Key string `json:"key,omitempty"`

	/**
	  监控项原型名字。
	*/
	Name string `json:"name,omitempty"`

	/**
	  监控项原型类型.可选值:0 - Zabbix agent;2 - Zabbix trapper;3 - simple check;5 - Zabbix internal;7 - Zabbix agent (active);10 - external check;11 - database monitor;12 - IPMI agent;13 - SSH agent;14 - TELNET agent;15 - calculated;16 - JMX agent;17 - SNMP trap;18 - Dependent item;19 - HTTP agent;20 - SNMP agent;21 - Script.
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  URL仅用于HTTP agent监控项原型. 支持 LLD宏,用户宏, {HOST.IP}, {HOST.CONN}, {HOST.DNS}, {HOST.HOST}, {HOST.NAME}, {ITEM.ID}, {ITEM.KEY}.
	*/
	Url string `json:"url,omitempty"`

	/**
	  监控项原型的值的类型可选项:0 - 浮点值;1 - 字符串;2 - 日志;3 - 无符号整数;4 - 文本。
	*/
	ValueType int64 `json:"value_type,omitempty"`

	/**
	  HTTP客户端监控项原型字段. 允许trapper监控项类型填充值。0 - (默认值) 不允许接受传入数据、1 - 允许接受传入数据。
	*/
	AllowTraps int64 `json:"allow_traps,omitempty"`

	/**
	  仅允许SSH客户端监控项原型或者HTTP客户端监控项原型。SSH客户端验证方法可选值:0 - (默认值) 密码;1 - 秘钥.HTTP客户端验证方法可选值:0 - (默认值) none1 - basic2 - NTLM3 - Kerberos
	*/
	Authtype int64 `json:"authtype,omitempty"`

	/**
	  监控项原型的描述.
	*/
	Description string `json:"description,omitempty"`

	/**
	  HTTP客户端监控现原型字段。数据重定向.0 - 不要重定向。1 - (默认值) 重定向。
	*/
	FollowRedirects int64 `json:"follow_redirects,omitempty"`

	/**
	  HTTP客户端监控现原型字段。 HTTP（S）请求头的对象，其中，头部名称用作键，头部值用作值.案例:{ "User-Agent": "Zabbix" }
	*/
	Headers HttpHeaderObject `json:"headers,omitempty"`

	/**
	  历史数据的可存储时间。可接受用户宏和LLD宏。默认值: 90d.
	*/
	History string `json:"history,omitempty"`

	/**
	  HTTP客户端监控现原型字段。 HTTP(S) 代理地址字符串.
	*/
	HttpProxy string `json:"http_proxy,omitempty"`

	/**
	  IPMI传感器. 仅用于IPMI监控现原型.
	*/
	IpmiSensor string `json:"ipmi_sensor,omitempty"`

	/**
	  JMX客户端连接地址。默认值:service:jmx:rmi:///jndi/rmi://{HOST.CONN}:{HOST.PORT}/jmxrmi
	*/
	JmxEndpoint string `json:"jmx_endpoint,omitempty"`

	/**
	  日志条目中的时间格式。仅由日志监控项原型使用。
	*/
	Logtimefmt string `json:"logtimefmt,omitempty"`

	/**
	  主监控项ID.最多可递归3个相关监控项目和监控项目原型，且相关监控项目和监控项目原型的最大计数等于29999。
	*/
	MasterItemid int64 `json:"master_itemid,omitempty"`

	/**
	  HTTP客户端监控现原型字段。可将响应内容转换为JSON0 - (默认值) 原始数据.1 - 转JSON.
	*/
	OutputFormat int64 `json:"output_format,omitempty"`

	/**
	  其他参数取决于监控现原型的类型:- SSH和Telnet监控项原型的执行脚本;- 数据库监控项原型的sql查询;- 可计算监控项目原型的公式。
	*/
	Params string `json:"params,omitempty"`

	/**
	  脚本监控项原型的其他参数。具有“名称”和“值”属性的对象数组，其中名称必须唯一。
	*/
	Parameters ParameterObject `json:"parameters,omitempty"`

	/**
	  验证密码。用于简单检查、SSH、Telnet、数据库监控项、JMX和HTTP客户端监控项原型。
	*/
	Password string `json:"password,omitempty"`

	/**
	  HTTP客户端监控现原型字段。 posts属性中存储的post数据体的类型。0 - (默认值) Raw数据。2 - JSON数据。3 - XML数据。
	*/
	PostType int64 `json:"post_type,omitempty"`

	/**
	  HTTP客户端监控现原型字段。HTTP(S)请求内容. 使用post类型.
	*/
	Posts string `json:"posts,omitempty"`

	/**
	  私钥文件的名称。
	*/
	Privatekey string `json:"privatekey,omitempty"`

	/**
	  公钥文件的名称。
	*/
	Publickey string `json:"publickey,omitempty"`

	/**
	  HTTP客户端监控现原型字段。 查询参数。具有“键”：“值”对的对象数组，其中值可以是空字符串。
	*/
	QueryFields QueryFieldObject `json:"query_fields,omitempty"`

	/**
	  HTTP客户端监控现原型字段。请求方法的类型。0 - (默认值) GET1 - POST2 - PUT3 - HEAD
	*/
	RequestMethod int64 `json:"request_method,omitempty"`

	/**
	  HTTP客户端监控现原型字段。 响应返数据的哪一部分。0 - (默认值) Body.1 - Headers.2 - body和headers.值为1时，返回request的HEAD。
	*/
	RetrieveMode int64 `json:"retrieve_mode,omitempty"`

	/**
	  SNMP OID.
	*/
	SnmpOid string `json:"snmp_oid,omitempty"`

	/**
	  HTTP代理监控项原型字段。公共SSL密钥文件路径。
	*/
	SslCertFile string `json:"ssl_cert_file,omitempty"`

	/**
	  HTTP代理监控项原型字段。私有SSL密钥文件路径。
	*/
	SslKeyFile string `json:"ssl_key_file,omitempty"`

	/**
	  HTTP代理监控项原型字段。 SSL密钥文件的密码。
	*/
	SslKeyPassword string `json:"ssl_key_password,omitempty"`

	/**
	  监控项原型状态。可选值:0 - (默认值) 激活监控项原型;1 - 禁用监控项原型;3 - 不支持监控项原型.
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  HTTP代理监控项原型字段。 指定HTTP状态码范围以，分隔。还支持用户宏或LLD宏作为逗号分隔列表的一部分。例如: 200,200-{$M},{$M},200-400
	*/
	StatusCodes string `json:"status_codes,omitempty"`

	/**
	  监控项原型的模板（只读）id.
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  监控项数据轮询请求超时。用于HTTP代理和脚本监控项原型. 支持用户宏或LLD宏。默认值: 3s最小值: 60s
	*/
	Timeout string `json:"timeout,omitempty"`

	/**
	  允许的主机。 用于trapper监控项原型或者HTTP监控现原型.
	*/
	TrapperHosts string `json:"trapper_hosts,omitempty"`

	/**
	  趋势数据存储多长时间。还接受用户宏和LLD宏。默认值: 365d.
	*/
	Trends string `json:"trends,omitempty"`

	/**
	  值的单位。
	*/
	Units string `json:"units,omitempty"`

	/**
	  验证的用户名. 用于简单检查、SSH、Telnet、数据库监控项、JMX和HTTP代理监控项原型。SSH和Telnet监控项目原型需要。
	*/
	Username string `json:"username,omitempty"`

	/**
	  通用唯一标识符，用于将导入的监控项目原型链接到现有的监控项目原型。仅用于模板上的监控项目原型。如果没有给出则自动生成。对于更新操作，此字段为只读.
	*/
	Uuid string `json:"uuid,omitempty"`

	/**
	  关联值映射值的ID.
	*/
	Valuemapid string `json:"valuemapid,omitempty"`

	/**
	  HTTP代理监控项原型字段。 验证URL中的主机名位于主机证书的公用名字段或使用者备用名字段中.0 - (默认值) 不验证.1 - 验证.
	*/
	VerifyHost int64 `json:"verify_host,omitempty"`

	/**
	  HTTP代理监控项原型字段。 验证主机证书是否真实.0 - (默认值) 不验证.1 - 验证.
	*/
	VerifyPeer int64 `json:"verify_peer,omitempty"`

	/**
	  监控项原型自动发现状态.可选值:0 - (默认值) 新监控项目将被发现;1 - 新项目将不会被发现，现有项目将被标记为丢失。
	*/
	Discover int64 `json:"discover,omitempty"`
}

type ItemPrototypeCreateParam struct {
	ItemPrototype

	/**
	  监控项所属的LLD规则ID
	*/
	Ruleid string `json:"ruleid,omitempty"`

	/**
	  监控项原型预处理 选项。
	*/
	Preprocessing []*PreprocessingObject `json:"preprocessing,omitempty"`

	/**
	  监控项原型标签。
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type ItemPrototypeUpdateParam struct {
	ItemPrototype

	/**
	  要替换当前应用的应用的ID。
	*/
	Applications []interface{}/* TODO */ `json:"applications,omitempty"`

	/**
	  要替换的当前监控项预处理选项。
	*/
	Preprocessing []*PreprocessingObject `json:"preprocessing,omitempty"`
}

type ItemPrototypeGetParam struct {
	GetBaseObject

	/**
	  只返回属于给定 LLD 规则的监控项原型。
	*/
	Discoveryids []string `json:"discoveryids,omitempty"`

	/**
	  只返回给定图形原型中使用的监控项原型。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回属于给定主机的监控项原型。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  如果设置为 true，则仅返回从模板继承的监控项原型。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  只返回给定 ID 的监控项原型。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  如果设置为 true，则仅返回属于受监控主机的已启用监控项原型。
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  如果设置为 true，则仅返回属于模板的监控项原型。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  只返回属于给定模板的监控项原型。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  只返回给定触发器原型中使用的监控项原型。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  返回带有监控项原型所属的低级发现规则的 discoveryRule 属性。
	*/
	SelectDiscoveryRule []string `json:"selectDiscoveryRule,omitempty"`

	/**
	  返回一个 manual/api/reference/graphprototype/object#graph_prototype 属性，其中包含使用监控项原型的图形原型。< br>支持count。
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  返回一个 hosts 属性，其中包含监控项原型所属的主机数组。
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回 tags 属性中的监控项原型标签。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  返回一个 triggers 属性，其中包含使用监控项原型的触发器原型。支持count。
	*/
	SelectTriggers []string `json:"selectTriggers,omitempty"`

	/**
	  返回带有监控项预处理选项的 preprocessing 属性。它具有以下属性：type - (string) 预处理选项类型：1 - 自定义乘数;2 - 右修剪;3 - 左修剪;4 - 修剪;5 - 正则表达式匹配;< br>6 - 布尔到十进制；7 - 八进制到十进制；8 - 十六进制到十进制；9 - 简单更改；10 - 每秒更改；11 - XML XPath ;12 - JSONPath;13 - 在范围内;14 - 匹配正则表达式;15 - 不匹配正则表达式;16 - 检查 JSON 中的错误; 17 - 检查 XML 中的错误；18 - 使用正则表达式检查错误；19 - 丢弃未更改；20 - 丢弃未更改的心跳；21 - JavaScript；22 - Prometheus 模式；23 - Prometheus 到 JSON；24 - CSV 到 JSON；25 - 替换；26 - 检查不支持的值；27- XML 到 JSON。<br >params - (str ing) 预处理选项使用的附加参数。多个参数由 LF (\n) 字符分隔。error_handler - (string) 预处理步骤失败时使用的操作类型：0 - Zabbix 服务器设置错误消息；<br >1 - 丢弃值；2 - 设置自定义值；3 - 设置自定义错误消息。error_handler_params - (string) 错误处理程序参数。
	*/
	SelectPreprocessing []string `json:"selectPreprocessing,omitempty"`

	/**
	  返回带有监控项原型值映射的 valuemap 属性。
	*/
	SelectValueMap []string `json:"selectValueMap,omitempty"`

	/**
	  限制子选择返回的记录数。适用于以下子选择：selectGraphs - 结果将按name排序；selectTriggers - 结果将按“描述”排序。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  按给定属性对结果进行排序。可能的值有：itemid、name、key_、delay、type 和 status。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
