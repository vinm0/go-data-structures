package datas

type linkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	value interface{}
	next  *node
}
