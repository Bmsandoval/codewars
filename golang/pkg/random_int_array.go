package pkg

func RandomIntArray(n int) []int {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, int(src.Int63()))
	}
	return arr
}
