package main

type Node struct {
	data interface{}
	next *Node
}

type QueueLinkedList struct {
	head *Node
	tail *Node
}
