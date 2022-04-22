package utils
func ReplaceSpace01(s string) string {
	var res string
	for _, v := range s {
		switch v {
		case ' ': //遇到空格则加上%20
			res += "%20"
		default: //默认加上v值，但是要转为string
			res += string(v)
		}
	}
	return res
}