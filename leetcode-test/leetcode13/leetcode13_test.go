package leetcode13

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	args := []struct {
		name        string
		f           func(int, [][]int, []int) int
		maxTime     int
		edges       [][]int
		passingFees []int
		want        int
	}{
		{
			name:        "one",
			f:           minCost,
			maxTime:     30,
			edges:       [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}},
			passingFees: []int{5, 1, 2, 20, 20, 3},
			want:        11,
		},
		{
			name:        "two",
			f:           minCost,
			maxTime:     29,
			edges:       [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}},
			passingFees: []int{5, 1, 2, 20, 20, 3},
			want:        48,
		},
		{
			name:    "three",
			f:       minCost,
			maxTime: 500,
			edges: [][]int{{9, 7, 18}, {26, 3, 12}, {28, 45, 33}, {47, 10, 27}, {34, 18, 38}, {32, 13, 39}, {32, 26, 32}, {12, 0, 2}, {4, 1, 7}, {5, 3, 2}, {39, 25, 27}, {45, 10, 34}, {3, 19, 5}, {25, 32, 23}, {30, 10, 47}, {37, 2, 31}, {10, 32, 15}, {23, 14, 19}, {22, 6, 14}, {45, 39, 38}, {39, 21, 30}, {42, 17, 42}, {20, 17, 15}, {24, 0, 27}, {2, 46, 11}, {2, 24, 13}, {36, 22, 30}, {2, 1, 31}, {41, 35, 45}, {4, 19, 20}, {32, 27, 33}, {38, 46, 1}, {21, 11, 15}, {33, 41, 2}, {45, 18, 30}, {8, 33, 50}, {37, 11, 6}, {25, 17, 42}, {45, 39, 33}, {7, 4, 49}, {17, 42, 36}, {36, 16, 9}, {46, 25, 24}, {43, 4, 6}, {35, 13, 28}, {1, 28, 1}, {34, 35, 15}, {38, 1, 15}, {16, 6, 28}, {13, 0, 42}, {3, 30, 24}, {43, 27, 35}, {8, 0, 45}, {27, 20, 47}, {6, 16, 47}, {0, 34, 35}, {0, 35, 3}, {40, 11, 24}, {1, 0, 49}, {44, 20, 32}, {26, 12, 17}, {3, 2, 25}, {37, 25, 42}, {27, 1, 15}, {36, 25, 38}, {24, 47, 33}, {33, 28, 15}, {25, 43, 37}, {47,
				31, 47}, {29, 10, 50}, {11, 1, 21}, {29, 3, 48}, {1, 25, 10}, {48, 17, 16}, {19, 24, 22}, {30, 7, 2}, {11, 22, 19}, {20, 42, 41}, {27, 3, 48}, {17, 0, 34}, {19, 14, 32}, {49, 2, 20}, {10, 3, 38}, {0, 49, 13}, {6, 3, 28}, {42, 23, 6}, {14, 8, 1}, {35, 16, 3}, {17, 7, 40}, {18, 7, 49}, {36, 35, 13}, {14, 40, 45}, {16, 33, 11}, {31, 22, 33}, {38, 15, 48}, {15, 14, 25}, {37, 13, 37}, {44, 32, 7}, {48, 1, 31}, {33, 12, 20}, {22, 26, 23}, {4, 10, 11}, {43, 28, 43}, {19, 8, 14}, {35, 31, 33}, {28, 27, 19}, {40, 11, 36}, {36, 43, 28}, {22, 21, 15}},
			passingFees: []int{199, 505, 107, 961, 682, 400, 304, 517, 512, 18, 334, 627, 893, 412, 922, 289, 19, 161, 206, 879, 336, 831, 577, 802, 139, 348, 440, 219, 273, 691, 99, 858, 389, 955, 561, 353, 937, 904, 858, 704, 548, 497, 787, 546, 241, 67, 743, 42, 87, 137},
			want:        336,
		},
	}

	for _, v := range args {
		t.Run(v.name, func(t *testing.T) {
			if got := v.f(v.maxTime, v.edges, v.passingFees); v.want != got {
				t.Errorf(" got = %v ,want = %v \n", got, v.want)
			}
		})

	}

}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	cityRoadIdx := make(map[int][]int)

	// target
	target := 0

	// 整理出 edge  到城市的映射
	// city  --> edge_idx
	for i, v := range edges {

		if target < v[0] {
			target = v[0]
		}
		if _, ok := cityRoadIdx[v[0]]; !ok {
			cityRoadIdx[v[0]] = make([]int, 0)
		}
		cityRoadIdx[v[0]] = append(cityRoadIdx[v[0]], i)

		if target < v[1] {
			target = v[1]
		}
		if _, ok := cityRoadIdx[v[1]]; !ok {
			cityRoadIdx[v[1]] = make([]int, 0)
		}
		cityRoadIdx[v[1]] = append(cityRoadIdx[v[1]], i)
	}

	//到达重点路线的全部 可能的 费用
	feesAll := make([]int, 0)

	// queue := make()
	// 组织队列
	// for timeCost, idx := 0, 0; timeCost < maxTime; {
	// 	nextStopIdxs := cityRoadIdx[idx]

	// 	for _, stopIdx := range nextStopIdxs {
	// 		road := edges[stopIdx]

	// 	}

	// }

	alreadyCity := make(map[int]bool)
	var bfs func(startCity, targetCity, timeConsume, feesConsume int) (feesCost int)
	bfs = func(startCity, targetCity, timeConsume, feesConsume int) (feesCost int) {
		alreadyCity[startCity] = true

		feesConsume += passingFees[startCity]
		log.Printf("add fees cost %d \n", passingFees[startCity])
		log.Printf("add fees cost res %d \n", feesConsume)
		if startCity == targetCity {
			feesAll = append(feesAll, feesConsume)
			return feesConsume
		}
		nextStopIdxs := cityRoadIdx[startCity]
		for _, stopIdx := range nextStopIdxs {
			timeTmp := timeConsume
			road := edges[stopIdx]

			timeTmp += road[2]
			if timeTmp > maxTime {
				return -1
			}

			// 找出到达的那个城市 road 1 2 排除掉 startCity
			toCity := 0
			if road[0] == startCity {
				toCity = road[1]
			} else if road[1] == startCity {
				toCity = road[0]
			} else {
				println("error cityRoadIdx 组织错误")
				return -1
			}

			if alreadyCity[toCity] {
				continue
			}

			log.Printf(" %d ---->  %d   time const %d \n", startCity, toCity, road[2])
			res := bfs(toCity, targetCity, timeTmp, feesConsume)
			if res == -1 {
				continue
			}
			//  else {
			// feesConsume = res
			// }

		}
		// 这里的返回值有问题 不准 怀疑是多加了
		return feesConsume
	}

	res := bfs(0, target, 0, 0)
	log.Printf("---- bfs res %d \n", res)
	log.Printf("---- fees All %v \n", feesAll)

	if len(feesAll) == 0 {
		return -1
	}
	min := feesAll[0]
	for _, v := range feesAll[1:] {
		if min < v {
			min = v
		}
	}

	return min
}
