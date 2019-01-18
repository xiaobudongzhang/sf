package array

func MaxProfit(prices []int) int {
	//每次都在最低点买入，在最高点卖出
	low := prices[0]
	sum := 0

	for i:=0; i<len(prices) - 1; i++  {

		if prices[i]  < low {
			low = prices[i]
		}

		if prices[i+1] < prices[i] {//顶点，卖出
			sum += prices[i] - low
			low = prices[i+1]
		}

	}

	//最后一位处理
	if prices[len(prices)-1] > low {
		sum += prices[len(prices)-1] - low
	}

	return sum
}
