package main

import "fmt"

type MaxQueue struct {
	valQueue []int
	deQueue  []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		valQueue: make([]int, 0, 16),
		deQueue:  make([]int, 0, 16),
	}
}
func (this *MaxQueue) Max_value() int {
	if len(this.valQueue) == 0 {
		return -1
	}
	return this.deQueue[0]
}
func (this *MaxQueue) Push_back(value int) {
	if len(this.valQueue) == 0 || this.deQueue[len(this.deQueue)-1] >= value {
		this.deQueue = append(this.deQueue, value)
	} else {
		tmp := make([]int, 0, len(this.deQueue))
		for i := len(this.deQueue) - 1; i >= 0; i-- {
			if this.deQueue[i] < value {
				tmp = append(tmp, this.deQueue[i])
				this.deQueue = this.deQueue[:len(this.deQueue)-1]
			}
			this.deQueue = append(this.deQueue, value)
			this.deQueue = append(this.deQueue, tmp...)
		}
	}
	this.valQueue = append(this.valQueue, value)
}
func (this *MaxQueue) Pop_front() int {
	if len(this.valQueue) == 0 {
		return -1
	}
	res := this.valQueue[0]
	this.valQueue = this.valQueue[1:]
	return res
}

func main() {
	mq := Constructor()
	array := make([]int, 0, 16)
	mq.Push_back(1)
	mq.Push_back(2)
	array = append(array, mq.Max_value())
	array = append(array, mq.Pop_front())
	array = append(array, mq.Max_value())
	fmt.Println(array)
}
