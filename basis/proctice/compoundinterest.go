package main

import "fmt"

const yearDay = 365

// getAllMoneyScheduledInvestment 代表定投，参数 interestRate 代表年复利
func getAllMoneyScheduledInvestment(principal float64, interestRate float64, cycle int, cycleMoney float64, duration int) (float64, float64) {
	// 定期投资，按日计算收益，利息计入新的资金
	allMoney := principal
	nowCycle := cycle
	for ; duration > 0; duration-- {
		allMoney += allMoney * interestRate * float64(1) / yearDay
		if nowCycle == 0 {
			principal += cycleMoney
			allMoney += cycleMoney
			nowCycle = cycle - 1
		} else {
			nowCycle--
		}
	}
	return principal, allMoney
}

func getAllMoney(principal float64, interestRate float64, duration int) float64 {
	allMoney := principal*interestRate*float64(duration)/yearDay + principal
	return allMoney
}

func getInterest(principal float64, allmoney float64) float64 {
	return allmoney - principal
}

func interestTest() {
	principal := 10000.0
	interestRate := 0.1
	duration := 365 * 30
	allMoney := getAllMoney(principal, interestRate, duration)
	fmt.Printf(`
方案为标准活期投资: 
  本金为 %.2f 元，年利率为 %.2f %%，存放时间为 %d 天，
  利息为 %.2f 元，本息合计为 %.2f 元

`, principal, interestRate*100, duration, getInterest(principal, allMoney), allMoney)
}

func interestScheduledInvestmentTest() {
	allMoneyScheduledInvestmentMessage := `
方案为定期投资：
  原始本金为 %.2f 元，总本金为 %.2f 元，年利率为 %.2f %%，存放时间为 %d 天，
  定投周期为 %d 天，定投金额为 %.2f 元，
  利息为 %.2f 元，本息合计为 %.2f 元
`
	principal := 10000.0
	interestRate := 0.1
	duration := 365 * 30
	cycle := 30
	cycleMoney := 1000.0
	allprincipal, allMoney := getAllMoneyScheduledInvestment(principal, interestRate, cycle, cycleMoney, duration)
	fmt.Printf(allMoneyScheduledInvestmentMessage, principal, allprincipal, interestRate*100, duration, cycle, cycleMoney, getInterest(allprincipal, allMoney), allMoney)
}

//方案为定期投资：
//   原始本金为 10000.00 元，总本金为 131000.00 元，年利率为 10.00 %，存放时间为 3650 天，
//   定投周期为 30 天，定投金额为 1000.00 元，
//   利息为 103639.69 元，本息合计为 234639.69 元

// 方案为定期投资：
//   原始本金为 10000.00 元，总本金为 374000.00 元，年利率为 10.00 %，存放时间为 10950 天，
//   定投周期为 30 天，定投金额为 1000.00 元，
//   利息为 2137002.05 元，本息合计为 2511002.05 元

// 方案为定期投资：
//   原始本金为 10000.00 元，总本金为 1226000.00 元，年利率为 10.00 %，存放时间为 36500 天，
//   定投周期为 30 天，定投金额为 1000.00 元，
//   利息为 2883487538.70 元，本息合计为 2884713538.70 元

// 利滚利真可怕
