package gdstructure

import (
	"bytes"
	"fmt"
	"sync"
)

// Stack 线程安全堆栈
type Stack struct {
	baseSlice []interface{}
	locker    *sync.Mutex
}

// NewStack 创建buffer堆栈
func NewStack() (stack *Stack) {
	stack = new(Stack)
	stack.locker = new(sync.Mutex)
	return
}

// Size 返回栈中元素个数
func (s *Stack) Size() int {
	s.locker.Lock()
	defer s.locker.Unlock()
	return len(s.baseSlice)
}

// Push 将元素添加到栈顶
func (s *Stack) Push(elem interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.baseSlice = append(s.baseSlice, elem)
}

// Pop 从堆栈返回栈顶元素
func (s *Stack) Pop() (elem interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	stackSize := len(s.baseSlice)
	if stackSize < 1 {
		elem = nil
		return
	}

	elem = s.baseSlice[stackSize-1]
	s.baseSlice = s.baseSlice[:stackSize-1]

	return
}

// String 打印同步队列
// 该方法回打印当前队列元素个数和所有元素
func (s *Stack) String() string {
	s.locker.Lock()
	defer s.locker.Unlock()

	byteSlice := make([]byte, defaultSize)
	buffer := bytes.NewBuffer(byteSlice)
	stackLen := len(s.baseSlice)

	buffer.WriteString(fmt.Sprintf("Stack size: %d, elements: ", stackLen))
	for idx, elem := range s.baseSlice {
		buffer.WriteString(fmt.Sprintf("%v", elem))
		if idx != stackLen-1 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}
