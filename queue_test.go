package gdstructure

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Queue(t *testing.T) {
	safeQueue := NewQueue()
	finishWaiter := make(chan struct{})
	rand.Seed(time.Now().UnixNano())
	operateNum := 10000
	go func() {
		for i := 0; i < operateNum; i++ {
			elem := safeQueue.Dequeue()
			if elem == nil {
				fmt.Sprintln("empty queue")
				time.Sleep(1 * time.Second)
				continue
			}

			if (i % 100) == 0 {
				t.Logf("dequeue %d\n", elem.(int))
				t.Log(safeQueue)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}

		finishWaiter <- struct{}{}
	}()

	for i := 0; i < operateNum; i++ {
		elem := rand.Int()
		safeQueue.Enqueue(elem)
		if (i % 100) == 0 {
			t.Logf("enqueue %d\n", elem)
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			t.Log(safeQueue)
		}
	}

	<-finishWaiter
	t.Log("test finished")
}
