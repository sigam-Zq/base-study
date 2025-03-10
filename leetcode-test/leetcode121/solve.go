package leetcode121

import (
	"sort"
)

func maximumBeautyLeetCode(items [][]int, queries []int) []int {
	// 将物品按价格升序排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	n := len(items)
	// 按定义修改美丽值
	// fmt.Printf("befor---%v \n", items)
	for i := 1; i < n; i++ {
		if items[i][1] < items[i-1][1] {
			items[i][1] = items[i-1][1]
		}
	}
	// fmt.Printf("after---%v \n", items)
	// 二分查找处理查询
	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = query(items, q)
	}
	return res
}

func query(items [][]int, q int) int {
	l, r := 0, len(items)
	for l < r {
		mid := l + (r-l)/2
		if items[mid][0] > q {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == 0 {
		// 此时所有物品价格均大于查询价格
		return 0
	} else {
		// 返回小于等于查询价格的物品的最大美丽值
		return items[l-1][1]
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/most-beautiful-item-for-each-query/solutions/1101845/mei-yi-ge-cha-xun-de-zui-da-mei-li-zhi-b-d8jw/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
