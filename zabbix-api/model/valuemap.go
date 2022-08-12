package model

type ValueMap struct {

	/**
	  (只读) 值映射的ID。
	*/
	Valuemapid string `json:"valuemapid,omitempty"`

	/**
	  值映射的主机ID。
	*/
	Hostid string `json:"hostid,omitempty"`

	/**
	  值映射的名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  当前值映射的值映射关系。映射对象 described in detail below。
	*/
	Mappings []interface{}/* TODO */ `json:"mappings,omitempty"`

	/**
	  通用唯一标识，用于将引入的值映射关联已经存在的对象。仅用于模板上的值映射。如果没有提供，则自动生成。用于更新操作，这个字段是只读。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type ValueMapCreateParam struct {
	ValueMap
}

type ValueMapUpdateParam struct {
	ValueMap
}

type ValueMapGetParam struct {
	GetBaseObject

	/**
	  仅返回给定ID的值映射。
	*/
	Valuemapids []string `json:"valuemapids,omitempty"`

	/**
	  返回 mappings 属性中当前值映射的值映射关系。支持count。
	*/
	SelectMappings map[string][]string `json:"selectMappings,omitempty"`

	/**
	  按照给定的属性对结果进行排序。可用值： valuemapid，name。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}

type ValueMapMappingObject struct {

	/**
	  映射匹配类型。等于 0,1,2,3,4 类型的 value 字段不能为空，类型 5 的value 字段必须为空。可用值：0 - (默认) 精确匹配 ;1 - 如果值大于或者等于，映射将会被应用1；2 - 如果值小于或者等于，映射将会被应用1；3 - 如果值在一个范围（包含范围边界），允许定义用逗号符号分隔的多个范围，映射将会被应用1；4 - 如果值和正真表达式匹配，映射将会被应用2；5 - 默认值，如果没有找到其他的匹配，映射将会被应用。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  原始值。“默认”类型的映射不是必填的。
	*/
	Value string `json:"value,omitempty"`

	/**
	  原始值的映射值。
	*/
	Newvalue string `json:"newvalue,omitempty"`
}
