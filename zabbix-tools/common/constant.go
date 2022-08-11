package common

import (
	"fmt"
	"strings"
)

const (
	ZabbixServer = "https://www.zabbix.com/documentation/6.0/zh/manual/api/reference"
)

func GetDocUrl(path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s%s", ZabbixServer, path)
	}
	return fmt.Sprintf("%s/%s", ZabbixServer, path)
}
