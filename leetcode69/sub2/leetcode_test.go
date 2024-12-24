package leetcode69

import (
	"log"
	"reflect"
	"sort"
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
			}, isDebug: false,
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
			want:    []int{-1, 0, -1, 0, -1, 0, -1, 0, -1, 0, -1, 0, 0, 0, 0, 0},
			isDebug: false,
		},
		{
			o:       ExamRoom{},
			action:  []string{"ExamRoom", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave", "seat", "leave", "leave", "seat", "seat", "leave", "leave", "seat", "seat", "leave"},
			args:    []int{10, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 7, 0, 3, 0, 3, 0, 9, 0, 0, 8, 0, 0, 0, 8, 0, 0, 2},
			want:    []int{-1, 0, 9, 4, -1, -1, 0, 4, 2, 6, 1, 3, 5, 7, 8, -1, -1, 0, 4, -1, 7, -1, 3, -1, 3, -1, 9, -1, -1, 0, 8, -1, -1, 0, 8, -1},
			isDebug: true,
		},
		{
			o:       ExamRoom{},
			action:  []string{"ExamRoom", "seat", "seat", "seat", "leave", "leave", "seat", "seat", "seat", "seat", "seat", "seat", "seat"},
			args:    []int{8, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0},
			want:    []int{-1, 0, 7, 3, -1, -1, 7, 0, 5, 1, 2, 4, 6},
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
						log.Println("Seat this.isSit")
						log.Println(v.o.isSit)
						log.Println("Seat this.haveKey")
						log.Println(v.o.haveKey)
					}

				case "leave":
					v.o.Leave(v.args[i])
					t = -1
					if v.isDebug {
						log.Println("Leave this.isSit")
						log.Println(v.o.isSit)
						log.Println("Leave this.haveKey")
						log.Println(v.o.haveKey)
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
	// map[i]=true i座位在坐
	isSit map[int]bool
	// 存在map中的key
	haveKey []int
	len     int
}

func Constructor(n int) ExamRoom {

	return ExamRoom{
		isSit:   make(map[int]bool),
		haveKey: make([]int, 0),
		len:     n,
	}
}

func (this *ExamRoom) Seat() int {

	if len(this.isSit) == 0 {
		this.haveKey = append(this.haveKey, 0)
		this.isSit[0] = true
		return 0
	} else if len(this.isSit) == 1 {
		if this.len > 2*this.haveKey[0] {
			this.haveKey = append(this.haveKey, this.len-1)
			this.isSit[this.len-1] = true
			return this.len - 1
		} else {
			this.haveKey = append(this.haveKey, 0)
			this.isSit[0] = true
			return 0
		}
	} else {
		// len(this.isSit) >= 2
		var maxDisIdx, maxDis int

		// 当0位置没有人坐时候 初始化取第一个
		if _, ok := this.isSit[0]; !ok {
			init := 1
			for i := 1; i < this.len; i++ {
				if _, ok := this.isSit[i]; ok {
					maxDis = init + 1
					break
				}
				init++
			}
		}
		sort.Ints(this.haveKey)
		log.Println("sort haveKey")
		log.Println(this.haveKey)

		log.Println("init maxDis")
		log.Println(maxDis)
		for i := 1; i < len(this.haveKey); i++ {
			// if maxDis+1 < this.haveKey[i]-this.haveKey[i-1] {
			if maxDis/2 < (this.haveKey[i]-this.haveKey[i-1])/2 {
				maxDis = this.haveKey[i] - this.haveKey[i-1]
				maxDisIdx = i
			}
		}

		if maxDisIdx == 0 {

			if this.len == len(this.isSit)+1 {
				for i := 0; i < this.len; i++ {
					if _, ok := this.isSit[i]; !ok {
						this.haveKey = append(this.haveKey, i)
						this.isSit[i] = true
						return i
					}
				}
			}
			this.haveKey = append(this.haveKey, 0)
			this.isSit[0] = true
			return 0
		}

		log.Println("maxDisIdx,maxDis")
		log.Println(maxDisIdx, maxDis)
		k := (this.haveKey[maxDisIdx] + this.haveKey[maxDisIdx-1]) / 2
		log.Println("this.haveKey[maxDisIdx],this.haveKey[maxDisIdx-1],k")
		log.Println(this.haveKey[maxDisIdx], this.haveKey[maxDisIdx-1], k)
		this.haveKey = append(this.haveKey, k)
		this.isSit[k] = true
		return k
	}
}

func (this *ExamRoom) Leave(p int) {

	delete(this.isSit, p)

	newHaveKey := make([]int, len(this.isSit))
	i := 0
	for k := range this.isSit {
		newHaveKey[i] = k
		i++
	}
	this.haveKey = newHaveKey

}
