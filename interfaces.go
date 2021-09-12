package main

type Stack interface {
	Push(interface{}) bool
	Pop() interface{}
}

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
}

type Deque interface {
	Queue
	PushFront(interface{}) bool
	Pop() interface{}
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
