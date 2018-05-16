package hoge

func Hoge1(x int) int {
	return (x + 1)
}

// can not call from other packages
func hoge1(x int) int {
	return (x + 1)
}
