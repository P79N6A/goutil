package datastructure

import "fmt"

type Node struct {
	value string
}

type Queue struct {
	nodes []*Node
	len int64
	head int64
	tail int64
	size int64
}

func NewQueue(size int64) *Queue {
	return &Queue{
		nodes:make([]*Node, size),
		size:size,
	}
}

func (q *Queue) Pop() *Node {
	if q.len == 0 {
		return nil
	}

	q.head++
	if q.head >= q.size {
		q.head = 0
	}
	q.len--
	return q.nodes[q.head]
}

func (q *Queue)Push(node *Node) error {

	if q.len == q.size {
		return fmt.Errorf("queue is full, size:%v", q.size)
	}
	q.tail++
	if q.tail >= q.size {
		q.tail = 0
	}
	q.nodes[q.tail] = node
	q.len++
	return nil
}



