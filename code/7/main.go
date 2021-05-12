package main

func findKth(a []int, n int, K int) int {
	// write code here
	quick(a, 0, n)
	return a[K-1]

}

func main() {
	arr := []int{2, 3, 1}
	quick(arr, 0, 3)

}
func quick(a []int, l, r int) {
	if l >= r {
		return
	}
	// 基准点
	mid := (r-l)>>1 + l
	i, j := l, r-1
	tmp := a[mid]
	for i < j {
		for i < j && a[j] <= tmp {
			j--
		}
		a[mid] = a[j]
		for i < j && a[i] >= tmp {
			i++
		}
		a[j] = a[i]
	}
	a[i] = tmp
	mid = i
	quick(a, l, mid-1)
	quick(a, mid+1, r)
}
