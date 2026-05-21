package main

func main() {

	var investmentAmount = 1000
	var annualInterestRate = 5.5
	var years = 10

	var futureValue = calculateFutureValue(investmentAmount, annualInterestRate, years)
	println("Future Value of Investment:", futureValue)
}
