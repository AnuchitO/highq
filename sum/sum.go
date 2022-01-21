package sum

func Sum(xs ...int) int {
	r := 0
	for _, v := range xs {
		r += v
	}
	return r
}