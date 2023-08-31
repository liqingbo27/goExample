package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Front      int
	Rear       int
	MaxSize    int
	QueueArray [6]int
}

func main() {
	//初始化队列
	queue := initQueue()

	//队列最大能存储6个元素，最大队尾下标为5
	queue.addQueue(100)
	queue.addQueue(200)
	fmt.Println("入队两个元素")
	fmt.Printf("队列: %v, 队头front: %v, 队尾rear: %v\n", queue.QueueArray, queue.Front, queue.Rear)

	v, err := queue.deleteQueue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("出队元素: %v, 队头front: %v, 队尾rear: %v\n", v, queue.Front, queue.Rear)
	}

	//入队5个元素
	for i := 1; i <= 5; i++ {
		err := queue.addQueue(i)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Printf("队列: %v, 队头front: %v, 队尾rear: %v\n", queue.QueueArray, queue.Front, queue.Rear)
}

// 队列初始化
func initQueue() (queue *Queue) {
	//初始化队列
	queue = &Queue{
		Front:      0,
		Rear:       -1,
		MaxSize:    6,
		QueueArray: [6]int{},
	}
	return queue
}

// 队列-入队
func (q *Queue) addQueue(v int) error {
	//判断队列是否已满
	if q.Rear+1 == q.MaxSize {
		return errors.New("队列已满")
	}

	q.Rear += 1              //队尾下标+1
	q.QueueArray[q.Rear] = v //数据插入队尾

	return nil
}

// 队列-出队
func (q *Queue) deleteQueue() (int, error) {
	//判断队列否为空
	if q.Rear-q.Front == -1 {
		return -1, errors.New("队列为空")
	}

	v := q.QueueArray[q.Front] //获取队列头部元素值
	q.QueueArray[q.Front] = 0  //出队数据置为0
	q.Front += 1               //队头下标+1

	return v, nil
}
