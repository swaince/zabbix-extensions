package fetch

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/swaince/zabbix-tools/common"
)

func FetchDoc(path string, tableIndex int) *ClassObject {

	fields := make([]*FieldObject, 0)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Accept-Encoding", "identity")
	})
	// #dokuwiki__content > div > div.page.group > div:nth-child(6) > table
	// #dokuwiki__content > div > div.page.group > div:nth-child(9) > table
	c.OnHTML("#dokuwiki__content .table-container", func(table *colly.HTMLElement) {
		if table.Index != tableIndex {
			return
		}

		table.ForEach("table > tbody > tr", func(i int, tr *colly.HTMLElement) {
			f := &FieldObject{}
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				text := td.Text
				if i == 0 {
					f.RawKey = text
				} else if i == 1 {
					f.RawType = text
				} else if i == 2 {
					f.RawDesc = text
				}
			})

			f.Parse()
			fmt.Printf("rawkey: %s, newKey: %s, rawType: %s, newType: %s, rawDesc: %s\n", f.RawKey, f.NewKey, f.RawType, f.NewType, f.RawDesc)
			fields = append(fields, f)
		})
	})

	c.Visit(common.GetDocUrl(path))

	class := &ClassObject{
		Fields: fields,
	}

	return class
}
