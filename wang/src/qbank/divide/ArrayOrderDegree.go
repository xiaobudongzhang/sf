package divide



//求数组中逆序对的个数（O()）
/**
 * 将数组A分成A1 A2两部分，
 * 分别求出A1的逆序度K1，A2的逆序度K2，
 * A1与A2之间的逆序度K3，那么数组A的逆序度就是K1+K2+K3
 */
 var num int
func ArrayOrderDegree(arr []int) int {
	num = 0
	arrayOrderDegree(arr, 0, len(arr) - 1)
	return num
}

func arrayOrderDegree(arr []int, p int, r int)  {
	if p >= r {//终止条件
		return
	}
	q := (p + r) / 2
	arrayOrderDegree(arr, p, q)
	arrayOrderDegree(arr, q + 1, r)
	merge(arr, p, q, r)
}

func merge(arr []int, p int, q int, r int)  {
	i, j, k := p, q+1, 0
	tmpArr := make([]int, r - p + 1)

	for i <= q && j <= r {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			num += q -i + 1
			tmpArr[k] = arr[j]
			j++
		}
		k++
	}

	for i <= q  {
		tmpArr[k] = arr[i]
		k++
		i++
	}

	for j <= r  {
		tmpArr[k] = arr[j]
		k++
		j++
	}
	for i := 0;i < r - p + 1; i++ {
		arr[p+i] = tmpArr[i]
	}
	tmpArr = nil
}
