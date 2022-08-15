package model

type Item struct {

	/**
	  （只读） 监控项 ID。
	*/
	Itemid string `json:"itemid,omitempty"`

	/**
	  更新监控项的时间间隔。格式为有后缀的时间单位（30s、1m、2h、1d）。可以通过 自定义间隔 进行组合，从而更灵活的设置间隔时间。多个间隔使用分号分隔。可以使用用户宏。单个宏必须填充完整个字段。不支持字段中包含多个宏或文本混合宏。通过编写由正斜杠分隔的两段宏能更加灵活地设置间隔时间 (e.g. {$FLEX_INTERVAL}/{$FLEX_PERIOD}).可选功能 Zabbix trapper 依赖于捕捉监控项，还需要 Zabbix agent（启用状态）能够获取 mqtt.get 密钥。
	*/
	Delay string `json:"delay,omitempty"`

	/**
	  监控项所属的主机或模板的 ID。更新操作中，此字段为只读。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  监控项主机接口的 ID。模板监控项不需要此属性。此属性可用于监控内部、启用状态代理、捕捉、计算、依赖和数据库监控项。
	*/
	Interfaceid string `json:"interfaceid,omitempty"`

	/**
	  监控项关键字。
	*/
	Key string `json:"key_,omitempty"`

	/**
	  监控项名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  监控项类型。可用值：0 - Zabbix agent（zabbix 代理）；2 - Zabbix trapper（zabbix 捕获器）；3 - Simple check（通用检查）；5 - Zabbix internal（zabbix 内部）；7 - Zabbix agent (active)（zabbix agent-工作状态）；9 - Web item（web 监控项）；10 - External check（外部检查）；11 - Database monitor（数据库监控）；12 - IPMI agent（IPMI 代理）；13 - SSH agent（SSH 代理）；14 - Telnet agent（Telnet代理）；15 - Calculated（计算值）；16 - JMX agent（JMX 代理）；17 - SNMP trap（SNMP 捕获）；18 - Dependent item（依赖监控项）；19 - HTTP agent（HTTP 代理）；20 - SNMP agent（SNMP 代理）；21 - Script（脚本）。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  URL 字符串，用于 HTTP agent 监控项类型。支持用户宏编辑、{HOST.IP}、{HOST.CONN}、{HOST.DNS}、{HOST.HOST}、{HOST.NAME}、{ITEM.ID}、{ITEM.KEY}。
	*/
	Url string `json:"url,omitempty"`

	/**
	  监控项数据类型可用值：0 - numeric float（数字浮点型）；1 - character（字符）；2 - log（日志）；3 - numeric unsigned（正数）；4 - text（文本）。
	*/
	ValueType int64 `json:"value_type,omitempty"`

	/**
	  HTTP agent 监控项字段。同样适用 trapper 监控项类型。0 - （默认） 不允许接受传入数据。1 - 允许接受传入数据。
	*/
	AllowTraps int64 `json:"allow_traps,omitempty"`

	/**
	  仅 SSH agent 监控项或 HTTP agent 监控项使用。SSH agent 身份验证方法可用值：0 - （默认） 密码；1 - 公钥。HTTP agent 身份验证方法可用值：0 - （默认） 无认证；1 - basic 认证；2 - NTLM 认证3 - Kerberos 认证；
	*/
	Authtype int64 `json:"authtype,omitempty"`

	/**
	  监控项说明。
	*/
	Description string `json:"description,omitempty"`

	/**
	  （只读） 如果更新监控项时报错，则显示错误信息。
	*/
	Error string `json:"error,omitempty"`

	/**
	  （只读） 监控项源可用值：0 - 平台监控项；4 - 发现监控项。
	*/
	Flags int64 `json:"flags,omitempty"`

	/**
	  HTTP agent 监控项字段。在合并数据时遵循响应重定向。0 - 不遵循重定向；1 - （默认） 遵循重定向。
	*/
	FollowRedirects int64 `json:"follow_redirects,omitempty"`

	/**
	  HTTP agent 监控项字段。具有 HTTP（S）请求头的对象，其中请求头名用作键，请求头值用作值。示例：{ "User-Agent": "Zabbix" }
	*/
	Headers HttpHeaderObject `json:"headers,omitempty"`

	/**
	  历史数据应保留多长时间的单位。同样接受用户宏。默认： 90天。
	*/
	History string `json:"history,omitempty"`

	/**
	  HTTP agent 监控项字段。HTTP（S）agent 连接字符串。
	*/
	HttpProxy string `json:"http_proxy,omitempty"`

	/**
	  由监控项填充的主机清单字段的 ID。请参阅 主机资源清单页 ，了解更多受支持的主机资源清单字段及其 ID 列表。默认： 0。
	*/
	InventoryLink int64 `json:"inventory_link,omitempty"`

	/**
	  IPMI 传感器。仅用于 IPMI 监控项。
	*/
	IpmiSensor string `json:"ipmi_sensor,omitempty"`

	/**
	  JMX agent 自定义连接字符串。默认值：service:jmx:rmi:///jndi/rmi://{HOST.CONN}:{HOST.PORT}/jmxrmi
	*/
	JmxEndpoint string `json:"jmx_endpoint,omitempty"`

	/**
	  （只读） 上一次更新监控项的时间。默认情况下，仅显示过去24小时内的值。您可以通过更改 管理 → 通用 菜单中的 最大历史显示周期 参数值来延长此时间段。
	*/
	Lastclock int64 `json:"lastclock,omitempty"`

	/**
	  （只读） 上一次更新监控项时的纳秒数。默认情况下，仅显示过去24小时内的值。您可以通过更改 管理 → 通用 菜单中的 最大历史显示周期 参数值来延长此时间段。
	*/
	Lastns int64 `json:"lastns,omitempty"`

	/**
	  （只读） 监控项最后获取的值。默认情况下，仅显示过去24小时内的值。您可以通过更改 管理 → 通用 菜单中的 最大历史显示周期 参数值来延长此时间段。
	*/
	Lastvalue string `json:"lastvalue,omitempty"`

	/**
	  日志条目中的时间格式。仅日志监控项使用。
	*/
	Logtimefmt string `json:"logtimefmt,omitempty"`

	/**
	  主监控项 ID。最多允许递归 3 个依赖监控项，且依赖监控项的最大计数为 29999。依赖监控项的先决条件。
	*/
	MasterItemid int64 `json:"master_itemid,omitempty"`

	/**
	  HTTP agent 监控项字段。将响应转换为 JSON 格式。0 - （默认） 原格式；1 - 转换为 JSON 格式。
	*/
	OutputFormat int64 `json:"output_format,omitempty"`

	/**
	  其他参数取决于监控项的类型：- SSH 和 Telnet agent 的可执行脚本；- 数据库监控项的 SQL 查询；- 计算监控项的公式；- 脚本监控项的脚本。
	*/
	Params string `json:"params,omitempty"`

	/**
	  脚本监控项的其他参数。具有 “名称” 和 “值” 属性的对象数组，其中名称必须唯一。
	*/
	Parameters ParameterObject `json:"parameters,omitempty"`

	/**
	  验证密码。用于一般检查、SSH、Telnet、数据库监控、JMX 和 HTTP agent 监控项。当 JMX 使用时，用户名应与密码一起指定，或者两个属性都留空。
	*/
	Password string `json:"password,omitempty"`

	/**
	  HTTP agent 监控项字段。Posts 属性中存储的 post 数据体的类型。0 - （默认） 原始数据；2 - JSON 数据；3 - XML 数据。
	*/
	PostType int64 `json:"post_type,omitempty"`

	/**
	  HTTP agent 监控项字段。HTTP（S）请求正文数据。与 post_type 一同使用。
	*/
	Posts string `json:"posts,omitempty"`

	/**
	  （只读） 监控项的上一个值。默认情况下，仅显示过去24小时内的值。您可以通过更改 管理 → 通用 菜单中的 最大历史显示周期 参数值来延长此时间段。
	*/
	Prevvalue string `json:"prevvalue,omitempty"`

	/**
	  私钥文件的名称。
	*/
	Privatekey string `json:"privatekey,omitempty"`

	/**
	  公钥文件的名称。
	*/
	Publickey string `json:"publickey,omitempty"`

	/**
	  HTTP agent 监控项字段。查询参数。具有 “键”：“值” 对的对象数组，其中值可以是空字符串。
	*/
	QueryFields QueryFieldObject `json:"query_fields,omitempty"`

	/**
	  HTTP agent 监控项字段。请求方法类型：0 - （默认） GET；1 - POST；2 - PUT；3 - HEAD。
	*/
	RequestMethod int64 `json:"request_method,omitempty"`

	/**
	  HTTP agent 监控项字段。具体存储响应的部分。0 - （默认） Body；1 - Headers；2 - Body and headers。对于 request_method HEAD 仅允许方法 1。
	*/
	RetrieveMode int64 `json:"retrieve_mode,omitempty"`

	/**
	  SNMP OID。
	*/
	SnmpOid string `json:"snmp_oid,omitempty"`

	/**
	  HTTP agent 监控项字段。SSL 公钥文件路径。
	*/
	SslCertFile string `json:"ssl_cert_file,omitempty"`

	/**
	  HTTP agent 监控项字段。SSL 私钥文件路径。
	*/
	SslKeyFile string `json:"ssl_key_file,omitempty"`

	/**
	  HTTP agent 监控项字段。SSL 密钥文件的密码。
	*/
	SslKeyPassword string `json:"ssl_key_password,omitempty"`

	/**
	  （只读） 监控项状态。可用值：0 - （默认） 标准；1 - 不受支持。
	*/
	State int64 `json:"state,omitempty"`

	/**
	  监控项状态。可用值：0 - （默认） 启用监控项；1 - 禁用监控项。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  HTTP agent 监控项字段。目标HTTP状态代码的范围，用逗号分隔。还支持将用户宏作为逗号分隔列表的一部分。示例： 200,200-{$M},{$M},200-400
	*/
	StatusCodes string `json:"status_codes,omitempty"`

	/**
	  （只读） 父模板监控项的 ID。提示：使用 hostid 属性指定监控项所属的模板。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  监控项数据轮询请求超时。用于 HTTP agent 和脚本监控项。支持用户宏。默认：3秒。最大值：60秒。
	*/
	Timeout string `json:"timeout,omitempty"`

	/**
	  允许的主机。trapper 监控项或 HTTP agent 监控项使用。
	*/
	TrapperHosts string `json:"trapper_hosts,omitempty"`

	/**
	  趋势数据应保留多长时间的单位。可接受用户宏编辑。默认：365天。
	*/
	Trends string `json:"trends,omitempty"`

	/**
	  值的单位。
	*/
	Units string `json:"units,omitempty"`

	/**
	  用于身份验证的用户名。用于一般检查、SSH、Telnet、数据库监控项、JMX 和 HTTP agent 监控项。SSH 和 Telnet 监控项先决条件。当 JMX 使用时，用户名应与密码一起指定，或者两个属性都留空。
	*/
	Username string `json:"username,omitempty"`

	/**
	  通用唯一标识符，用于将导入的监控项链接到现有监控项。仅模板包含项目使用。如果未指定会自动生成。对于更新操作，此字段为 只读。
	*/
	Uuid string `json:"uuid,omitempty"`

	/**
	  关联值映射 ID。
	*/
	Valuemapid string `json:"valuemapid,omitempty"`

	/**
	  HTTP agent 监控项字段。验证 URL 中的主机名是否位于主机证书的公用名字段或使用者备用名字段中。0 - （默认） 不验证；1 - 验证。
	*/
	VerifyHost int64 `json:"verify_host,omitempty"`

	/**
	  HTTP agent 监控项字段。验证主机证书的真实性。0 - （默认） 不验证；1 - 验证。
	*/
	VerifyPeer int64 `json:"verify_peer,omitempty"`
}

type ItemCreateParam struct {
	Item

	/**
	  监控项 预处理 选项。
	*/
	Preprocessing []*PreprocessingObject `json:"preprocessing,omitempty"`

	/**
	  监控项 标签 。
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type ItemUpdateParam struct {
	Item

	/**
	  监控项预处理 用于替换当前预处理选项。
	*/
	Preprocessing []*PreprocessingObject `json:"preprocessing,omitempty"`

	/**
	  标签 用于替换当前标签。
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type ItemGetParam struct {
	GetBaseObject

	/**
	  仅返回具有给定 ID 的监控项。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  仅返回属于给定组中主机的监控项。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回属于给定模板的监控项。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  仅返回属于给定主机的监控项。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回由给定代理监控的监控项。
	*/
	Proxyids []string `json:"proxyids,omitempty"`

	/**
	  仅返回使用给定主机接口的监控项。
	*/
	Interfaceids []string `json:"interfaceids,omitempty"`

	/**
	  仅返回给定图表中使用的监控项。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  仅返回给定触发器中使用的监控项。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  设置返回结果中包含 web 监控项。
	*/
	Webitems bool `json:"webitems,omitempty"`

	/**
	  如果设置为 true ，仅返回从模板继承的监控项。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  如果设置为 true ，仅返回属于模板的监控项。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  如果设置为 true ，仅返回属于受监控主机的已启用监控项。
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  仅返回属于给定名称的组的监控项。
	*/
	Group string `json:"group,omitempty"`

	/**
	  仅返回属于给定名称的主机的监控项。
	*/
	Host string `json:"host,omitempty"`

	/**
	  标签搜索规则。可用值：0 - (默认) And/Or（与/或）；2 - Or（或）。
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  仅返回给定标签的监控项。根据标签进行精确匹配，并根据运算符对标签进行区分大小写或不区分大小写搜索。格式： [{"tag": "<tag>", "value": "<value>", "operator": "<operator>"}, ...].空数组会返回所有监控项。可选操作类型：0 - (默认) Like（相似）；1 - Equal（相等）；2 - Not like（不相似）；3 - Not equal（不相等）；4 - Exists（存在）；5 - Not exists（不存在）。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  如果设置为 true ，则只返回在触发器中使用的监控项。
	*/
	WithTriggers bool `json:"with_triggers,omitempty"`

	/**
	  返回 主机 属性，其中包含监控项所属主机。
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回 接口 属性，其中包含监控项所使用的主机接口。
	*/
	SelectInterfaces []string `json:"selectInterfaces,omitempty"`

	/**
	  返回 触发器 属性，其中包含监控项所使用的触发器。支持 count 。
	*/
	SelectTriggers []string `json:"selectTriggers,omitempty"`

	/**
	  返回 图表 属性，其中包含监控项所使用的图表。支持 count 。
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  返回 发现规则 属性，其中包含用于创建监控项的 LLD 规则。
	*/
	SelectDiscoveryRule []string `json:"selectDiscoveryRule,omitempty"`

	/**
	  返回 itemDiscovery 属性，其中包含监控项的发现对象。监控项发现对象将监控项链接到创建此监控项的监控项原型中。包含属性如下：itemdiscoveryid - (string) 监控项发现 ID ；itemid - (string) 已发现的监控项 ID ；parent_itemid - (string) 创建监控项的监控项原型的 ID ；key_ - (string) 监控项原型关键字；lastcheck - (timestamp) 监控项最后发现时间；ts_delete - (timestamp) 监控项不再被发现，将被删除的时间。
	*/
	SelectItemDiscovery []string `json:"selectItemDiscovery,omitempty"`

	/**
	  返回 预处理 属性，其中包含监控项的预处理选项。包含属性如下：type - (string) 预处理选项类型：1 - Custom multiplier（自定义乘数）；2 - Right trim（移除右侧空白字符）；3 - Left trim（移除左侧空白字符）；4 - Trim（移除空白字符）；5 - Regular expression matching（正则表达式匹配）；6 - Boolean to decimal（布尔值转换十进制）；7 - Octal to decimal（八进制转换十进制）；8 - Hexadecimal to decimal（十六进制转十进制）；9 - Simple change（先前值到新值的基本变化）；10 - Change per second（每秒钟变化量）；11 - XML XPath（XML 解析）；12 - JSONPath（JSON解析）；13 - In range（生成序列）；14 - Matches regular expression（匹配正则表达式）；15 - Does not match regular expression（不匹配正则表达式）；16 - Check for error in JSON（检查 JSON 错误）；17 - Check for error in XML（检查 XML 错误）；18 - Check for error using regular expression（检查正则表达式使用错误）；19 - Discard unchanged（丢弃重复数据）；20 - Discard unchanged with heartbeat（设置心跳检查周期，丢弃重复数据）；21 - JavaScript（JS格式）；22 - Prometheus pattern（Prometheus 模板）；23 - Prometheus to JSON（Prometheus 转换 JSON）；24 - CSV to JSON（CSV 转换 JSON）；25 - Replace（替换）；26 - Check for not supported value（检查不支持的值）；27 - XML to JSON（XML 转换 JSON）。params - (string) 预处理选项使用的其他参数。多个参数由LF（\n）字符分隔。error_handler - (string) 预处理步骤失败时使用的操作类型：0 - Error message is set by Zabbix server（Zabbix 服务器自带错误消息）；1 - Discard value（丢弃值）；2 - Set custom value（设置自定义值）；3 - Set custom error message（设置自定义错误信息）。error_handler_params - (string) 错误处理器参数。
	*/
	SelectPreprocessing []string `json:"selectPreprocessing,omitempty"`

	/**
	  返回 标签 属性，其中包含监控项的标签。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  返回 映射表 属性，其中包含监控项的映射表。
	*/
	SelectValueMap []string `json:"selectValueMap,omitempty"`

	/**
	  限制子查询返回的记录数。适用以下子查询：selectGraphs - 按 name 对结果进行排序；selectTriggers - 按 description 对结果进行排序。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  按给定的属性对结果进行排序。可用值： itemid, name, key_, delay, history, trends, type 以及 status.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}

type PreprocessingObject struct {

	/**
	  预处理选项类型。可用值：1 - Custom multiplier（自定义乘数）；2 - Right trim（移除右侧空白字符）；3 - Left trim（移除左侧空白字符）；4 - Trim（移除空白字符）；5 - Regular expression matching（正则表达式匹配）；6 - Boolean to decimal（布尔值转换十进制）；7 - Octal to decimal（八进制转换十进制）；8 - Hexadecimal to decimal（十六进制转十进制）；9 - Simple change（先前值到新值的基本变化）；10 - Change per second（每秒钟变化量）；11 - XML XPath（XML 解析）；12 - JSONPath（JSON解析）；13 - In range（生成序列）；14 - Matches regular expression（匹配正则表达式）；15 - Does not match regular expression（不匹配正则表达式）；16 - Check for error in JSON（检查 JSON 错误）；17 - Check for error in XML（检查 XML 错误）；18 - Check for error using regular expression（检查正则表达式使用错误）；19 - Discard unchanged（丢弃重复数据）；20 - Discard unchanged with heartbeat（设置心跳检查周期，丢弃重复数据）；21 - JavaScript（JS格式）；22 - Prometheus pattern（Prometheus 模式）；23 - Prometheus to JSON（Prometheus 转换 JSON）；24 - CSV to JSON（CSV 转换 JSON）；25 - Replace（替换）；26 - Check unsupported（检查不支持的值）；27 - XML to JSON（XML 转换 JSON）。
	*/
	Type int64 `json:"type,omitempty"`

	/**

	 */
	Params string `json:"params,omitempty"`

	/**
	  预处理步骤失败时使用的操作类型：可用值：0 - Error message is set by Zabbix server（Zabbix 服务器自带错误消息）；1 - Discard value（丢弃值）；2 - Set custom value（设置自定义值）；3 - Set custom error message（设置自定义错误信息）。
	*/
	ErrorHandler int64 `json:"error_handler,omitempty"`

	/**
	  错误处理器参数。与 error_handler 搭配使用。如果 error_handler 类型为 0 或 1，此值必须为空。如果 error_handler 类型为 2，此值可以为空。如果 error_handler 类型为 3，此值不能为空。
	*/
	ErrorHandlerParams string `json:"error_handler_params,omitempty"`
}
