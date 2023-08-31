package main

import (
	"errors"
	"fmt"
)

type LinkQueue struct {
	Frount *QueueNode
	Rear   *QueueNode
}

type QueueNode struct {
	data int
	next *QueueNode
}

func main() {
	//初始化队列
	queue := initQueue()

	//循环入队&出队
	fmt.Println("---入队---")
	for i := 100; i <= 105; i++ {
		queue.addQueue(i)
		fmt.Println("入队元素:", queue.Rear.data)
	}

	fmt.Println("---出队---")
	for queue.Frount != nil {
		v, err := queue.deleteQueue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("出队元素:", v)
		}
	}
}

// 队列初始化
func initQueue() (queue *LinkQueue) {
	//初始化队列
	queue = &LinkQueue{
		Frount: nil,
		Rear:   nil,
	}
	return queue
}

// 队列-入队
func (lq *LinkQueue) addQueue(v int) {
	//链式存储结构，暂不考虑队列是否已满

	//新建入队节点
	newNode := &QueueNode{
		data: v,
		next: nil,
	}

	//如果Frount为空，表明为入队的第一个节点
	if lq.Frount == nil {
		lq.Frount = newNode
		lq.Rear = newNode
	}

	//原队尾指向新入队节点
	lq.Rear.next = newNode
	lq.Rear = newNode
}

// 队列-出队
func (lq *LinkQueue) deleteQueue() (int, error) {
	//判断循环队列否为空
	if lq.Frount == nil {
		return -1, errors.New("队列为空")
	}

	v := lq.Frount.data        //获取队列头部元素值
	lq.Frount = lq.Frount.next //frount指向下一个节点

	return v, nil
}
