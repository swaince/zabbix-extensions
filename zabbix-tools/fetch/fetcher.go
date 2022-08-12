package fetch

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/swaince/zabbix-tools/common"
)

func FetchDoc(url string, tableIndex int, excludes ...string) []*FieldObject {

	excludeFields := make([]string, 0)
	if excludes != nil {
		excludeFields = append(excludeFields, excludes...)
	}

	fields := make([]*FieldObject, 0)
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("Accept-Encoding", "identity")
	})

	c.OnHTML(".table-container", func(table *colly.HTMLElement) {
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
			if common.IsChinese(f.ParseKey) {
				fmt.Println(f.ParseKey)
				return
			}
			if len(excludeFields) > 0 {
				for _, field := range excludeFields {
					if f.ParseKey == field {
						return
					}
				}
			}
			// fmt.Printf("rawKey: %s, newKey: %s, rawType: %s, newType: %s, rawDesc: %s\n", f.RawKey, f.NewKey, f.RawType, f.NewType, f.RawDesc)
			fields = append(fields, f)
		})
	})

	c.Visit(url)

	return fields
}
