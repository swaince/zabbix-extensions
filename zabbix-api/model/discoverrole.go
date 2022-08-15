package model

type DiscoverRole struct {

	/**
	  （只读） LLD规则ID。
	*/
	Itemid string `json:"itemid,omitempty"`

	/**
	  LLD 规则的更新间隔。接受带后缀的秒或时间单位，可以带或不带一个或多个 自定义间隔，它们由灵活间隔和调度间隔组成，作为序列化字符串。还接收用户宏。灵活的间隔可以写成两个由正斜杠分隔的宏。间隔由分号分隔。对于 Zabbix trapper、依赖项和带有 mqtt.get 键值的 Zabbix agent (active) 是可选的。
	*/
	Delay string `json:"delay,omitempty"`

	/**
	  LLD规则的主机ID。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  LLD规则的主机接口ID。只用于主机LLD规则。Zabbix agent (active)的选项, Zabbix internal, Zabbix trapper和数据库监控 LLD规则。
	*/
	Interfaceid string `json:"interfaceid,omitempty"`

	/**
	  LLD规则键值。
	*/
	Key string `json:"key_,omitempty"`

	/**
	  LLD规则名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  LLD规则类型。可用值：0 - Zabbix agent；2 - Zabbix trapper；3 - simple check；5 - Zabbix internal；7 - Zabbix agent (active)；10 - external check；11 - database monitor；12 - IPMI agent；13 - SSH agent；14 - TELNET agent；16 - JMX agent；18 - Dependent item；19 - HTTP agent；20 - SNMP agent；21 - Script。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  URL字符串，HTTP代理LLD规则请求。 支持用户宏，{HOST.IP}， {HOST.CONN}，{HOST.DNS}，{HOST.HOST}，{HOST.NAME}，{ITEM.ID}，{ITEM.KEY}。
	*/
	Url string `json:"url,omitempty"`

	/**
	  HTTP代理LLD规则字段。也允许填充值作为trapper监控项类型。0 - （默认值） 不允许接收传入数据；1 - 允许接收传入数据。
	*/
	AllowTraps int64 `json:"allow_traps,omitempty"`

	/**
	  只有SSH客户端或者HTTP代理的LLD规则能使用.SSH客户端认证方法可用值：0 - （默认值） password；1 - public key。HTTP客户端认证方法可用值：0 - （默认值） none；1 - basic；2 - NTLM。
	*/
	Authtype int64 `json:"authtype,omitempty"`

	/**
	  LLD规则描述。
	*/
	Description string `json:"description,omitempty"`

	/**
	  （只读） 更新 LLD 规则时出现问题时的错误文本。
	*/
	Error string `json:"error,omitempty"`

	/**
	  HTTP代理LLD规则字段。接收数据时进行重定向。0 - 不重定向；1 - （默认值） 重定向。
	*/
	FollowRedirects int64 `json:"follow_redirects,omitempty"`

	/**
	  HTTP代理LLD规则字段。带有 HTTP(S) 请求标头的对象，其中头部名称用作键，头部值用作值。例如：{ "User-Agent": "Zabbix" }
	*/
	Headers interface{}/* TODO */ `json:"headers,omitempty"`

	/**
	  HTTP代理LLD规则字段。HTTP(S)代理连接地址。
	*/
	HttpProxy string `json:"http_proxy,omitempty"`

	/**
	  IPMI 传感器。只用于IPMI LLD规则。
	*/
	IpmiSensor string `json:"ipmi_sensor,omitempty"`

	/**
	  JMX agent自定义的连接地址。默认值：service:jmx:rmi:///jndi/rmi://{HOST.CONN}:{HOST.PORT}/jmxrmi
	*/
	JmxEndpoint string `json:"jmx_endpoint,omitempty"`

	/**
	  监控项的过期时间。接受秒、带后缀的时间单位和用户宏。默认值：30d。
	*/
	Lifetime string `json:"lifetime,omitempty"`

	/**
	  主监控项ID。最多允许递归 3 个依赖监控项，依赖监控项的最大计数等于 999。发现规则不能是另一个发现规则的主监控项。依赖监控项是必需的。
	*/
	MasterItemid int64 `json:"master_itemid,omitempty"`

	/**
	  HTTP代理LLD规则字段。 返回内容格式化JSON输出。0 - （默认值） 不转换；1 - 转JSON。
	*/
	OutputFormat int64 `json:"output_format,omitempty"`

	/**
	  附加参数取决于 LLD 规则的类型：- 执行SSH 和 Telnet 脚本的LLD 规则；- 数据库监控SQL查询SQL的LLD规则；- 可计算LLD规则。
	*/
	Params string `json:"params,omitempty"`

	/**
	  除了LLD规则脚本类型参数之外。还有“名称”和“值”属性的对象数组，其中名称必须是唯一的。
	*/
	Parameters []interface{}/* TODO */ `json:"parameters,omitempty"`

	/**
	  密码认证。适用于简单检查、SSH、Telnet、数据库监控、JMX 和 HTTP代理LLD规则。
	*/
	Password string `json:"password,omitempty"`

	/**
	  HTTP代理LLD规则字段。在post属性中的数据body的类型 。0 - (默认值) 行数据；2 - JSON数据。3 - XML数据。
	*/
	PostType int64 `json:"post_type,omitempty"`

	/**
	  HTTP代理LLD规则字段。HTTP(S) 请求body数据。使用post类型。
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
	  HTTP代理LLD规则字段。查询参数。具有 'key':'value' 对的对象数组，其中值可以是空字符串。
	*/
	QueryFields []interface{}/* TODO */ `json:"query_fields,omitempty"`

	/**
	  HTTP代理LLD规则字段。请求方法的类型。0 - （默认值） GET1 - POST；2 - PUT；3 - HEAD。
	*/
	RequestMethod int64 `json:"request_method,omitempty"`

	/**
	  HTTP代理LLD规则字段。应该存储响应的哪些部分。0 - （默认值） Body；1 - Headers；2 - body和headers均存储。HEAD请求方法值只能为1。
	*/
	RetrieveMode int64 `json:"retrieve_mode,omitempty"`

	/**
	  SNMP OID。
	*/
	SnmpOid string `json:"snmp_oid,omitempty"`

	/**
	  HTTP代理LLD规则字段。公钥文件路径。
	*/
	SslCertFile string `json:"ssl_cert_file,omitempty"`

	/**
	  HTTP代理LLD规则字段。 私钥文件路径。
	*/
	SslKeyFile string `json:"ssl_key_file,omitempty"`

	/**
	  HTTP代理LLD规则字段。秘钥文件的密码。
	*/
	SslKeyPassword string `json:"ssl_key_password,omitempty"`

	/**
	  （只读） LLD规则状态。可用值：0 - （默认值） 支持；1 - 不支持。
	*/
	State int64 `json:"state,omitempty"`

	/**
	  LLD规则状态。可用值：0 - （默认值） 激活LLD规则；1 - 禁用LLD规则。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  HTTP代理LLD规则字段。 用逗号分隔的所需 HTTP 状态代码范围。还支持用户宏作为逗号分隔列表的一部分。例如：200，200-{$M}，{$M}，200-400。
	*/
	StatusCodes string `json:"status_codes,omitempty"`

	/**
	  （只读）父模板 LLD 规则的 ID。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  监控项数据轮询请求超时。用于HTTP代理和LLD规则脚本。支持用户宏或LLD宏。默认值：3s最大值：60s
	*/
	Timeout string `json:"timeout,omitempty"`

	/**
	  允许的主机。 用于trapper监控项或者HTTP代理LLD规则。
	*/
	TrapperHosts string `json:"trapper_hosts,omitempty"`

	/**
	  验证的用户名。用于简单检查、SSH、Telnet、数据库监控、JMX和HTTP代理LLD规则。SSH和Telnet监控项LLD规则需要。
	*/
	Username string `json:"username,omitempty"`

	/**
	  通用唯一标识符，用于将导入的LLD规则链接到现有的LLD规则中。仅用于模板上的LLD规则。如果没有给出，则自动生成。对于更新操作，此字段为只读。
	*/
	Uuid string `json:"uuid,omitempty"`

	/**
	  HTTP代理LLD规则字段。 验证URL中的主机名位于主机证书的公用名字段或使用者备用名字段中。0 - （默认值） 不验证；1 - 验证。
	*/
	VerifyHost int64 `json:"verify_host,omitempty"`

	/**
	  HTTP代理LLD规则字段。验证主机证书是否真实。0 - （默认值） 不验证；1 - 验证。
	*/
	VerifyPeer int64 `json:"verify_peer,omitempty"`
}

type DiscoverRoleCreateParam struct {
	DiscoverRole

	/**
	  LLD规则过滤LLD规则的对象。
	*/
	Filter interface{}/* TODO */ `json:"filter,omitempty"`

	/**
	  LLD规则预处理选项。
	*/
	Preprocessing []interface{}/* TODO */ `json:"preprocessing,omitempty"`

	/**
	  LLD规则lld_macro_path选项。
	*/
	LldMacroPaths []interface{}/* TODO */ `json:"lld_macro_paths,omitempty"`

	/**
	  LLD规则覆盖选项。
	*/
	Overrides []interface{}/* TODO */ `json:"overrides,omitempty"`
}

type DiscoverRoleUpdateParam struct {
	DiscoverRole

	/**
	  LLD规则过滤替换当前过滤器的对象
	*/
	Filter interface{}/* TODO */ `json:"filter,omitempty"`

	/**
	  LLD规则预处理替换当前预处理选项的选项。
	*/
	Preprocessing []interface{}/* TODO */ `json:"preprocessing,omitempty"`

	/**
	  LLD规则lld_macro_path选项。
	*/
	LldMacroPaths []interface{}/* TODO */ `json:"lld_macro_paths,omitempty"`

	/**
	  LLD规则覆盖选项。
	*/
	Overrides []interface{}/* TODO */ `json:"overrides,omitempty"`
}

type DiscoverRoleGetParam struct {
	GetBaseObject

	/**
	  只返回给定 ID 的 LLD 规则。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  只返回属于给定组中主机的 LLD 规则。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  只返回属于给定主机的 LLD 规则。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  如果设置为 true，则仅返回从模板继承的 LLD 规则。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  仅返回使用给定主机接口的 LLD 规则。
	*/
	Interfaceids []string `json:"interfaceids,omitempty"`

	/**
	  如果设置为 true，则仅返回属于受监控主机的已启用 LLD 规则。
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  如果设置为 true，则仅返回属于模板的 LLD 规则。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  只返回属于给定模板的 LLD 规则。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  返回一个 filter 属性，其中包含 LLD 规则使用的过滤器的数据。
	*/
	SelectFilter []string `json:"selectFilter,omitempty"`

	/**
	  返回具有属于 LLD 规则的图形原型的 graphs 属性。支持count。
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  返回具有属于 LLD 规则的主机原型的 hostPrototypes 属性。支持count。
	*/
	SelectHostPrototypes []string `json:"selectHostPrototypes,omitempty"`

	/**
	  返回一个 hosts 属性，其中包含 LLD 规则所属的主机数组。
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回一个 items 属性，其中包含属于 LLD 规则的项目原型。支持 count。
	*/
	SelectItems []string `json:"selectItems,omitempty"`

	/**
	  返回带有属于 LLD 规则的触发器原型的 triggers 属性。支持count。
	*/
	SelectTriggers []string `json:"selectTriggers,omitempty"`

	/**
	  返回一个 lld_macro_paths 属性，其中包含 LLD 宏列表和分配给每个相应宏的值的路径。
	*/
	SelectLLDMacroPaths []string `json:"selectLLDMacroPaths,omitempty"`

	/**
	  返回带有 LLD 规则预处理选项的 preprocessing 属性。它具有以下属性：type - (string) 预处理选项类型：5 - 正则表达式匹配；11 - XML XPath；12 - JSONPath；15 - 不匹配正则表达式；16 - 检查 JSON 中的错误；17 - 检查错误在 XML 中；20 - 丢弃未更改的心跳；23 - Prometheus 到 JSON；24 - CSV 到 JSON；25 - 替换；27 - XML 到 JSON。 params - (string) 预处理选项使用的附加参数。多个参数由 LF (\n) 字符分隔。error_handler - (string) 预处理步骤失败时使用的操作类型：0 - Zabbix 服务器设置错误消息；<br >1 - 丢弃值；2 - 设置自定义值；3 - 设置自定义错误消息。error_handler_params - (string) 错误处理程序参数。
	*/
	SelectPreprocessing []string `json:"selectPreprocessing,omitempty"`

	/**
	  返回一个 lld_rule_overrides 属性，其中包含在原型对象上执行的覆盖过滤器、条件和操作的列表。
	*/
	SelectOverrides []string `json:"selectOverrides,omitempty"`

	/**
	  限制子选择返回的记录数。适用于以下子选择：selctItems；selectGraphs；selectTriggers。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  按给定属性对结果进行排序。可用值：itemid、name、key_、delay、type 和 status。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
