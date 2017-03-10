package gdstructer

import (
	"bytes"
	"sync"
)

type bufferNode struct {
	data *bytes.Buffer
	next *bufferNode
}

// BufferStack buffer队列，避免重复申请内存
type BufferStack struct {
	head   *bufferNode
	locker *sync.Mutex
	Size   int
}

// NewBufferStack 创建buffer堆栈
func NewBufferStack() (bq *BufferStack) {
	head := new(bufferNode)
	locker := new(sync.Mutex)
	return &BufferStack{head, locker, 0}
}

// Push 添加buffer到堆栈
func (p *BufferStack) Push(b *bytes.Buffer) {
	bn := new(bufferNode)
	bn.data = b

	p.locker.Lock()
	bn.next = p.head.next
	p.head.next = bn
	// p.Size = p.Size + 1
	p.Size++
	p.locker.Unlock()
}

// Pop 从堆栈拿buffer
func (p *BufferStack) Pop() (b *bytes.Buffer) {
	p.locker.Lock()
	if p.head.next != nil {
		bn := p.head.next
		p.head.next = bn.next
		p.Size--
		b = bn.data
	}
	p.locker.Unlock()

	return
}
