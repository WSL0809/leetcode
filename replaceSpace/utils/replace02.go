package utils
func ReplaceSpace02(s string) string{
	
	var r = []rune(s)
	res := make([]rune, 0)
	for i := 0; i < len(r); i++{
		if r[i] == ' '{
			res = append(res, []rune("%20")...)
		}else{
			res = append(res, r[i])
		}
	}
	return string(res)

}