package math

func Div(a int, b int) (div int) {
	if b != 0 {
		div = a/b
		return div
	} else {
		panic("a problem")
		return 0
	}

}
