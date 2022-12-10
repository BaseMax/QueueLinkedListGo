# Queue Linked-List Go

This is a simple implementation of a queue using a linked-list in Go. This is a good example of how to use a linked-list in Go.

## Functions

**Public functions:**

- `func NewQueueLinkedList() *QueueLinkedList`: Create a new queue

**Methods:**

- `func (q *QueueLinkedList) Enqueue(data interface{})`: Add a new element to the queue (to the tail)
- `func (q *QueueLinkedList) EnqueueHead(data interface{})`: Add a new element to the queue (to the head)
- `func (q *QueueLinkedList) Dequeue() interface{}`: Remove the first element from the queue (from the head)
- `func (q *QueueLinkedList) DequeueTail() interface{}`: Remove the first element from the queue (from the tail)
- `func (q *QueueLinkedList) Peek() interface{}`: Get the first element from the queue (from the head)
- `func (q *QueueLinkedList) PeekTail() interface{}`: Get the first element from the queue (from the tail)
- `func (q *QueueLinkedList) IsEmpty() bool`: Check if the queue is empty
- `func (q *QueueLinkedList) Size() int`: Get the size of the queue
- `func (q *QueueLinkedList) Clear()`: Clear the queue
- `func (q *QueueLinkedList) Print()`: Print the queue
- `func (q *QueueLinkedList) Reverse()`: Reverse the queue
- `func (q *QueueLinkedList) ReverseRecursive(Node *Node) *Node`: Reverse the queue (recursive)

**Types:**

```go
type Node struct {
	data interface{}
	next *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
}
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

© Copyright (c) 2022, Max Base
