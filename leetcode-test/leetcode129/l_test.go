package leetcode125

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

// 3635. 最早完成陆地和水上游乐设施的时间 II
// https://leetcode.cn/problems/earliest-finish-time-for-land-and-water-rides-ii/description/?envType=daily-question&envId=2026-06-03
func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f              func([]int, []int, []int, []int) int
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		want           int
	}{
		{
			f:              earliestFinishTime,
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			want:           9,
		},
		{
			f:              earliestFinishTime,
			landStartTime:  []int{5},
			landDuration:   []int{3},
			waterStartTime: []int{1},
			waterDuration:  []int{10},
			want:           14,
		},
		{
			f:              earliestFinishTime,
			landStartTime:  []int{82, 14},
			landDuration:   []int{42, 30},
			waterStartTime: []int{6, 54},
			waterDuration:  []int{91, 71},
			want:           125,
		},
		{
			f:              earliestFinishTime,
			landStartTime:  []int{80, 71},
			landDuration:   []int{27, 47},
			waterStartTime: []int{45, 3, 64, 7, 25, 45, 6},
			waterDuration:  []int{40, 39, 66, 64, 24, 74, 49},
			want:           107,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.landStartTime, v.landDuration, v.waterStartTime, v.waterDuration)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	landFinalBestSortTimeIdx, waterFinalBestSortTimeIdx := 0, 0
	landFinalBestSortTimeVal, waterFinalBestSortTimeVal := math.MaxInt, math.MaxInt
	minRes := math.MaxInt
	// 最早可以结束的陆上设施
	for i := range landStartTime {
		if landFinalBestSortTimeVal > landStartTime[i]+landDuration[i] {
			landFinalBestSortTimeVal = landStartTime[i] + landDuration[i]
			landFinalBestSortTimeIdx = i
		} else if landFinalBestSortTimeVal == landStartTime[i]+landDuration[i] {
			if landStartTime[landFinalBestSortTimeIdx] < landStartTime[i] {
				landFinalBestSortTimeVal = landStartTime[i] + landDuration[i]
				landFinalBestSortTimeIdx = i
			}
		}
	}

	// 最早可以结束的水上设施
	for i := range waterStartTime {
		if waterFinalBestSortTimeVal > waterStartTime[i]+waterDuration[i] {
			waterFinalBestSortTimeVal = waterStartTime[i] + waterDuration[i]
			waterFinalBestSortTimeIdx = i
		} else if waterFinalBestSortTimeVal == waterStartTime[i]+waterDuration[i] {
			if waterStartTime[waterFinalBestSortTimeIdx] < waterStartTime[i] {
				waterFinalBestSortTimeVal = waterStartTime[i] + waterDuration[i]
				waterFinalBestSortTimeIdx = i
			}
		}

		// 判断是否可以结合当前的水上活动
		// 结束了上 d个项目P 下个项目已经开始的情况
		if landFinalBestSortTimeVal >= waterStartTime[i] {
			if v := landFinalBestSortTimeVal; minRes > v && minRes > v+waterDuration[i] {
				minRes = v + waterDuration[i]
			}
			// 可以无缝把路上项目塞到水上项目前的情况
		} else if landFinalBestSortTimeVal <= waterStartTime[i] {
			minRes = min(minRes, waterStartTime[i]+waterDuration[i])
		}

	}

	// 最早结束的上一个水上项目.前后插一个路上的项目
	for i := range landStartTime {
		if waterFinalBestSortTimeVal >= landStartTime[i] {
			if v := waterFinalBestSortTimeVal; minRes > v && minRes > v+landDuration[i] {
				minRes = v + landDuration[i]
			}
			// 可以无缝把路上项目塞到水上项目的情况
		} else if waterFinalBestSortTimeVal <= landStartTime[i] {
			minRes = min(minRes, landStartTime[i]+landDuration[i])
		}

	}

	return minRes
}

// 局部最优导致错过全局最优. 不是 这里的确可以单个的局部最优去匹配便利找到第二个全局最优

func earliestFinishTimeLeetcode(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	solve := func(start1, duration1, start2, duration2 []int) int {
		finish1 := 2147483647
		for i := 0; i < len(start1); i++ {
			if val := start1[i] + duration1[i]; val < finish1 {
				finish1 = val
			}
		}
		finish2 := 2147483647
		for i := 0; i < len(start2); i++ {
			curStart := start2[i]
			if finish1 > curStart {
				curStart = finish1
			}
			if val := curStart + duration2[i]; val < finish2 {
				finish2 = val
			}
		}
		return finish2
	}

	land_water := solve(landStartTime, landDuration, waterStartTime, waterDuration)
	water_land := solve(waterStartTime, waterDuration, landStartTime, landDuration)
	if land_water < water_land {
		return land_water
	}
	return water_land
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/earliest-finish-time-for-land-and-water-rides-ii/solutions/3976312/zui-zao-wan-cheng-lu-di-he-shui-shang-yo-bcew/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
