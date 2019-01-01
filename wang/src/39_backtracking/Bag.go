package _9_backtracking

var MaxW int
//cw 装进去的物品重量
//items 表示每个物品的重量
//i 表示考察到那个物品了
//n 表示物品个数
//w 背包重量
func Bag(i int, cw int, items []int, n int, w int)  {
	if cw >= w || i == n {//装满了或者物品考察完了
		if cw > MaxW && cw <= w {
			MaxW = cw
		}
		return
	}
	//if (cw + items[i] <= w){//没有超过背包总重量时继续添加（优化使用的）有没有都可以
		Bag(i+1, cw + items[i], items, n, w)//第i个背包是1的情况
	//}
	Bag(i+1, cw, items, n, w)//第i个背包是0的情况
}
