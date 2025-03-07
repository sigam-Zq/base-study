package leetcode78

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
729. 我的日程安排表 I

实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。

当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。

日程可以用一对整数 startTime 和 endTime 表示，这里的时间是半开区间，即 [startTime, endTime), 实数 x 的范围为，  startTime <= x < endTime 。

实现 MyCalendar 类：

MyCalendar() 初始化日历对象。
boolean book(int startTime, int endTime) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false 并且不要将该日程安排添加到日历中。


示例：

输入：
["MyCalendar", "book", "book", "book"]
[[], [10, 20], [15, 25], [20, 30]]
输出：
[null, true, false, true]

解释：
MyCalendar myCalendar = new MyCalendar();
myCalendar.book(10, 20); // return True
myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15 已经被另一个日程安排预订了。
myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于 20 ，且不包含时间 20 。


提示：

0 <= start < end <= 109
每个测试用例，调用 book 方法的次数最多不超过 1000 次。
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
			act:    []string{"MyCalendar", "book", "book", "book"},
			params: [][]int{{}, {10, 20}, {15, 25}, {20, 30}},
			want:   []bool{true, true, false, true},
		},
		{
			act:    []string{"MyCalendar", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			params: [][]int{{}, {47, 50}, {33, 41}, {39, 45}, {33, 42}, {25, 32}, {26, 35}, {19, 25}, {3, 8}, {8, 13}, {18, 27}},
			want:   []bool{true, true, true, false, false, true, false, true, true, true, false},
		},
		{
			act:    []string{"MyCalendar", "book", "book", "book", "book", "book"},
			params: [][]int{{}, {37, 50}, {33, 50}, {4, 17}, {35, 48}, {8, 25}},
			want:   []bool{true, true, false, true, false, false},
		},
		{
			act:     []string{"MyCalendar", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			params:  [][]int{{}, {97, 100}, {33, 51}, {89, 100}, {83, 100}, {75, 92}, {76, 95}, {19, 30}, {53, 63}, {8, 23}, {18, 37}, {87, 100}, {83, 100}, {54, 67}, {35, 48}, {58, 75}, {70, 89}, {13, 32}, {44, 63}, {51, 62}, {2, 15}},
			want:    []bool{true, true, true, false, false, true, false, true, true, false, false, false, false, false, false, false, false, false, false, false, true},
			isDebug: true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var calendar MyCalendar
			var ans []bool
			for ii, vAct := range v.act {
				switch vAct {
				case "MyCalendar":
					calendar = Constructor()
					ans = append(ans, true)
				case "book":
					res := calendar.Book(v.params[ii][0], v.params[ii][1])
					ans = append(ans, res)
					if v.isDebug {
						fmt.Printf("ii %d res %v \n", ii, res)
						fmt.Printf("v.params %v \n", v.params[ii])
						fmt.Println("===================")
						for _, v := range calendar {
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

type pair struct{ start, end int }
type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	for _, p := range *c {
		if p.start < end && start < p.end {
			return false
		}
	}
	*c = append(*c, pair{start, end})
	return true
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/my-calendar-i/solutions/1643942/wo-de-ri-cheng-an-pai-biao-i-by-leetcode-nlxr/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
