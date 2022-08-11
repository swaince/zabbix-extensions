package template

import (
	"github.com/swaince/zabbix-tools/fetch"
	"testing"
	"time"
)

func TestRender(t *testing.T) {
	c := fetch.FetchDoc("host/get", 0)
	c.StructName = "Host"

	time.Sleep(5 * time.Second)
	Render(c)

}
