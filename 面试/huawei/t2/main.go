package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type file struct {
	filename string
	row      int
}
type SortBy []*file

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := make(map[string]*file, 0)
	order := make(SortBy, 0, 2)
	var a string
	newLine, _, _ := reader.ReadLine()
	for len(newLine) != 0 {
		a = string(newLine)
		if _, ok := m[a]; ok {
			m[a].row++
		} else {
			m[a] = &file{filename: a, row: 1}
		}
		newLine, _, _ = reader.ReadLine()
	}
	for index := range m {
		order = append(order, m[index])
	}

	sort.Sort(order)
	for i := 0; i < len(order); i++ {
		fmt.Println(order[i].filename + " " + fmt.Sprintf("%d", order[i].row))
	}
}

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].row > a[j].row }
