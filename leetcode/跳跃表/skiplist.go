package skiplist

import "fmt"

type Node struct {
	Data      int
	NextPoint *Node
	PrePoint  *Node
	NextLevel *Node // 索引等级
}

type LinkedList struct {
	Head    *Node
	Current *Node
	Tail    *Node
	Length  int
	Level   int
}
type SkipList struct {
	// 这里为什么用 值类型？
	List       LinkedList
	FirstIndex LinkedList
	// 一级索引
	SecondIndex LinkedList
	// 二级索引
}

func InitSkipList() {
	data := []int{11, 12, 13, 19, 21, 24}
	sl := SkipList{}
	sl.initSkiplist(data)
	sl.add(11)
}

func (sl *SkipList) initSkiplist(list []int) {
	sl.List = LinkedList{}
	sl.FirstIndex = LinkedList{}
	sl.SecondIndex = LinkedList{}
	var currrentNode *Node
	for i := 0; i < len(list); i++ {
		currrentNode = new(Node)
		currrentNode.Data = list[i]
		addNode(sl, currrentNode)
	}
}
func addNode(skipList *SkipList, node *Node) {
	// 首先在List中添加
	insertIntoList(&skipList.List, node)
	if skipList.FirstIndex.Length == 0 || (skipList.List.Length%2 == 0 && skipList.List.Length > 2) {
		newNode := new(Node)
		newNode.Data = node.Data
		newNode.NextLevel = node
		insertIntoList(&skipList.FirstIndex, newNode)
		if skipList.SecondIndex.Length == 0 || (skipList.FirstIndex.Length%2 == 0 && skipList.FirstIndex.Length > 2) {
			newNode2 := new(Node)
			newNode2.Data = node.Data
			newNode2.NextLevel = newNode2
			insertIntoList(&skipList.SecondIndex, newNode2)
		}
	}
}

func insertIntoList(link *LinkedList, node *Node) {
	// LinkList 为空插入到第一个位置
	if link.Head == nil {
		link.Head = node
		link.Current = node
		link.Tail = node
	} else {
		// LinkList 不空在末尾添加
		link.Tail.NextPoint = node
		node.PrePoint = link.Tail
		link.Tail = node
	}
	link.Length++
}

/*
零级索引、一级索引和二级索引Head节点都保存着链表的首节点
添加策略：
		1、在链表头部插入，更新各级索引的首节点
		2、在链表尾部插入，

*/
func (sl *SkipList) add(num int) {
	var current *Node
	current = sl.SecondIndex.Head
	if current.Data == num {
		panic("Exist！！！")
	}
	// 策略一：直接在链表首部插入
	if num < current.Data {
		// 更新二级索引
		newNode2 := new(Node)
		newNode2.Data = num
		newNode2.NextPoint = sl.SecondIndex.Head
		sl.SecondIndex.Head.PrePoint = newNode2
		sl.SecondIndex.Head = newNode2
		// 更新一级索引
		newNode1 := new(Node)
		newNode1.Data = num
		newNode1.NextPoint = sl.FirstIndex.Head
		sl.FirstIndex.Head.PrePoint = newNode2
		sl.FirstIndex.Head = newNode2
		// 更新零级索引
		newNode := new(Node)
		newNode.Data = num
		newNode.NextPoint = sl.List.Head
		sl.List.Head.PrePoint = newNode2
		sl.List.Head = newNode2
	}
	// 在链表末尾/中间插入  重点是找到插入的位置
	for {
		if num > current.Data {
			if current.NextPoint == nil {
				// 在二级索引中遍历到末尾之后，进入一级索引查找
				if current.NextLevel != nil {
					current = current.NextLevel
				} else {
					// 当进入到零级索引，且遍历到末尾则只需要在末尾插入即可
					newNode := new(Node)
					newNode.Data = num
					current.NextPoint = newNode
					newNode.PrePoint = current
					return
				}
			} else {
				// 首先在二级索引中查找，二级索引可以跳跃的区间更大，位置更容易锁定
				fmt.Println(current.Data)
				current = current.NextPoint
			}
		} else if num < current.Data {
			// 在二级索引中如果找到正序的插入位置，则需要进入到一级索引再次查找正序的插入位置
			// 直到在零级索引中确定插入的位置
			if current.PrePoint.NextLevel != nil {
				// 向零级索引逼近
				current = current.PrePoint.NextLevel.NextPoint
			} else {
				// 到达零级索引，插入
				newNode := new(Node)
				newNode.Data = num
				current.PrePoint.NextPoint = newNode
				newNode.NextPoint = current
				current.PrePoint = newNode
				return
			}
		} else {
			fmt.Println(current.Data)
		}
	}
}
