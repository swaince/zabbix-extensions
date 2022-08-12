package common

import "regexp"

func IsChinese(str string) bool {
	var a = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	for _, v := range str {
		//golang中string的底层是byte类型，所以单纯的for输出中文会出现乱码，这里选择for-range来输出
		if a.MatchString(string(v)) {
			return true
		}
	}

	return false
}
