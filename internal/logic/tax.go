package logic

func CalculateTax(income float64) float64 {
	// 假设的税率计算逻辑
	var tax float64
	if income <= 5000 {
		tax = 0
	} else if income <= 8000 {
		tax = (income - 5000) * 0.03
	} else if income <= 17000 {
		tax = (income-8000)*0.1 + 90
	} else {
		tax = (income-17000)*0.2 + 1090
	}
	return tax
}