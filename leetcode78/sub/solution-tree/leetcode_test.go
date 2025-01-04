package leetcode78

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"tree/learn/solution-tree/redblacktree"
)

/*
731. 我的日程安排表 II

实现一个程序来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。

当三个日程安排有一些时间上的交叉时（例如三个日程安排都在同一时间内），就会产生 三重预订。

事件能够用一对整数 startTime 和 endTime 表示，在一个半开区间的时间 [startTime, endTime) 上预定。实数 x 的范围为  startTime <= x < endTime。

实现 MyCalendarTwo 类：

MyCalendarTwo() 初始化日历对象。
boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致三重预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。


示例 1：

输入：
["MyCalendarTwo", "book", "book", "book", "book", "book", "book"]
[[], [10, 20], [50, 60], [10, 40], [5, 15], [5, 10], [25, 55]]
输出：
[null, true, true, true, false, true, true]

解释：
MyCalendarTwo myCalendarTwo = new MyCalendarTwo();
myCalendarTwo.book(10, 20); // 返回 True，能够预定该日程。
myCalendarTwo.book(50, 60); // 返回 True，能够预定该日程。
myCalendarTwo.book(10, 40); // 返回 True，该日程能够被重复预定。
myCalendarTwo.book(5, 15);  // 返回 False，该日程导致了三重预定，所以不能预定。
myCalendarTwo.book(5, 10); // 返回 True，能够预定该日程，因为它不使用已经双重预订的时间 10。
myCalendarTwo.book(25, 55); // 返回 True，能够预定该日程，因为时间段 [25, 40) 将被第三个日程重复预定，时间段 [40, 50) 将被单独预定，而时间段 [50, 55) 将被第二个日程重复预定。


提示：

0 <= start < end <= 109
最多调用 book 1000 次。
*/

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f       func([]int, []int, int) []int
		act     []string
		params  [][]int
		want    []bool
		isDebug bool
	}{
		{
			act:     []string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book"},
			params:  [][]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}},
			want:    []bool{true, true, true, true, false, true, true},
			isDebug: false,
		},
		{
			act:     []string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			params:  [][]int{{}, {47, 50}, {1, 10}, {27, 36}, {40, 47}, {20, 27}, {15, 23}, {10, 18}, {27, 36}, {17, 25}, {8, 17}, {24, 33}, {23, 28}, {21, 27}, {47, 50}, {14, 21}, {26, 32}, {16, 21}, {2, 7}, {24, 33}, {6, 13}, {44, 50}, {33, 39}, {30, 36}, {6, 15}, {21, 27}, {49, 50}, {38, 45}, {4, 12}, {46, 50}, {13, 21}},
			want:    []bool{true, true, true, true, true, true, true, true, true, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, true, false, false, false},
			isDebug: true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var calendar MyCalendarTwo
			var ans []bool
			for ii, vAct := range v.act {
				switch vAct {
				case "MyCalendarTwo":
					calendar = Constructor()
					ans = append(ans, true)
				case "book":
					res := calendar.Book(v.params[ii][0], v.params[ii][1])
					ans = append(ans, res)
					if v.isDebug {
						fmt.Printf("ii %d res %v \n", ii, res)
						fmt.Printf("v.params %v \n", v.params[ii])
						fmt.Println("===================")
						for _, v := range calendar.booked {
							fmt.Printf("%+v\n", v)
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

type MyCalendarTwo struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{redblacktree.NewWithIntComparator()}
}

func (c MyCalendarTwo) add(key, value int) {
	if v, ok := c.Get(key); ok {
		c.Put(key, v.(int)+value)
	} else {
		c.Put(key, value)
	}
}

func (c MyCalendarTwo) Book(start, end int) bool {
	c.add(start, 1)
	c.add(end, -1)
	maxBook := 0
	it := c.Iterator()
	for it.Next() {
		maxBook += it.Value().(int)
		if maxBook > 2 {
			c.add(start, -1)
			c.add(end, 1)
			return false
		}
	}
	return true
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/my-calendar-ii/solutions/1678660/wo-de-ri-cheng-an-pai-biao-ii-by-leetcod-wo6n/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
