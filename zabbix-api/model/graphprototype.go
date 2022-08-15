package model

type GraphPrototype struct {

	/**
	  （只读） 图表原型的 ID。
	*/
	Graphid string `json:"graphid,omitempty"`

	/**
	  图表形原型的高度，以像素为单位。
	*/
	Height int64 `json:"height,omitempty"`

	/**
	  图表形原型的名称。
	*/
	Name string `json:"name,omitempty"`

	/**
	  图表形原型的宽度，以像素为单位。
	*/
	Width int64 `json:"width,omitempty"`

	/**
	  图表原型的布局类型。可能的取值：0 - （默认） 普通；1 - 堆积图；2 - 饼图；3 - 分散饼图。
	*/
	Graphtype int64 `json:"graphtype,omitempty"`

	/**
	  左百分比。默认：0.
	*/
	PercentLeft float64 `json:"percent_left,omitempty"`

	/**
	  右百分比。默认：0.
	*/
	PercentRight float64 `json:"percent_right,omitempty"`

	/**
	  是否以3D形式展示饼图和分散饼图。可能的取值：0 - （默认） 以2D展示；1 - 以3D展示。
	*/
	Show3D int64 `json:"show_3d,omitempty"`

	/**
	  是否在已发现的图表上显示图例。可能的取值：0 - 隐藏；1 - （默认） 显示。
	*/
	ShowLegend int64 `json:"show_legend,omitempty"`

	/**
	  是否在已发现的图表上显示工作时间。可能的取值：0 - 隐藏；1 - （默认） 显示。
	*/
	ShowWorkPeriod int64 `json:"show_work_period,omitempty"`

	/**
	  （只读） 父模板图表原型的ID。
	*/
	Templateid string `json:"templateid,omitempty"`

	/**
	  Y轴的固定最大值。
	*/
	Yaxismax float64 `json:"yaxismax,omitempty"`

	/**
	  Y轴的固定最小值。
	*/
	Yaxismin float64 `json:"yaxismin,omitempty"`

	/**
	  用于作为Y轴最大值的监控项的ID.
	*/
	YmaxItemid string `json:"ymax_itemid,omitempty"`

	/**
	  Y轴最大值的计算方式。可能的取值：0 - （默认） 可计算的；1 - 固定的；2 - 监控项。
	*/
	YmaxType int64 `json:"ymax_type,omitempty"`

	/**
	  用于作为Y轴最小值的监控项的ID.
	*/
	YminItemid string `json:"ymin_itemid,omitempty"`

	/**
	  Y轴最小值的计算方式。可能的取值：0 - （默认） 可计算的；1 - 固定的；2 - 监控项。
	*/
	YminType int64 `json:"ymin_type,omitempty"`

	/**
	  图表原型的发现状态。可能的取值：0 - （默认） 将会发现新的图表。1 - 将不会发现新的图表，现有的图表将被标记为丢失。
	*/
	Discover int64 `json:"discover,omitempty"`

	/**
	  唯一通用标识符，用于将导入的图表原型与已有的图表原型联系起来。只用于模板上的图表原型。如果没有给出，则自动生成。 对于更新操作，这个字段是只读的。
	*/
	Uuid string `json:"uuid,omitempty"`
}

type GraphPrototypeCreateParam struct {
	GraphPrototype

	/**
	  要为图表原型创建的图表项目。图形项目可以同时引用项目和项目原型，但至少要有一个项目原型。
	*/
	Gitems []*GraphItem `json:"gitems,omitempty"`
}

type GraphPrototypeUpdateParam struct {
	GraphPrototype

	/**
	  图表项目将替换现有的图表项目。如果一个图表项目定义了 gitemid 属性，它将会被更新，否则会创建一个新的图表项目。
	*/
	Gitems []*GraphItem `json:"gitems,omitempty"`
}

type GraphPrototypeGetParam struct {
	GetBaseObject

	/**
	  只返回属于给定发现规则的图表原型。
	*/
	Discoveryids []string `json:"discoveryids,omitempty"`

	/**
	  只返回具有给定ID的图表原型。
	*/
	Graphids []string `json:"graphids,omitempty"`

	/**
	  只返回属于给定主机组中的主机的图表原型。
	*/
	Groupids []string `json:"groupids,omitempty"`

	/**
	  只返回属于给定主机的图表原型。
	*/
	Hostids []string `json:"hostids,omitempty"`

	/**
	  如果设置为true，只返回从模板继承的图表原型。
	*/
	Inherited bool `json:"inherited,omitempty"`

	/**
	  只返回包含给定监控项原型的图表原型。
	*/
	Itemids []string `json:"itemids,omitempty"`

	/**
	  如果设置为true，只返回属于模板的图表原型。
	*/
	Templated bool `json:"templated,omitempty"`

	/**
	  只返回属于给定模板的图表原型。
	*/
	Templateids []string `json:"templateids,omitempty"`

	/**
	  返回一个发现规则属性，其带有图表原型所属的LLD规则。
	*/
	SelectDiscoveryRule []string `json:"selectDiscoveryRule,omitempty"`

	/**
	  返回一个图表项属性，其中包含图表原型中使用的图表项目。
	*/
	SelectGraphItems []string `json:"selectGraphItems,omitempty"`

	/**
	  返回一个主机组属性，包含图表原型所属的主机组。
	*/
	SelectGroups []string `json:"selectGroups,omitempty"`

	/**
	  返回一个主机属性，包含图表原型所属的主机。
	*/
	SelectHosts []string `json:"selectHosts,omitempty"`

	/**
	  返回一个items属性，包含图表原型中使用的监控项和监控项原型。
	*/
	SelectItems []string `json:"selectItems,omitempty"`

	/**
	  返回一个模板属性，其带有含图表原型所属的模板。
	*/
	SelectTemplates []string `json:"selectTemplates,omitempty"`

	/**
	  按给定的属性对结果进行排序。可能的取值： graphid、 name 和 graphtype。
	*/
	Sortfield []string `json:"sortfield,omitempty"`
}
