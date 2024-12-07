package leetcode54

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f     func(board [][]byte) int
		board [][]byte
		want  int
	}{
		{
			f: numRookCaptures,
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', 'p'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'}},
			want: 3,
		},
		{
			f: numRookCaptures,
			board: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'}},
			want: 0,
		},
	} {
		t.Run(strconv.Itoa(i)+"-test", func(t *testing.T) {
			if got := v.f(v.board); got != v.want {
				t.Errorf(" err got %v  want %v  \n", got, v.want)
			}
		})
	}
}

func numRookCaptures(board [][]byte) int {
	centerIdx := make([]int, 2)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				centerIdx[0] = i
				centerIdx[1] = j
			}
		}
	}

	res := 0
	log.Printf(" centerIdx %v \n", centerIdx)
	// 代表移动的四个方向
	r, l, t, d := centerIdx[1], centerIdx[1], centerIdx[0], centerIdx[0]
	for (r + l + t + d) > -3 {
		//向 右方向走
		if r < 7 && r != -1 {
			r += 1
			if board[centerIdx[0]][r] != '.' {
				if board[centerIdx[0]][r] == 'p' {
					res++
				}
				r = -1
			}
		} else {
			r = -1
		}

		//向 左方向走
		if l > 0 {
			l -= 1
			if board[centerIdx[0]][l] != '.' {
				if board[centerIdx[0]][l] == 'p' {
					res++
				}
				l = -1
			}
		} else {
			l = -1
		}

		// 向上走
		if t < 7 && t != -1 {
			t += 1
			if board[t][centerIdx[1]] != '.' {
				if board[t][centerIdx[1]] == 'p' {
					res++
				}
				t = -1
			}
		} else {
			t = -1
		}
		// 向上走
		if d > 0 && d != -1 {
			d -= 1
			if board[d][centerIdx[1]] != '.' {
				if board[d][centerIdx[1]] == 'p' {
					res++
				}
				d = -1
			}
		} else {
			d = -1
		}
	}
	log.Println(r, l, t, d)

	return res
}

func TestPrint(t *testing.T) {
	/*
	   t:. = 46
	   t:R = 82
	   t:p = 112
	   t:B = 66
	*/

	fmt.Printf("t:%c = %d\n", '.', '.')
	fmt.Printf("t:%c = %d\n", 'R', 'R')
	fmt.Printf("t:%c = %d\n", 'p', 'p')
	fmt.Printf("t:%c = %d\n", 'B', 'B')
}
