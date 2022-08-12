package model

type History struct {

	/**
	  历史条目的 ID。
	*/
	Id string `json:"id,omitempty"`

	/**
	  收到值的时间。
	*/
	Clock int64 `json:"clock,omitempty"`

	/**
	  相关监控项的 ID。
	*/
	Itemid string `json:"itemid,omitempty"`

	/**
	  Windows 事件日志条目的 ID。
	*/
	Logeventid int64 `json:"logeventid,omitempty"`

	/**
	  收到值时的纳秒数。
	*/
	Ns int64 `json:"ns,omitempty"`

	/**
	  Windows 事件日志条目的级别。
	*/
	Severity int64 `json:"severity,omitempty"`

	/**
	  Windows 事件日志条目的来源。
	*/
	Source string `json:"source,omitempty"`

	/**
	  Windows 事件日志条目的时间。
	*/
	Timestamp int64 `json:"timestamp,omitempty"`

	/**
	  收到的值。
	*/
	Value string `json:"value,omitempty"`
}

type HistoryGetParam struct {
	GetBaseObject

	/**
	  要返回的历史对象的类型。可能的取值：0 - 浮点数；1 - 字符；2 - 日志；3 - 无符号数；4 - 文本。默认： 3.
	*/
	History int64 `json:"history,omitempty"`

	/**
	  只返回给定主机的历史。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  只返回给定监控项的历史。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  只返回在给定时间或在给定时间之后收到的值。
	*/
	TimeFrom int64 `json:"time_from,omitempty"`

	/**
	  只返回在给定时间之前或在给定时间收到的值。
	*/
	TimeTill int64 `json:"time_till,omitempty"`

	/**
	  按给定的属性对结果进行排序。可能的取值： itemid 和 clock.
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
