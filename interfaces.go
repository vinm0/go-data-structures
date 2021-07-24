package main

type Stack interface {
	Push(interface{}) bool
	Pop() interface{}
	Peek() interface{}
}

type Queue interface {
	Enqueue(interface{}) bool
	Dequeue() interface{}
	Peek() interface{}
}

type Deque interface {
	Queue
	PushFront(interface{}) bool
	Pop() interface{}
	PeekFront() interface{}
}

type TreeNode interface {
	Parent() TreeNode
	LeftChild() TreeNode
	RightChild() TreeNode
}

type TrieNode interface {
	Parent() TrieNode
	Children() map[interface{}]TrieNode
}

type GraphNode interface {
	Neighbors() []TreeNode
	// TODO: Edges() []Edge
}
