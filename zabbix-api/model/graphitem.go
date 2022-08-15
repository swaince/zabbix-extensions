package model

type GraphItem struct {

	/**
	  （只读） 图表监控项ID。
	*/
	Gitemid string `json:"gitemid,omitempty"`

	/**
	  图表监控项的绘制颜色是十六进制的颜色代码。
	*/
	Color string `json:"color,omitempty"`

	/**
	  项目ID。
	*/
	Itemid string `json:"itemid,omitempty"`

	/**
	  监控项的值将被展示。可用值：1 - 最小值；2 - （默认） 平均值；4 - 最大值；7 - 所有值；9 - 最后一个值，仅用于饼图和分散饼图。
	*/
	CalcFnc int64 `json:"calc_fnc,omitempty"`

	/**
	  图表监控项的绘制样式。可用值：0 - （默认） 折线图；1 - 填充图；2 - 加粗折线图；3 - 散点图；4 - 虚线折线图；5 - 梯度直方图。
	*/
	Drawtype int64 `json:"drawtype,omitempty"`

	/**
	  图表监控项所属图表的ID。
	*/
	Graphid string `json:"graphid,omitempty"`

	/**
	  该监控项在图表中的位置。默认值：从0开始，每增加一个条目就增加一。
	*/
	Sortorder int64 `json:"sortorder,omitempty"`

	/**
	  图表监控项类型。可用值：0 - （默认） 简单；2 - 图表的总和，仅用于饼图和离散图。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  该图表项目的Y轴将被绘制在图表的一侧。可用值：0 - （默认） 左侧；1 - 右侧。
	*/
	Yaxisside int64 `json:"yaxisside,omitempty"`
}

type GraphItemGetParam struct {
	GetBaseObject

	/**
	  只返回属于给定图表的图表监控项。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回具有给定监控项ID的图表监控项。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  只返回具有给定类型的图表监控项。请参考图表监控项对象页，了解支持的图表监控项类型列表。
	*/
	Type int64 `json:"type,omitempty"`

	/**
	  返回一个图表属性，其中包含该监控项所属的图表的数组。
	*/
	SelectGraphs []string `json:"selectGraphs,omitempty"`

	/**
	  按给定的属性对结果进行排序。可用值： gitemid。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
