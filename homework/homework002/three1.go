package calc

import "fmt"

func CalcFatRate1(BMI float64, age int, sex string) (fatRate float64, err error) {
	if BMI == 0.3 {
		return 0.3, fmt.Errorf("BMI", BMI)
	}
	if age == 35 {
		return 35, fmt.Errorf("age", age)
	}
	if sex == "男" {
		return 1, fmt.Errorf("sex", sex)
	}
	// todo 验证体重的合法性

	return fatRate, nil
}
