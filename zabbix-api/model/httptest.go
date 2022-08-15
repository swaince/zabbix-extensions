package model

type HttpTest struct {

	/**
	  (只读) Web场景的ID。
	*/
	Httptestid string `json:"httptestid,omitempty"`

	/**
	  Web场景所属主机的ID。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  Web场景的名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  Web场景将使用的用户代理字符串。默认： Zabbix
	*/
	Agent string `json:"agent,omitempty"`

	/**
	  Web场景将使用的身份认证方法。可用值：0 - (默认) none；1 - 基本的HTTP身份认证；2 - NTLM身份认证。
	*/
	Authentication int64 `json:"authentication,omitempty"`

	/**
	  Web场景的执行间隔。接受秒，带后缀的时间单位和用户宏。默认： 1m。
	*/
	Delay string `json:"delay,omitempty"`

	/**
	  执行请求时将发送的HTTP请求头。
	*/
	Headers HttpHeaderObject `json:"headers,omitempty"`

	/**
	  用于基本的HTTP或NTLM身份认证的密码。
	*/
	HttpPassword string `json:"http_password,omitempty"`

	/**
	  Web场景将使用的代理，如下所示： http://[username[:password]@]proxy.example.com[:port]。
	*/
	HttpProxy string `json:"http_proxy,omitempty"`

	/**
	  用于基本的HTTP或NTLM身份认证的用户名。
	*/
	HttpUser string `json:"http_user,omitempty"`

	/**
	  (只读)下一个Web场景执行的时间。
	*/
	Nextcheck int64 `json:"nextcheck,omitempty"`

	/**
	  Web场景在失败之前尝试执行每个步骤的次数。默认： 1。
	*/
	Retries int64 `json:"retries,omitempty"`

	/**
	  用于客户端身份认证的SSL证书文件的名称(必须是PEM格式)。
	*/
	SslCertFile string `json:"ssl_cert_file,omitempty"`

	/**
	  用于客户端身份认证的SSL私钥文件的名称(必须是PEM格式)。
	*/
	SslKeyFile string `json:"ssl_key_file,omitempty"`

	/**
	  SSL私钥密码。
	*/
	SslKeyPassword string `json:"ssl_key_password,omitempty"`

	/**
	  Web场景是否可用。可用值：0 - (默认) 可用；1 - 不可用。
	*/
	Status int64 `json:"status,omitempty"`

	/**
	  (只读) 父模板Web场景的ID。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  Web场景变量。
	*/
	Variables []*HttpHeaderObject `json:"variables,omitempty"`

	/**
	  是否验证SSL证书里指定的主机名与Web场景中使用的主机名匹配。可能的值：0 - (默认) 跳过主机验证；1 - 验证主机。
	*/
	VerifyHost int64 `json:"verify_host,omitempty"`

	/**
	  是否验证Web服务器的SSL证书。可用值：0 - (默认) 跳过对等验证；1 - 验证对等。
	*/
	VerifyPeer int64 `json:"verify_peer,omitempty"`

	/**
	  (在现有的Web场景上只读)全局唯一标识符，用于将导入的Web场景连接到现有场景。仅用于模板上的web场景。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type HttpTestCreateParam struct {
	HttpTest

	/**
	  Web场景steps。
	*/
	Steps []interface{}/* TODO */ `json:"steps,omitempty"`

	/**
	  Web场景tags。
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type HttpTestUpdateParam struct {
	HttpTest

	/**
	  场景步骤用于替换现有步骤。
	*/
	Steps []interface{}/* TODO */ `json:"steps,omitempty"`

	/**
	  Web场景标签。
	*/
	Tags []*TagObject `json:"tags,omitempty"`
}

type HttpTestGetParam struct {
	GetBaseObject

	/**
	  仅返回属于给定主机组的Web场景。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  仅返回属于给定主机的Web场景。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  仅返回给定ID的Web场景。
	*/
	Httptestids []string `json:"httptestids,omitempty"`

	/**
	  如果设置为true，仅返回从模板继承的Web场景。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  如果设置为true，仅返回属于被监控主机的可用Web场景。
	*/
	Monitored bool `json:"monitored,omitempty"`

	/**
	  如果设置为true，仅返回属于模板的Web场景。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  仅返回属于给定模板的Web场景。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  以Web场景的名称展开宏。
	*/
	ExpandName bool `json:"expandName,omitempty"`

	/**
	  以场景步骤的名称展开宏。
	*/
	ExpandStepName bool `json:"expandStepName,omitempty"`

	/**
	  用于标签搜索的规则。可用值：0 - (默认) And/Or；2 - Or。
	*/
	Evaltype int64 `json:"evaltype,omitempty"`

	/**
	  仅返回给定标签的Web场景。根据标签进行精确匹配，并根据运算符值按标签值进行区分大小写或不区分大小写的搜索。格式：[{"tag": "<tag>", "value": "<value>", "operator": "<operator>"}, ...]。一个空数组返回所有的Web场景。可能的运算符类型：0 - (默认) Like；1 - Equal；2 - Not like；3 - Not equal；4 - Exists；5 - Not exists。
	*/
	Tags []*TagObject `json:"tags,omitempty"`

	/**
	  在hosts属性中，以一个数组的方式返回Web场景所属的主机。
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  在steps属性中返回Web场景步骤。支持count。
	*/
	SelectSteps []string `json:"selectSteps,omitempty"`

	/**
	  在tags属性中返回Web场景标签。
	*/
	SelectTags []string `json:"selectTags,omitempty"`

	/**
	  按照给定属性对结果进行排序。可能的值：httptestid和name。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
