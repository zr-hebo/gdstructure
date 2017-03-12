package gdstructure

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_Stack(t *testing.T) {
	safeStack := NewStack()
	finishWaiter := make(chan struct{})
	rand.Seed(time.Now().UnixNano())
	go func() {
		time.Sleep(1 * time.Second)
		defer func() {
			finishWaiter <- struct{}{}
		}()

		i := 0
		for {
			elem := safeStack.Pop()
			if elem == nil {
				fmt.Sprintln("empty stack")
				return
			}

			if (i % 100) == 0 {
				t.Logf("pop %d\n", elem.(int))
				t.Log(safeStack)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}

			i++
		}

	}()

	operateNum := 10000
	for i := 0; i < operateNum; i++ {
		elem := rand.Int()
		safeStack.Push(elem)
		if (i % 100) == 0 {
			t.Logf("push %d\n", elem)
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			t.Log(safeStack)
		}
	}

	<-finishWaiter
	t.Log(safeStack)
}
