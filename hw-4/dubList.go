package dublist

import (
	"fmt"
)

type Item struct {
	Value interface{}
	Next *Item
	Prev *Item
}

func newItem(value interface{}) *Item {
	var item = new(Item)
	item.Value = value
	return item
}

type DubList struct {
	front *Item
	back *Item
}

func (list DubList) First() *Item {
	return list.front
}

func (list DubList) Last() *Item {
	return list.back
}

func (list DubList) Print() {
	curr := list.front
	if curr == nil {
		fmt.Println("")
	}
	for ;curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}

func (list *DubList) PushFront(value interface{}) {
	var newItem = newItem(value)
	if list.front == nil {
		list.front = newItem
		list.back = newItem
	} else {
		list.insertBefore(list.front, newItem)
	}
}

func (list *DubList) PushBack(value interface{}) {
	if list.back == nil {
		list.PushFront(value)
	} else {
		list.insertAfter(list.back, newItem(value))
	}
}

func (list *DubList) Remove(item *Item) {
	if item.Prev == nil {
		list.front = item.Next
	} else {
		item.Prev.Next = item.Next
	}
	if item.Next == nil {
		list.back = item.Prev
	} else {
		item.Next.Prev = item.Prev
	}
}

func (list *DubList) insertAfter(item *Item, newItem *Item) {
	newItem.Prev = item
	if item.Next == nil {
		list.back = newItem
	} else {
		newItem.Next = item.Next
		item.Next.Prev = newItem
	}
	item.Next = newItem
}

func (list *DubList) insertBefore(item *Item, newItem *Item) {
	newItem.Next = item
	if item.Prev == nil {
		list.front = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}
	item.Prev = newItem
}
