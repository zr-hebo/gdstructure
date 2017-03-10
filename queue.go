package gdstructure

import (
	"bytes"
	"fmt"
	"sync"
)

// Queue 基于动态数组的同步队列
type Queue struct {
	baseSlice []interface{}
	queue     []interface{}
	locker    *sync.Mutex
	startPos  int
}

// NewQueue 创建同步队列
func NewQueue() (q *Queue) {
	q = new(Queue)
	q.baseSlice = make([]interface{}, 0, defaultSize)
	q.queue = q.baseSlice
	q.locker = new(sync.Mutex)
	return
}

// Size 同步队列元素个数
func (q *Queue) Size() int {
	return len(q.queue)
}

// Enqueue 同步队列添加元素
func (q *Queue) Enqueue(elem interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	if q.startPos+q.Size()+1 <= cap(q.baseSlice) {
		q.queue = append(q.queue, elem)

	} else {
		if q.startPos < (cap(q.baseSlice) / 4) {
			q.baseSlice = make(
				[]interface{}, 0, recommendCap(cap(q.baseSlice)))
		}

		newQueue := q.baseSlice[:q.Size()]
		copy(newQueue, q.queue)
		q.startPos = 0
		newQueue = append(newQueue, elem)
		q.queue = newQueue
	}
}

func recommendCap(oldCap int) int {
	newCap := oldCap * 2
	if oldCap > bigCap {
		newCap = oldCap + oldCap/4
	}

	return newCap
}

// Dequeue 元素出同步队列
func (q *Queue) Dequeue() (elem interface{}) {
	q.locker.Lock()
	defer q.locker.Unlock()

	if q.Size() < 1 {
		return
	}

	elem = q.queue[0]
	q.queue = q.queue[1:]
	q.startPos++
	return
}

// String 打印同步队列
// 该方法回打印当前队列元素个数和所有元素
func (q *Queue) String() string {
	q.locker.Lock()
	defer q.locker.Unlock()

	byteSlice := make([]byte, defaultSize)
	buffer := bytes.NewBuffer(byteSlice)
	buffer.WriteString(fmt.Sprintf("Queue size: %d, elements: ", q.Size()))
	for idx, elem := range q.queue {
		buffer.WriteString(fmt.Sprintf("%v", elem))
		if idx != q.Size()-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
