package main

import "fmt"

// type node struct {
// 	key, value int
// 	pre, next  *node
// }
// type LRUCache struct {
// 	size       int
// 	cap        int
// 	data       map[int]*node
// 	head, tail *node
// }

// func initLRU(capicity int) (l *LRUCache) {
// 	l = &LRUCache{
// 		data: map[int]*node{},
// 		cap:  capicity,
// 		head: &node{
// 			0, 0, nil, nil,
// 		},
// 		tail: &node{
// 			0, 0, nil, nil,
// 		},
// 	}
// 	l.head.next = l.tail
// 	l.tail.pre = l.head
// 	return
// }
// func (l *LRUCache) moveToHead(key int) {
// 	nodeData := l.data[key]
// 	// 1.先连接
// 	nodeData.pre.next = nodeData.next
// 	nodeData.next.pre = nodeData.pre

// 	// 2.插入
// 	// 3.先连接到头部再让头部断开 先连后断
// 	nodeData.next = l.head.next
// 	l.head.next.pre = nodeData
// 	nodeData.pre = l.head
// 	l.head.next = nodeData
// }
// func (l *LRUCache) removeTail() {
// 	l.tail.pre.pre.next, l.tail.pre = l.tail, l.tail.pre.pre

// }
// func (l *LRUCache) set(key, value int) {
// 	// 1.检查缓存中是否存在
// 	if v, ok := l.data[key]; ok {
// 		l.data[key] = v
// 		// 2.若存在则更新最常用记录
// 		// 将当前节点移动到头部
// 		l.moveToHead(key)
// 		return
// 	}
// 	// 3.若不存在则添加新记录
// 	nodeData := &node{
// 		key, value, nil, nil,
// 	}
// 	nodeData.next = l.head.next
// 	l.head.next.pre = nodeData

// 	l.head.next = nodeData
// 	nodeData.pre = l.head
// 	l.data[key] = nodeData
// 	l.size++
// 	// 4.若超出缓存大小则移除最久远的记录即尾部
// 	if l.size > l.cap {
// 		// 删除尾部元素
// 		delete(l.data, l.tail.pre.key)
// 		l.removeTail()

// 	}
// }
// func (l *LRUCache) get(key int) int {
// 	if v, ok := l.data[key]; ok {
// 		// 更新最新记录，移动到头部
// 		l.moveToHead(key)
// 		return v.value
// 	}
// 	return -1
// }
type LRUCache struct {
	Size       int
	Cap        int
	Data       map[int]*DLL
	Head, Tail *DLL
}

type DLL struct {
	key, val   int
	prev, next *DLL
}

func initDLinkedNode(key, value int) *DLL {
	return &DLL{
		key: key,
		val: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Data: map[int]*DLL{},
		Head: initDLinkedNode(0, 0),
		Tail: initDLinkedNode(0, 0),
		Cap:  capacity,
	}
	l.Head.next = l.Tail
	l.Tail.prev = l.Head
	return l
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Data[key]; ok {
		l.moveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.Data[key]; ok {
		n.val = value
		l.moveToHead(n)
	} else {
		n := initDLinkedNode(key, value)
		l.Data[key] = n
		l.addToHead(n)
		l.Size++
		if l.Size > l.Cap {
			removed := l.removeTail()
			delete(l.Data, removed.key)
			l.Size--
		}
	}
}

func (l *LRUCache) addToHead(n *DLL) {
	n.prev, n.next = l.Head, l.Head.next
	l.Head.next.prev, l.Head.next = n, n
}

func (l *LRUCache) removeNode(n *DLL) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (l *LRUCache) moveToHead(n *DLL) {
	l.removeNode(n)
	l.addToHead(n)
}

func (l *LRUCache) removeTail() *DLL {
	n := l.Tail.prev
	l.removeNode(n)
	return n
}

func main() {
	L := Constructor(3)

	data := [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	res := make([]int, 0, len(data))
	for _, v := range data {
		if v[0] == 1 {
			L.Put(v[1], v[2])
		} else if v[0] == 2 {
			res = append(res, L.Get(v[1]))
		}
	}
	fmt.Println(res)
}
func LRU(operators [][]int, k int) []int {
	result := make([]int, 0)
	key := make([]int, 0, k)
	value := make([]int, 0, k)

	for _, op := range operators {
		switch op[0] {
		case 1:
			set(&key, &value, k, op[1], op[2])
		case 2:
			get(&key, &value, &result, k, op[1])
		}
	}
	return result
}

// set 数组的尾部始终为最新的数据，头部为最久的
func set(key, value *[]int, k, vkey, vvalue int) {
	if len(*key) == k {
		*key = (*key)[1:]
		*value = (*value)[1:]
	}
	// 添加在尾部
	*key = append(*key, vkey)
	*value = append(*value, vvalue)
}

// get]
func get(key, value, result *[]int, k, vkey int) {
	index := -1
	for i, v := range *key {
		if v == vkey {
			index = i
			break
		}
	}
	if index == -1 {
		*result = append(*result, -1)
	} else {
		*result = append(*result, (*value)[index])
		// 更新尾部数据即最新的数据，删除最久远的数据
		if index < k-1 {
			*value = append((*value)[0:index], append((*value)[index+1:], (*value)[index])...)
			*key = append((*key)[0:index], append((*key)[index+1:], (*key)[index])...)
		}
	}
}
