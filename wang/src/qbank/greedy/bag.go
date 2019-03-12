package greedy

//i 表示考察到那个物品了
//cw 装进去的物品重量
//items 表示每个物品的重量
//n 表示物品个数
//w 背包重量
var MaxW int
func FindMax(i int, cw int, items []int, n int, w int)  {
	//顺序不可以乱掉
	if cw > w {
		return
	}

	if (cw > MaxW) {
		MaxW = cw
	}

	if i == n {
		return
	}

	FindMax(i+1, cw, items, n, w)
	FindMax(i+1, cw + items[i], items, n, w)
}
