package calc

import "fmt"

func CalcFatRate(BMI float64, age int, sex string) (fatRate float64, err error) {
	if BMI <= 0 {
		return 0, fmt.Errorf("BMI不能是0或者负数")
	}
	if age <= 0 {
		return 0, fmt.Errorf("age不能是0或者负数")
	}
	if age >= 150 {
		return 0, fmt.Errorf("age不能是大于150岁")
	}
	if sex == "非男非女" {
		return 0, fmt.Errorf("sex不能是非男非女")
	}
	// todo 验证体重的合法性

	return fatRate, nil
}
