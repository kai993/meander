package meander

import (
	"fmt"
	"strings"
)

// Cost : 価格帯
type Cost int8

const (
	_ Cost = iota
	// Cost1 : 最安値
	Cost1
	// Cost2 : 安値
	Cost2
	// Cost3 : 普通
	Cost3
	// Cost4 : 高値
	Cost4
	// Cost5 : 最高値
	Cost5
)

var costSetting = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (c Cost) String() string {
	for s, cost := range costSetting {
		if c == cost {
			return s
		}
	}

	return "Invalid cost value"
}

// CostRange : 価格帯の範囲
type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return fmt.Sprintf("%s...%s", r.From.String(), r.To.String())
}

// ParseCost : Costを文字列に変換する
func ParseCost(s string) Cost {
	return costSetting[s]
}

// ParseCostRange : 価格帯の範囲を設定する
func ParseCostRange(s string) *CostRange {
	r := strings.Split(s, "...")
	return &CostRange{
		From: ParseCost(r[0]),
		To:   ParseCost(r[1]),
	}
}
