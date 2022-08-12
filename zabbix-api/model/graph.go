package model

type Graph struct {

	/**
	  （只读） 图表的 ID。
	*/
	Graphid string `json:"graphid,omitempty"`

	/**
	  图表的高度，以像素为单位。
	*/
	Height int64 `json:"height,omitempty"`

	/**
	  图表的名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  图表的宽度，以像素为单位。
	*/
	Width int64 `json:"width,omitempty"`

	/**
	  (只读) 图表的起源。可用值：0 - （默认） 普通图表；4 - 发现的图表。
	*/
	Flags int64 `json:"flags,omitempty"`

	/**
	  图表的布局类型。可用值：0 - （默认） 常规；1 - 堆积图；2 - 饼图；3 - 分散饼图。
	*/
	Graphtype int64 `json:"graphtype,omitempty"`

	/**
	  左百分比。默认：0。
	*/
	PercentLeft float64 `json:"percent_left,omitempty"`

	/**
	  右百分比。默认：0。
	*/
	PercentRight float64 `json:"percent_right,omitempty"`

	/**
	  是否以3D形式展示饼图和分散饼图。可用值：0 - （默认） 以2D展示；1 - 以3D展示。
	*/
	Show3D int64 `json:"show_3d,omitempty"`

	/**
	  是否在图表上显示图例。可用值：0 - 隐藏；1 - （默认） 显示。
	*/
	ShowLegend int64 `json:"show_legend,omitempty"`

	/**
	  是否在图表上显示工作时间。可用值：0 - 隐藏；1 - （默认） 显示。
	*/
	ShowWorkPeriod int64 `json:"show_work_period,omitempty"`

	/**
	  是否在图表上显示触发线。可用值：0 - 隐藏；1 - （默认） 显示。
	*/
	ShowTriggers int64 `json:"show_triggers,omitempty"`

	/**
	  （只读） 父模板图表的ID。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  Y轴的固定最大值。默认：100。
	*/
	Yaxismax float64 `json:"yaxismax,omitempty"`

	/**
	  Y轴的固定最小值。默认：0。
	*/
	Yaxismin float64 `json:"yaxismin,omitempty"`

	/**
	  用于作为Y轴最大值的监控项的ID。
	*/
	YmaxItemid string `json:"ymax_itemid,omitempty"`

	/**
	  Y轴最大值的计算方式。可用值：0 - （默认） 可计算的；1 - 固定的；2 - 监控项。
	*/
	YmaxType int64 `json:"ymax_type,omitempty"`

	/**
	  用于作为Y轴最小值的监控项的ID。
	*/
	YminItemid string `json:"ymin_itemid,omitempty"`

	/**
	  Y轴最小值的计算方式。可用值：0 - （默认） 可计算的；1 - 固定的；2 - 监控项。
	*/
	YminType int64 `json:"ymin_type,omitempty"`

	/**
	  唯一通用标识符，用于将导入的图表与已有的图表联系起来。只用于模板上的图形。如果没有给出，则自动生成。 对于更新操作，这个字段是只读的。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type GraphCreateParam struct {
	Graph

	/**
	  要为图表创建的图表监控项。
	*/
	Gitems []interface{}/* TODO */ `json:"gitems,omitempty"`
}

type GraphUpdateParam struct {
	Graph

	/**
	  图表监控项来替换现有的图表监控项。如果一个图表监控项有定义gitemid属性，它将被更新，否则将创建一个新的图表监控项。
	*/
	Gitems []interface{}/* TODO */ `json:"gitems,omitempty"`
}

type GraphGetParam struct {
	GetBaseObject

	/**
	  只返回具有给定ID的图表。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回属于给定主机组中的主机的图表。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  只返回属于给定模板的图表。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  只返回属于给定主机的图表。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  只返回包含给定监控项的图表。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  如果设置为true，只返回属于模板的图表。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  如果设置为true，只返回从模板继承的图表。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  在图表名称中扩展宏。
	*/
	ExpandName bool `json:"expandName,omitempty"`

	/**
	  返回一个组属性，其中包含该图表所属的主机组。
	*/
	SelectGroups map[string][]string `json:"selectGroups,omitempty"`

	/**
	  返回一个模板属性，其中包含该图表所属的模板。
	*/
	SelectTemplates map[string][]string `json:"selectTemplates,omitempty"`

	/**
	  返回一个主机属性，其中包含该图表所属的主机。
	*/
	SelectHosts map[string][]string `json:"selectHosts,omitempty"`

	/**
	  返回一个监控项属性，其中包含该图表中使用的监控项。
	*/
	SelectItems map[string][]string `json:"selectItems,omitempty"`

	/**
	  返回一个 graphDiscovery 属性和图表发现对象。图表发现对象将图表和创建图标的原型联系起来。它具有以下属性：graphid - (string) 图表的ID；parent_graphid - (string) 创建图表的图表原型的ID。
	*/
	SelectGraphDiscovery map[string][]string `json:"selectGraphDiscovery,omitempty"`

	/**
	  返回一个图表监控项属性，其中包含该图表中使用的监控项。
	*/
	SelectGraphItems map[string][]string `json:"selectGraphItems,omitempty"`

	/**
	  返回一个发现规则属性，其中包含创建该图表的低级发现规则。
	*/
	SelectDiscoveryRule map[string][]string `json:"selectDiscoveryRule,omitempty"`

	/**
	  按给定的属性对结果进行排序。可用值： graphid, name 和 graphtype。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
