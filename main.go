package main

// maxProfit returns the maximum profit possible for the stock
func maxProfit(stockPricesByDay []int) int {
	var profit int
	var grandprofit int
	var key int
	grandprofit = 0

	for outerloop := 0; outerloop < len(stockPricesByDay); outerloop++ {
		if outerloop != key {
			continue
		}
		for innerloop := outerloop + 1; innerloop < len(stockPricesByDay); innerloop++ {

			if stockPricesByDay[innerloop] < stockPricesByDay[outerloop] {
				key = innerloop
				break
			}
			if profit < stockPricesByDay[innerloop]-stockPricesByDay[outerloop] {
				profit = stockPricesByDay[innerloop] - stockPricesByDay[outerloop]
			}

		}
		if profit > grandprofit {
			grandprofit = profit
		}
	}
	return grandprofit
}
