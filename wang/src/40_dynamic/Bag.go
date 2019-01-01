package _0_dynamic



//items 每个物品的重量
//n 物品个数
//w 背包最大承载重量
func Knapsack(items []int, n int, w int) int {
	var states [][]bool

	states = make([][]bool, n,n)
	for i:=0;i<n ;i++  {
		tmp := make([]bool, w+1)
		states[i] = tmp
	}

	states[0][0] = true//第一个物品不放入
	states[0][items[0]] = true //第一个物品放入

	for i:=1;i<n ;i++  {
		//不放入的情况
		for j:=0;j<w ;j++  {
			if states[i-1][j] == true{//不放入的情况下，本个物品的重量跟上个物品的重量相同
				states[i][j] = true
			}
		}
		//放入的情况
		for j:=0;j+items[i]<w ;j++  {//防止超过最大值，数组越界
			if states[i-1][j] == true{
				states[i][j+items[i]] = true //放入物品的话，在以前的基础上加上本次的物品重量
			}
		}
	}

	//找到最大的重量
	for j:=w;j>=0 ;j--  {//states的j就是重量
		if (states[n-1][j] == true){
			return j
		}
	}
	return 0
}