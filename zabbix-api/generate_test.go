package main

import "testing"

func Test_generateGetBaseObject(t *testing.T) {
	t.Run("getBase", func(t *testing.T) {
		generateGetBaseObject()
	})
}

func Test_generateHost(t *testing.T) {
	t.Run("getHost", func(t *testing.T) {
		generateHost()
	})
}

func Test_generateHostGroup(t *testing.T) {
	t.Run("getHostGroup", func(t *testing.T) {
		generateHostGroup()
	})
}

func Test_generateHostInterface(t *testing.T) {
	t.Run("getHostInterface", func(t *testing.T) {
		generateHostInterface()
	})
}

func Test_generateDiscoverRole(t *testing.T) {
	t.Run("getDiscoverRole", func(t *testing.T) {
		generateDiscoverRole()
	})
}

func Test_generateItem(t *testing.T) {
	t.Run("Test_generateItem", func(t *testing.T) {
		generateItem()
	})
}

func Test_generateItemPrototype(t *testing.T) {
	t.Run("getItemPrototype", func(t *testing.T) {
		generateItemPrototype()
	})
}

func Test_generateValueMap(t *testing.T) {
	t.Run("getValueMap", func(t *testing.T) {
		generateValueMap()
	})
}

func Test_generateTemplate(t *testing.T) {
	t.Run("getTemplate", func(t *testing.T) {
		generateTemplate()
	})
}

func Test_generateMacro(t *testing.T) {
	t.Run("getMacro", func(t *testing.T) {
		generateMacro()
	})
}

func Test_generateHttpTest(t *testing.T) {
	t.Run("getHttpTest", func(t *testing.T) {
		generateHttpTest()
	})
}

func Test_generateEvent(t *testing.T) {
	t.Run("getEvent", func(t *testing.T) {
		generateEvent()
	})
}

func Test_generateHistory(t *testing.T) {
	t.Run("getHistory", func(t *testing.T) {
		generateHistory()
	})
}

func Test_generateAlert(t *testing.T) {
	t.Run("getAlert", func(t *testing.T) {
		generateAlert()
	})
}

func Test_generateGraphItem(t *testing.T) {
	t.Run("getGraphItem", func(t *testing.T) {
		generateGraphItem()
	})
}

func Test_generateGraph(t *testing.T) {
	t.Run("getGraph", func(t *testing.T) {
		generateGraph()
	})
}

func Test_generateGraphPrototype(t *testing.T) {
	t.Run("getPrototype", func(t *testing.T) {
		generateGraphPrototype()
	})
}

func Test_generateMediaType(t *testing.T) {
	t.Run("getMediaType", func(t *testing.T) {
		generateMediaType()
	})
}

func Test_generateAuditLog(t *testing.T) {
	t.Run("getAuditLog", func(t *testing.T) {
		generateAuditLog()
	})
}

func Test_generateTrigger(t *testing.T) {
	t.Run("generateTrigger", func(t *testing.T) {
		generateTrigger()
	})
}

func Test_generateTriggerPrototype(t *testing.T) {
	t.Run("generateTriggerPrototype", func(t *testing.T) {
		generateTriggerPrototype()
	})
}

func Test_generateProblem(t *testing.T) {
	t.Run("Test_generateProblem", func(t *testing.T) {
		generateProblem()
	})
}
