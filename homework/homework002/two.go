package calc

import (
	"fmt"
	"testing"
)

func TestCalcBMI(t *testing.T) {
	inputHeight, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("开始计算,输入：height: %f,期望结果：%f", inputHeight, inputWeight, expectedOutput)
	actualOutput, err := CalcBMI(inputHeight, inputWeight)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	if tall <= 0 {
		return 0, fmt.Errorf("身高不能是0或者负数")
	}
	// todo 验证体重的合法性
	if weight <= 0 {
		return 0, fmt.Errorf("体重不能是0或者负数")
	}
	return weight / (tall * tall), nil
}
