package catalan

//for循环

func CataFor(n int) uint64 {
	if n <= 1 {
		return 1
	}
	h := make([]uint64, n+1)
	h[0] = 1
	h[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			//根据递归式计算 h(i)= h(0)*h(i-1)+h(1)*h(i-2) + ... + h(i-1)h(0)
			h[i] += (h[j] * h[i-1-j])
		}
	}
	return h[n]
}
