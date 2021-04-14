package main

import (
	"fmt"
)

func main() {
	var n int
	speed := []int{1, 2, 3}
	fmt.Scanln(&n)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanln(&tmp)
		speed = append(speed, tmp)
	}

	// fmt.Println(speed)

	for i := range speed {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", speed[i])
	}

}
func report(speed []int) []int {
	clock := 0
	aebTime := 1
	abeduring := false
	preSpeed := speed[0]
	ans := []int{preSpeed}
	for i := 0; i < len(speed); i++ {
		// 进入AEB状态
		if preSpeed-speed[i] >= 9 {
			pre := speed[i]
			abeduring = true
			for j := i + 1; j < len(speed); j++ {
				if pre-speed[j] >= 9 {
					// 速度降为0的情况
					if speed[i] == 0 {

					}
					aebTime++
					pre = speed[j]
					continue
				} else {
					if aebTime < 4 {
						break
					}
					if abeduring {
						// 统计 aeb 前后的速度
						ans = append(ans, speed[i-4:i+1]...)
						ans = append(ans, speed[i+1:j+4]...)
						clock = 0
						aebTime = 1
						i = j + 1
						abeduring = false
						break
					}
				}
			}
		}
		clock++
		if clock >= 61 && clock%61 == 0 {
			ans = append(ans, speed[i])
		}
		preSpeed = speed[i]
	}
	return ans
}
