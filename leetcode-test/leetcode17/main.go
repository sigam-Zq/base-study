package main

import "log"

func main() {

	log.Printf("-- %d \n", maxHeightOfTriangle(4, 9))
}

func maxHeightOfTriangle(red int, blue int) int {

	isPreRed := true
	// 第一次循环 （先蓝色）
	red1, blue1 := red, blue
	countB := 1
	for {

		if isPreRed {
			blue1 -= countB
			if red1 < countB+1 {
				break
			}
			isPreRed = false
		} else {
			red1 -= countB
			if blue1 < countB+1 {
				break
			}
			isPreRed = true
		}
		countB++

	}

	isPreRed = false
	red2, blue2 := red, blue
	countR := 1
	// 第一次循环 （先红色）
	for {
		if isPreRed {
			blue2 -= countR
			if red2 < countR+1 {
				break
			}
			isPreRed = false
		} else {
			red2 -= countR
			if blue2 < countR+1 {
				break
			}
			isPreRed = true
		}
		countR++
	}
	log.Printf("countB  %d  countR %d \n", countB, countR)
	if countB > countR {
		return countB
	} else {
		return countR
	}

}
