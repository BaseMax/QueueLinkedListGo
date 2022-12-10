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

func main() {
	// Create a new queue
	queue := NewQueueLinkedList()

	// Add some elements to the queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	// Peek the first element
	fmt.Println(queue.Peek())

	// Peek the last element
	fmt.Println(queue.PeekTail())

	// Print the queue
	queue.Print()

	// Reverse the queue
	queue.Reverse()

	// Print the queue
	queue.Print()

	// Pop the first element
	fmt.Println(queue.Dequeue())

	// Pop the last element
	fmt.Println(queue.DequeueTail())

	// Print the queue
	queue.Print()

	// Reverse the queue (recursive)
	queue.head = queue.ReverseRecursive(queue.head)

	// Print the queue
	queue.Print()

	// Clear the queue
	queue.Clear()

	// Print the queue
	queue.Print()

	// Check if the queue is empty
	fmt.Println(queue.IsEmpty())

	// Add some elements to the queue
	queue.Enqueue(1)

	// Check if the queue is empty
	fmt.Println(queue.IsEmpty())
}
