package main

/*
查找数组中某个数的父结点和左右孩子结点，比如已知索引为i的数，那么

1.父结点索引：(i-1)/2（这里计算机中的除以2，省略掉小数）
2.左孩子索引：2*i+1
3.右孩子索引：2*i+2
*/

func sift(li []int, low, high int) {
	i := low
	j := 2*i + 1 // 左孩子
	tmp := li[i]
	for j <= high {
		if j < high && li[j] < li[j+1] {
			j++
		}
		if tmp < li[j] {
			li[i] = li[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	li[i] = tmp
}

func heapSort(li []int) {
	for i := len(li)/2 - 1; i >= 0; i-- {
		sift(li, i, len(li)-1)
	}
	for j := len(li) - 1; j > 0; j-- {
		li[0], li[j] = li[j], li[0]
		sift(li, 0, j-1)
	}
}
