package modify

func Modifystring(chain string) string {
	result := ""
	for _, char := range chain {
		result += "#" + string(char)
	}
	result += "#"
	return result
}
