package dynamic

//import "fmt"

//items 每个物品的重量
//n 物品个数
//w 背包最大承载重量
func Knapsack(items []int, n int, w int) int {
	caches := make([][]bool, n)
	for k, _ := range items {
		caches[k] = make([]bool, w+1)//包含0
	}
	//第一行特殊处理
	caches[0][0] = true
	caches[0][items[0]] = true

	for i:=1;i<n;i++{

		for j:=0;j<=w ;j++  {
			if (caches[i-1][j] == false) {
				continue;
			}
			//不放入背包
			caches[i][j] = true
			//放入背包
			if (j + items[i]> w) {
				continue
			}
			caches[i][j+items[i]] = true
		}
	}
	//fmt.Printf("%v\n", caches)
	last := n -1
	for j:=w;j>=0 ;j--  {
		if caches[last][j] == true {
			return j
		}
	}
	return 0
}
