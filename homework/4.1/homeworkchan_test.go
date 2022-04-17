package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"time"
)

type Rank interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

type Client interface {
	UpdateFR(name string, fr float64)
	GetRank(name string) int
}

var rank Rank

var clients []Client

type PersonalInformation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
}

type inputFromStd struct {
}

func (inputFromStd) GetInput() *PersonalInformation {
	// 录入各项
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高(米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	sex := "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	return &PersonalInformation{
		Name:   name,
		Sex:    sex,
		Tall:   tall,
		Weight: weight,
		Age:    age,
	}
}

func TestChan(t *testing.T) {

	for i := 0; i < len(clients); i++ {
		go func(idx int) {
			// todo add context to control exit
			go func() {
				Ch1 := make(chan int)
				workerCount := 1000
				for i := 0; i < workerCount; i++ {
					go func(i int) {
						fmt.Println(i, "更新体脂：", time.Now())
						for {
							clients[idx].UpdateFR("", 0.23) // 0.23 to be replaced with base +- delta // 完成客户端的更新
						}
						Ch1 <- i
						fmt.Println("结束更新", time.Now())
					}(i)
				}
				fmt.Println("关闭 Ch1")
				close(Ch1)
			}()

			go func() {
				Ch2 := make(chan int)
				workerCount := 1000
				for i := 0; i < workerCount; i++ {
					go func(i int) {
						fmt.Println(i, "获取排行榜", time.Now())
						for {
							clients[idx].GetRank("")
						}
						Ch2 <- i
						fmt.Println("结束获取", time.Now())
					}(i)
				}

				fmt.Println("关闭 Ch2")
				close(Ch2)
			}()

		}(i)
	}
}

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RankItem
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}