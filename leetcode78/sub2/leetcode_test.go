package leetcode78

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*

732. 我的日程安排表 III

当 k 个日程存在一些非空交集时（即, k 个日程包含了一些相同时间），就会产生 k 次预订。

给你一些日程安排 [startTime, endTime) ，请你在每个日程安排添加后，返回一个整数 k ，表示所有先前日程安排会产生的最大 k 次预订。

实现一个 MyCalendarThree 类来存放你的日程安排，你可以一直添加新的日程安排。

MyCalendarThree() 初始化对象。
int book(int startTime, int endTime) 返回一个整数 k ，表示日历中存在的 k 次预订的最大值。


示例：

输入：
["MyCalendarThree", "book", "book", "book", "book", "book", "book"]
[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
输出：
[null, 1, 1, 2, 3, 3, 3]

解释：
MyCalendarThree myCalendarThree = new MyCalendarThree();
myCalendarThree.book(10, 20); // 返回 1 ，第一个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
myCalendarThree.book(50, 60); // 返回 1 ，第二个日程安排可以预订并且不存在相交，所以最大 k 次预订是 1 次预订。
myCalendarThree.book(10, 40); // 返回 2 ，第三个日程安排 [10, 40) 与第一个日程安排相交，所以最大 k 次预订是 2 次预订。
myCalendarThree.book(5, 15); // 返回 3 ，剩下的日程安排的最大 k 次预订是 3 次预订。
myCalendarThree.book(5, 10); // 返回 3
myCalendarThree.book(25, 55); // 返回 3


提示：

0 <= startTime < endTime <= 109
每个测试用例，调用 book 函数最多不超过 400次
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, []int, int) []int
		act     []string
		params  [][]int
		want    []int
		isDebug bool
	}{
		{
			act:    []string{"MyCalendarThree", "book", "book", "book", "book", "book", "book"},
			params: [][]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}},
			want:   []int{1, 1, 1, 2, 3, 3, 3},
		},
		{
			act:    []string{"MyCalendarThree", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			params: [][]int{{}, {24, 40}, {43, 50}, {27, 43}, {5, 21}, {30, 40}, {14, 29}, {3, 19}, {3, 14}, {25, 39}, {6, 19}},
			want:   []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4},
			// isDebug: true,
		},
		{
			act:     []string{"MyCalendarThree", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			params:  [][]int{{}, {47, 50}, {1, 10}, {27, 36}, {40, 47}, {20, 27}, {15, 23}, {10, 18}, {27, 36}, {17, 25}, {8, 17}, {24, 33}, {23, 28}, {21, 27}, {47, 50}, {14, 21}, {26, 32}, {16, 21}, {2, 7}, {24, 33}, {6, 13}, {44, 50}, {33, 39}, {30, 36}, {6, 15}, {21, 27}, {49, 50}, {38, 45}, {4, 12}, {46, 50}, {13, 21}},
			want:    []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7},
			isDebug: true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var calendar MyCalendarThree
			var ans []int
			for ii, vAct := range v.act {
				switch vAct {
				case "MyCalendarThree":
					calendar = Constructor()
					ans = append(ans, 1)
				case "book":
					res := calendar.Book(v.params[ii][0], v.params[ii][1])
					ans = append(ans, res)
					if v.isDebug {
						fmt.Printf("ii %d res %v \n", ii, res)
						fmt.Printf("v.params %v \n", v.params[ii])
						fmt.Println("===================")
						for _, v := range calendar.layerArea {
							for _, vv := range v {
								fmt.Printf("%+v", *vv)
							}
							fmt.Printf("\n")
						}
						fmt.Println("===================")
					}
				default:
					panic("not support act")
				}

			}
			if !reflect.DeepEqual(ans, v.want) {
				t.Errorf("got %v want %v", ans, v.want)
			}
		})
	}
}

type MyCalendarThree struct {
	layerArea [][]*CalendarArea
}

type CalendarArea struct{ start, end int }

func (t CalendarArea) merge(t2 CalendarArea) (res *CalendarArea, isMerge bool) {
	if t.end > t2.start && t.start < t2.end {
		return &CalendarArea{start: min(t.start, t2.start), end: max(t.end, t2.end)}, true
	}
	return nil, false
}
func (t CalendarArea) isMerge(t2 CalendarArea) (isMerge bool) {
	return t.end > t2.start && t.start < t2.end
}

func (t *MyCalendarThree) pushArea(layer int, area *CalendarArea) {
	if len(t.layerArea) < layer {
		t.layerArea = append(t.layerArea, make([]*CalendarArea, 0))
		t.layerArea[layer-1] = append(t.layerArea[layer-1], area)
		// fmt.Println("layerArea", t.layerArea)
		return
	}
	skipIdx := make([]int, 0)
	newAreaList := make([]*CalendarArea, 0)
	for i, v := range t.layerArea[layer-1] {
		// 存在交叉
		if newArea, isMerge := v.merge(*area); isMerge {
			skipIdx = append(skipIdx, i)
			newAreaList = append(newAreaList, newArea)
		}
	}
	if len(newAreaList) == 0 {
		t.layerArea[layer-1] = append(t.layerArea[layer-1], area)
		return
	}

	//TODO 清空newAreaList 重复部分
	t.layerArea[layer-1] = append(t.layerArea[layer-1], newAreaList...)
	// 删除交叉的区域
	for i := len(skipIdx) - 1; i >= 0; i-- {
		t.layerArea[layer-1] = append(t.layerArea[layer-1][:skipIdx[i]], t.layerArea[layer-1][skipIdx[i]+1:]...)
	}
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		layerArea: make([][]*CalendarArea, 1),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	layer := len(this.layerArea)
	for i := layer - 1; i >= 0; i-- {
		for j := len(this.layerArea[i]) - 1; j >= 0; j-- {
			if this.layerArea[i][j].isMerge(CalendarArea{start: startTime, end: endTime}) {
				// log.Println("add layer", i+1)
				this.pushArea(i+2, &CalendarArea{start: max(this.layerArea[i][j].start, startTime), end: min(this.layerArea[i][j].end, endTime)})
			}
		}
	}

	this.pushArea(1, &CalendarArea{start: startTime, end: endTime})
	return len(this.layerArea)
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */