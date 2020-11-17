package main

import "fmt"

type Node struct {
	name     string
	priority int64
	sum      int64
	prev     *Node
	next     *Node
}

type listNode struct {
	size         int64
	firstElement *Node
	lastElement  *Node
}

func (receiver *listNode) pushBack(name string, priority int64, sum int64) { //добавляет элемент в конец очереди
	current := Node{
		name:     name,
		priority: priority,
		sum:      sum,
		prev:     nil,
		next:     nil,
	}
	if receiver.size == 0 {
		receiver.firstElement = &current
		receiver.lastElement = &current
	} else {
		current.prev = receiver.lastElement
		receiver.lastElement.next = &current
		receiver.lastElement = &current
	}
	receiver.size++
}

func (receiver *listNode) popBack() { //удаляет последний элемент очереди
	if receiver.size > 0 {
		receiver.lastElement = receiver.lastElement.prev
		receiver.lastElement.next = nil
		receiver.size--
	}
}

func (receiver *listNode) pushFront(name string, priority int64, sum int64) { //добавляет элемент в начало очереди
	current := Node{
		name:     name,
		priority: priority,
		sum:      sum,
		prev:     nil,
		next:     nil,
	}

	if receiver.size == 0 {
		receiver.firstElement = &current
		receiver.lastElement = &current
	} else {
		current.next = receiver.firstElement
		receiver.firstElement.prev = &current
		receiver.firstElement = &current
	}
	receiver.size++
}

func (receiver *listNode) popFront() { //удаляем первый элемент очереди
	if receiver.size > 0 {
		receiver.firstElement = receiver.firstElement.next
		receiver.firstElement.prev = nil
		receiver.size--
	}
}

func (receiver *listNode) sortByPriority() { //сортировка очереди по приоритету(priority)
	for i := int64(1); i < receiver.size; i++ {
		current := receiver.firstElement
		Next := current.next
		for j := int64(0); j < receiver.size-i; j++ {
			if current.priority > Next.priority {
				helperName := current.name
				helperPriority := current.priority
				helperSum := current.sum

				current.name = Next.name
				current.priority = Next.priority
				current.sum = Next.sum

				Next.name = helperName
				Next.priority = helperPriority
				Next.sum = helperSum
			}
			current = Next
			Next = current.next
		}
	}
}

func (receiver *listNode) sortBySum() { //сортировка очереди по сумме(sum)
	for i := int64(1); i < receiver.size; i++ {
		current := receiver.firstElement
		Next := current.next
		for j := int64(0); j < receiver.size-i; j++ {
			if current.sum > Next.sum {
				helperName := current.name
				helperPriority := current.priority
				helperSum := current.sum

				current.name = Next.name
				current.priority = Next.priority
				current.sum = Next.sum

				Next.name = helperName
				Next.priority = helperPriority
				Next.sum = helperSum
			}
			current = Next
			Next = current.next
		}
	}
}

func main() {
	var queue listNode

	queue.pushBack("Ikrom", 3, 90)
	queue.pushFront("Ahmad", 5, 500)
	queue.pushBack("Sulya", 6, 502)
	fmt.Println("____________________Начальный вид очереди________________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.cout()
	fmt.Println("\n__________Очередь после удаления первого элемента___________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.popFront()
	queue.cout()
	fmt.Println("\n__________Очередь после удаления последнего элемента___________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.popBack()
	queue.cout()
	fmt.Println("\n__________Очередь после добавление элемента в перёд___________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.pushFront("Shoh", 1, 60)
	queue.cout()
	fmt.Println("\n__________Очередь после добавление элемента в конец___________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.pushBack("Alex", 1, 60)
	queue.cout()
	fmt.Println("\n__________Очередь после сортировки по priority________________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.sortByPriority()
	queue.cout()
	fmt.Println("\n__________Очередь после сортировки по sum________________")
	fmt.Println("Name:	", "Priority:	", "Sum:	", "prev	     ", "next")
	queue.sortBySum()
	queue.cout()

}

func (receiver *listNode) cout() {
	current := *receiver.firstElement
	fmt.Println(current.name, "	", current.priority, "		", current.sum, "	", current.prev, current.next)
	for current.next != nil {
		current = *current.next
		fmt.Println(current.name, "	", current.priority, "		", current.sum, "	", current.prev, current.next)
	}
}
