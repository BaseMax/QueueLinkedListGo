/**
 *
 * @Name: Go - Data Structure and Algorithm (Queue Linked-List)
 * @Author: Max Base
 * @Date: 2022/10/12
 * @Repository: https://github.com/BaseMax/QueueLinkedListGo
 * @License: GPL-3.0
 *
 */

package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
}

/**
 * @summary: Create a new queue
 * @return: A new queue
 */
func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{}
}

/**
 * @summary: Add a new element to the queue (to the tail)
 * @param: data - the data to add
 */
func (q *QueueLinkedList) Enqueue(data interface{}) {
	node := &Node{data: data}

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

/**
 * @summary: Add a new element to the queue (to the head)
 * @param: data - the data to add
 */
func (q *QueueLinkedList) EnqueueHead(data interface{}) {
	node := &Node{data: data}

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		node.next = q.head
		q.head = node
	}
}

/**
 * @summary: Remove the first element from the queue (from the head)
 * @return: The data of the first element
 */
func (q *QueueLinkedList) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}

	node := q.head
	q.head = q.head.next

	return node.data
}

/**
 * @summary: Remove the first element from the queue (from the tail)
 * @return: The data of the first element
 */
func (q *QueueLinkedList) DequeueTail() interface{} {
	if q.head == nil {
		return nil
	}

	node := q.head
	q.head = q.head.next
	return node.data
}

/**
 * @summary: Get the first element from the queue (from the head)
 * @return: The data of the first element
 */
func (q *QueueLinkedList) Peek() interface{} {
	if q.head == nil {
		return nil
	}

	return q.head.data
}

/**
 * @summary: Get the first element from the queue (from the tail)
 * @return: The data of the first element
 */
func (q *QueueLinkedList) PeekTail() interface{} {
	if q.head == nil {
		return nil
	}

	return q.tail.data
}

/**
 * @summary: Check if the queue is empty
 * @return: True if the queue is empty, false otherwise
 */
func (q *QueueLinkedList) IsEmpty() bool {
	return q.head == nil
}

/**
 * @summary: Get the size of the queue
 * @return: The size of the queue
 */
func (q *QueueLinkedList) Size() int {
	size := 0

	for node := q.head; node != nil; node = node.next {
		size++
	}

	return size
}

/**
 * @summary: Clear the queue
 */
func (q *QueueLinkedList) Clear() {
	q.head = nil
	q.tail = nil
}

/**
 * @summary: Print the queue
 */
func (q *QueueLinkedList) Print() {
	for node := q.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

/**
 * @summary: Reverse the queue
 */
func (q *QueueLinkedList) Reverse() {
	var prev *Node
	curr := q.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	q.head = prev
}

/**
 * @summary: Reverse the queue (recursive)
 */
func (q *QueueLinkedList) ReverseRecursive(Node *Node) *Node {
	if Node == nil || Node.next == nil {
		return Node
	}

	newHead := q.ReverseRecursive(Node.next)
	Node.next.next = Node
	Node.next = nil

	return newHead
}
