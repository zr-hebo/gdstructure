# gdstructure
gdstructure是一个golang实现的线程安全的FIFO队列
# Install

Use go get to install this package:<br>
$ go get github.com/zr-hebo/gdstructure

# Usage

	import (
		gds "github.com/zr-hebo/gdstructure"
	)
  
	// Int类型队列
	queue1 = gds.NewQueue()
	queue1.Enqueue(1)
	data := queue.Dequeue()
	intVal := data.(int)
	
	// string类型队列
	queue2 = gds.NewQueue()
	queue2.Enqueue("string")
	
	data := queue2.Dequeue()
	strVal := data.(string)
