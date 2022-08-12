package model

type Macro struct {

	/**
	  (readonly) 主机宏ID.
	*/
	Hostmacroid string `json:"hostmacroid,omitempty"`

	/**
	  宏所属主机的主机ID.
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  宏名.
	*/
	Macro string `json:"macro,omitempty"`

	/**
	  宏值.
	*/
	Value string `json:"value,omitempty"`

	/**
	  宏的类型.可接受的值:0 - (default) 文本宏;1 - 密文宏;2 - 密钥宏.
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  宏的描述信息.
	*/
	Description string `json:"description,omitempty"`

	/**
	  (readonly) 全局宏的ID.
	*/
	Globalmacroid string `json:"globalmacroid,omitempty"`
}
