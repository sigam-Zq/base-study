package leetcode67

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sort-the-students-by-their-kth-score/solutions/3021455/gou-mai-wu-pin-de-zui-da-kai-xiao-by-lee-0ov5/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f     func([][]int, int) [][]int
		score [][]int
		k     int
		want  [][]int
	}{
		{
			f:     sortTheStudents,
			score: [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}},
			k:     2,
			want:  [][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.score, v.k)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

func sortTheStudents(score [][]int, k int) [][]int {
	// // m 名学生
	// m := len(score)
	// // n 门课程
	// n := len(score[0])

	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}

// func sortTheStudents(score [][]int, k int) [][]int {
//     sort.Slice(score, func(i, j int) bool {
//         return score[i][k] > score[j][k]
//     })
//     return score
// }
