package textfmt


func TabText(text string, requiredSize int) string {
	l := len(text)
	if l >= requiredSize {
		return text
	}
	return MultiString(" ", requiredSize-l) + text
}

func MultiString(str string, multi int) string {
	var res string
	for i := 0; i < multi; i++ {
		res += str
	}
	return res
}
