/*
 *  @Author : huangzj
 *  @Time : 2020/7/20 10:27
 *  @Description：
 */

package deepSearch

import (
	"fmt"
	"testing"
)

func TestOriginDeepSearch(t *testing.T) {
	srv := NewDeepSearch()
	itemGroup := srv.GetItemMatchNum([]ItemValue{
		{
			ItemId:    1,
			ItemValue: 1,
		}, {
			ItemId:    2,
			ItemValue: 2,
		}, {
			ItemId:    3,
			ItemValue: 3,
		}, {
			ItemId:    4,
			ItemValue: 4,
		}, {
			ItemId:    5,
			ItemValue: 5,
		},
	}, 6)

	for _, r := range itemGroup {
		fmt.Println()
		for _, row := range r {
			fmt.Println(fmt.Sprintf("{%d  %d}", row.ItemId, row.ItemValue))
		}
	}
}

func TestDeepSearch(t *testing.T) {
	srv := NewDeepSearch()
	itemGroup := srv.GetItemMatchValue([]ItemValue{
		{
			ItemId:    1,
			ItemValue: 1,
		}, {
			ItemId:    2,
			ItemValue: 2,
		}, {
			ItemId:    3,
			ItemValue: 3,
		}, {
			ItemId:    4,
			ItemValue: 4,
		}, {
			ItemId:    5,
			ItemValue: 5,
		},
	}, 5)

	for _, r := range itemGroup {
		fmt.Println()
		for _, row := range r {
			fmt.Println(fmt.Sprintf("{%d  %d}", row.ItemId, row.ItemValue))
		}
	}
}

func TestDeepSearchLimitNum(t *testing.T) {
	srv := NewDeepSearch()
	itemGroup := srv.GetItemMatchValueLimitNum([]ItemValue{
		{
			ItemId:    1,
			ItemValue: 1,
		}, {
			ItemId:    2,
			ItemValue: 2,
		}, {
			ItemId:    3,
			ItemValue: 3,
		}, {
			ItemId:    4,
			ItemValue: 4,
		}, {
			ItemId:    5,
			ItemValue: 5,
		},
	}, 5, 4)

	for _, r := range itemGroup {
		fmt.Println()
		for _, row := range r {
			fmt.Println(fmt.Sprintf("{%d  %d}", row.ItemId, row.ItemValue))
		}
	}
}
func TestDeepSearchInSection(t *testing.T) {
	srv := NewDeepSearch()
	itemGroup := srv.GetItemMatchValueInSection([]ItemValue{
		{
			ItemId:    1,
			ItemValue: 1,
		}, {
			ItemId:    2,
			ItemValue: 2,
		}, {
			ItemId:    3,
			ItemValue: 3,
		}, {
			ItemId:    4,
			ItemValue: 4,
		}, {
			ItemId:    5,
			ItemValue: 5,
		},
	}, 5, 6)

	for _, r := range itemGroup {
		fmt.Println()
		for _, row := range r {
			fmt.Println(fmt.Sprintf("{%d  %d}", row.ItemId, row.ItemValue))
		}
	}
}
