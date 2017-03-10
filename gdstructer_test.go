package gdstructer

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
	go func() {
		for i := 0; i < 100000; i++ {
			elem := safeQueue.Dequeue()
			if elem == nil {
				fmt.Sprintln("empty queue")
				time.Sleep(1 * time.Second)
				continue
			}

			if (i % 100) == 0 {
				fmt.Printf("dequeue %d\n", elem.(int))
				fmt.Println(safeQueue)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}

		finishWaiter <- struct{}{}
	}()

	for i := 0; i < 100000; i++ {
		elem := rand.Int()
		safeQueue.Enqueue(elem)
		if (i % 100) == 0 {
			fmt.Printf("enqueue %d\n", elem)
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			fmt.Println(safeQueue)
		}
	}

	<-finishWaiter
	fmt.Println("test finished")
}
