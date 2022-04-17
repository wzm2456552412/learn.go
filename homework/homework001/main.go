package main

import "fmt"

func main() {
	var totalFatRate float64
	names := [3]string{}
	weights := [3]float64{}
	talls := [3]float64{}
	ages := [3]int{}
	bmis := [3]float64{}
	fatRates := [3]float64{}

	for i := 0; i < 3; i++ {

		// var name string
		fmt.Print("姓名：")
		fmt.Scanln(&names[i])

		// var weight float64
		fmt.Print("体重（千克）：")
		fmt.Scanln(&weights[i])
		// var tall float64
		fmt.Print("身高(米）：")
		fmt.Scanln(&talls[i])

		bmis[i] = weights[i] / (talls[i] * talls[i])
		// var age int
		fmt.Print("年龄：")
		fmt.Scanln(&ages[i])

		var sexWeight int
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}
		var fatRate float64 = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(sexWeight)) / 100
		fatRates[i] = fatRate
		fmt.Println("体脂率是：", fatRate)

		if sex == "男" {
			// 编写男性的体脂率与体脂状态表
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("目前是：标准，太棒了，要保持。")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("目前是：偏胖，吃完饭多散散步，消化消化。")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("目前是：肥胖，少吃点，多运动")
				} else {
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始")
				}
			}

		} else {
			// 编写女性的体脂率与体脂状态表
			if ages[i] >= 18 && ages[i] <= 39 {
				if fatRate <= 0.2 {
					fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					fmt.Println("目前是：标准，太棒了，要保持。")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println("目前是：偏胖，吃完饭多散散步，消化消化。")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println("目前是：肥胖，少吃点，多运动")
				} else {
					fmt.Println("目前是：非常肥胖，健身游泳跑步，即刻开始")
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		totalFatRate += fatRates[i]
		fmt.Println(names[i], weights[i], talls[i], ages[i], bmis[i], fatRates[i])
	}
	fmt.Print(totalFatRate / 3)
}

func calcBMI(tall float64, weight float64) float64 {
	return tall / (weight * weight)
}
