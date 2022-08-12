package model

type HostInterface struct {

	/**
	  (只读) 主机接口的可用性。可能的值有:0 - (默认) 未知;1 - 可用;2 - 不可用。
	*/
	Available int64 `json:"available,omitempty"`

	/**
	  接口的附加对象。 如果接口 'type'是SNMP，则必选
	*/
	Details []interface{}/* TODO */ `json:"details,omitempty"`

	/**
	  (只读) 不可用主机接口的下一次轮询时间。
	*/
	DisableUntil int64 `json:"disable_until,omitempty"`

	/**
	  接口使用的DNS名称如果通过IP直接连接，则可以为空。
	*/
	Dns string `json:"dns,omitempty"`

	/**
	  (只读) 主机接口不可用的错误文本。
	*/
	Error string `json:"error,omitempty"`

	/**
	  (只读) 主机接口不可用时间。
	*/
	ErrorsFrom int64 `json:"errors_from,omitempty"`

	/**
	  接口所属的主机ID。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  (只读) 接口的ID。
	*/
	Interfaceid string `json:"interfaceid,omitempty"`

	/**
	  接口使用的IP地址。如果使用DNS域名连接，可以设置为空。
	*/
	Ip string `json:"ip,omitempty"`

	/**
	  该接口是否在主机上用作默认的接口，在主机上只能将某种类型的一个接口设置为默认值。可能的值有:0 - 不是默认;1 - 默认。
	*/
	Main int64 `json:"main,omitempty"`

	/**
	  接口使用的端口号，可以包含用户宏。
	*/
	Port string `json:"port,omitempty"`

	/**
	  接口类型。可能的值有:1 - agent;2 - SNMP;3 - IPMI;4 - JMX。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  是否通过IP进行连接。可能的值有:0 - 使用主机DNS进行连接;1 - 使用主机接口的主机IP进行连接。
	*/
	Useip int64 `json:"useip,omitempty"`
}

type HostInterfaceCreateParam struct {
	HostInterface
}

type HostInterfaceUpdateParam struct {
	HostInterface
}

type HostInterfaceGetParam struct {
	GetBaseObject

	/**
	  返回给定主机使用的主机接口。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  返回给定ID的主机接口。
	*/
	Interfaceids []string `json:"interfaceids,omitempty"`

	/**
	  返回给定监控项使用的主机接口。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  返回给定的触发器中监控项使用的主机接口。
	*/
	Triggerids []string `json:"triggerids,omitempty"`

	/**
	  返回监控项 属性。其中包括使用该接口的监控项支持 count.
	*/
	SelectItems map[string][]string `json:"selectItems,omitempty"`

	/**
	  返回 主机属性，其中包括使用该接口的主机数组。
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  限制子选择返回的记录数适用于以下子选择:selectItems。
	*/
	LimitSelects int64 `json:"limitSelects,omitempty"`

	/**
	  对给定属性的结果进行排序。可能的值有: interfaceid, dns, ip。
	*/
	Sortfield []string `json:"sortfield,omitempty"`

	/**

	 */
	Nodeids []string `json:"nodeids,omitempty"`
}
