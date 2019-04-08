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
//因为不需要保存历史记录,所以可以用一维数组,利用数组中下一级的总量是顺序累计的特点
//items 每个物品的重量
//n 物品个数
//w 背包最大承载重量
func Knapsack2(items []int, n int, w int) int {
	var states []bool

	states = make([]bool, w+1)

	states[0] = true//第一个物品不放入

	for i:=1;i<n ;i++  {
		//不放入的情况
/*		for j:=0;j<w ;j++  {
			if states[j] == true{//不放入的情况下，本个物品的重量跟上个物品的重量相同

			}
		}*/
		//放入的情况
		for j:=0;j+items[i]<w ;j++  {//防止超过最大值，数组越界
			if states[j] == true{
				states[j+items[i]] = true //放入物品的话，在以前的基础上加上本次的物品重量
			}
		}
	}

	//找到最大的重量
	for j:=w;j>=0 ;j--  {//states的j就是重量
		if (states[j] == true){
			return j
		}
	}
	return 0
}


//items 每个物品的重量
//value 物品的价值
//n 物品个数
//w 背包最大承载重量
func Knapsack3(items []int, value []int,n int, w int) int {
	var states [][]int
	//初始化状态
	states = make([][]int, n,n)
	for i:=0;i<n ;i++  {
		statestmp := make([]int, w+1)//保存0的重量
		for j:=0;j<w+1 ;j++  {
			statestmp[j] = -1
		}
		states[i] = statestmp
	}

	states[0][0] = 0 //第一个物品不放入的价值
	states[0][items[0]] = value[0] //第一个物品放入的价值

	for i:=1;i<n ;i++  {
		//不放入的情况
		for j:=0;j<w ;j++  {
			if states[i-1][j] >= 0{//-1的直接过滤掉
				states[i][j] = states[i-1][j]
			}
		}
		//放入的情况
		for j:=0;j+items[i]<w ;j++  {//防止超过最大值，数组越界
			if states[i-1][j] >= 0{//-1的直接过滤掉
				v := states[i-1][j] + value[i]
				//states[i][j+items[i]] 来源于上面不放入的计算出来的值
				//states[i-1][j] + value[i] 来源于放入的值
				if v > states[i][j+items[i]]{
					states[i][j+items[i]] = v //放入物品的话，在以前的基础上加上本次的物品重量
				}
			}
		}
	}

	//找到最大的重量
	maxVal := -1
	for j:=0;j<=w ;j++  {//states的j就是重量
		if (states[n-1][j] > maxVal){
			maxVal = states[n-1][j]
		}
	}
	return maxVal
}