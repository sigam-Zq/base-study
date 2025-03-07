package leetcode69

import (
	"log"
	"reflect"
	"strconv"
	"testing"
)

// https://leetcode.cn/problems/exam-room/description/?envType=daily-question&envId=2024-12-23

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		o       ExamRoom
		action  []string
		args    []int
		want    []int
		isDebug bool
	}{
		{
			o:      ExamRoom{},
			action: []string{"ExamRoom", "seat", "seat", "seat", "seat", "leave", "seat"},
			args:   []int{10, 0, 0, 0, 0, 4, 0},
			want: []int{-1,
				0,
				9,
				4,
				2,
				-1,
				5,
			},
		},
		{
			o:       ExamRoom{},
			action:  []string{"ExamRoom", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "leave"},
			args:    []int{10, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:    []int{-1, 0, 9, 4, -1, -1, 0, 4, 2, 6, 1, 3, 5, 7, 8, -1},
			isDebug: false,
		},
		{
			o:       ExamRoom{},
			action:  []string{"ExamRoom", "seat", "seat", "leave", "leave", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "leave"},
			args:    []int{10, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:    []int{-1, 0, 9, -1, -1, 0, 9, 4, 2, 6, 1, 3, 5, 7, 8, -1},
			isDebug: false,
		},
		{
			o:       ExamRoom{},
			action:  []string{"ExamRoom", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave"},
			args:    []int{1000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:    []int{-1, 0, 9, -1, -1, 0, 9, 4, 2, 6, 1, 3, 5, 7, 8, -1},
			isDebug: false,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := make([]int, len(v.want))

			for i, vv := range v.action {
				if v.isDebug {
					log.Println("-----------------------")
					log.Println(i, vv, v.args[i])
					log.Println("-----------------------")
				}
				var t int
				switch vv {
				case "ExamRoom":
					v.o = Constructor(v.args[i])
					t = -1
				case "seat":
					t = v.o.Seat()
					if v.isDebug {
						log.Println("Seat this.dis")
						log.Println(v.o.dis)
						log.Println("Seat this.isSit")
						log.Println(v.o.isSit)
					}

				case "leave":
					v.o.Leave(v.args[i])
					t = -1
					if v.isDebug {
						log.Println("Leave this.dis")
						log.Println(v.o.dis)
						log.Println("Leave this.isSit")
						log.Println(v.o.isSit)
					}
				}
				got[i] = t
			}

			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	}
}

type ExamRoom struct {
	// 是否入座
	isSit []bool
	// 距离有座的距离
	dis []int
	len int
}

func Constructor(n int) ExamRoom {
	s := make([]bool, n)
	dis := make([]int, n)
	for i := range dis {
		dis[i] = n
	}
	return ExamRoom{
		isSit: s,
		dis:   dis,
		len:   n,
	}
}

func (this *ExamRoom) Seat() int {

	var sitIdx int
	for i := 0; i < this.len; i++ {
		if this.dis[i] > this.dis[sitIdx] {
			sitIdx = i
		}
	}

	// 更新 dis
	this.isSit[sitIdx] = true
	this.dis[sitIdx] = 0
	disT := 1
	for h, l := sitIdx+1, sitIdx-1; h < this.len || l >= 0; {
		if h < this.len {
			this.dis[h] = min(disT, this.dis[h])
		}
		if l >= 0 {
			this.dis[l] = min(disT, this.dis[l])
		}
		disT++
		h++
		l--
	}
	// log.Println("Seat this.dis")
	// log.Println(this.dis)
	// log.Println("Seat this.isSit")
	// log.Println(this.isSit)
	return sitIdx
}

func (this *ExamRoom) Leave(p int) {
	var maxIdx int
	for i := 0; i < this.len; i++ {
		if this.dis[i] > this.dis[maxIdx] {
			maxIdx = i
		}
	}

	this.isSit[p] = false

	// 重新计算dis
	this.reCalcDis()

	// log.Println("Leave this.isSit")
	// log.Println(this.isSit)
	// log.Println("Leave this.dis")
	// log.Println(this.dis)
}

func (t *ExamRoom) reCalcDis() {
	newDis := make([]int, t.len)
	for i := 0; i < t.len; i++ {

		if t.isSit[i] {
			newDis[i] = 0
			continue
		}

		disT := 1
		for h, l := i+1, i-1; h < t.len || l >= 0; {
			if (h < t.len && t.isSit[h]) || (l >= 0 && t.isSit[l]) {
				newDis[i] = disT
				break
			}
			disT++
			h++
			l--
		}
		if newDis[i] == 0 {
			newDis[i] = disT
		}
	}
	t.dis = newDis
}

/*
问:


func (this *ExamRoom) Seat() int {

	var sitIdx int
	for i := 0; i < this.len; i++ {
		if this.dis[i] > this.dis[sitIdx] {
			sitIdx = i
		}
	}

	// 更新 dis
	this.isSit[sitIdx] = true
	this.dis[sitIdx] = 0
	disT := 1
	for h, l := sitIdx+1, sitIdx-1; h < this.len || l >= 0; {
		if h < this.len {
			this.dis[h] = min(disT, this.dis[h])
		}
		if l >= 0 {
			this.dis[l] = min(disT, this.dis[l])
		}
		disT++
		h++
		l--
	}
	log.Println("Seat this.dis")
	log.Println(this.dis)
	log.Println("Seat this.isSit")
	log.Println(this.isSit)
	return sitIdx
}

func (this *ExamRoom) Leave(p int) {
	var maxIdx int
	for i := 0; i < this.len; i++ {
		if this.dis[i] > this.dis[maxIdx] {
			maxIdx = i
		}
	}

	this.isSit[p] = false

	// 重新计算dis
	this.reCalcDis()

	log.Println("Leave this.isSit")
	log.Println(this.isSit)
	log.Println("Leave this.dis")
	log.Println(this.dis)
}

func (t *ExamRoom) reCalcDis() {
	newDis := make([]int, t.len)
	for i := 0; i < t.len; i++ {
		disT := 1
		for h, l := i+1, i-1; h < t.len && l >= 0; {
			if (h < t.len && t.isSit[h]) || (l >= 0 && t.isSit[l]) {
				newDis[i] = disT
			}
			disT++
			h++
			l--
		}
	}
	t.dis = newDis
}


根据打印

2024/12/23 18:47:28 [0 1 2 1 0 1 2 2 1 0]
2024/12/23 18:47:28 Seat this.isSit
2024/12/23 18:47:28 [true false false false true false false false false true]
2024/12/23 18:47:28 Seat this.dis
2024/12/23 18:47:28 [0 1 0 1 0 1 2 2 1 0]
2024/12/23 18:47:28 Seat this.isSit
2024/12/23 18:47:28 [true false true false true false false false false true]
2024/12/23 18:47:28 Leave this.isSit
2024/12/23 18:47:28 [true false true false false false false false false true]
2024/12/23 18:47:28 Leave this.dis
2024/12/23 18:47:28 [0 1 2 3 4 4 3 2 1 0]
2024/12/23 18:47:28 Seat this.dis
2024/12/23 18:47:28 [0 1 2 1 0 1 2 2 1 0]
2024/12/23 18:47:28 Seat this.isSit
2024/12/23 18:47:28 [true false true false true false false false false true]

这里  的
Leave 和 reCalcDis

均未改变其结构体内部的值

这里是为什么，需要怎么解决


Fix

从你提供的代码和日志来看，`Leave` 方法调用了 `reCalcDis` 方法，但 `reCalcDis` 方法似乎没有正确更新 `dis` 切片的值。问题可能出在 `reCalcDis` 方法的逻辑上。

在 `reCalcDis` 方法中，`newDis[i]` 只有在找到一个已入座的位置时才会被更新为 `disT`。如果没有找到任何已入座的位置，`newDis[i]` 将保持为零值（即 0）。这可能导致 `dis` 切片没有被正确更新。

要解决这个问题，你可以在 `reCalcDis` 方法中确保 `newDis[i]` 被正确计算。以下是一个可能的修正：

```go
func (t *ExamRoom) reCalcDis() {
    newDis := make([]int, t.len)
    for i := 0; i < t.len; i++ {
        if t.isSit[i] {
            newDis[i] = 0
            continue
        }

        disT := 1
        for h, l := i+1, i-1; h < t.len || l >= 0; {
            if (h < t.len && t.isSit[h]) || (l >= 0 && t.isSit[l]) {
                newDis[i] = disT
                break
            }
            disT++
            h++
            l--
        }

        // 如果没有找到任何已入座的位置，设置为最大距离
        if newDis[i] == 0 {
            newDis[i] = disT
        }
    }
    t.dis = newDis
}
```

在这个修正中：

- 如果 `t.isSit[i]` 为 `true`，则 `newDis[i]` 被设置为 0。
- 如果在 `for` 循环中没有找到任何已入座的位置，`newDis[i]` 被设置为 `disT`，即到达边界的距离。

请确保在调用 `reCalcDis` 之后，`dis` 切片被正确更新。这样可以确保 `Leave` 方法在调用 `reCalcDis` 后，`dis` 切片反映出正确的距离信息。

*/

func (t *ExamRoom) reCalcDisFix() {
	newDis := make([]int, t.len)
	for i := 0; i < t.len; i++ {
		if t.isSit[i] {
			newDis[i] = 0
			continue
		}

		disT := 1
		for h, l := i+1, i-1; h < t.len || l >= 0; {
			if (h < t.len && t.isSit[h]) || (l >= 0 && t.isSit[l]) {
				newDis[i] = disT
				break
			}
			disT++
			h++
			l--
		}

		// 如果没有找到任何已入座的位置，设置为最大距离
		if newDis[i] == 0 {
			newDis[i] = disT
		}
	}
	t.dis = newDis
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */

func CalPoint(i int) *int {
	return &i
}
