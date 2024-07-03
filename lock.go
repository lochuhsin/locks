package lock

import (
	"sync"
	"sync/atomic"
)

func caslock(target int64) int64 {
	var count int64 = 0
	group := &sync.WaitGroup{}
	group.Add(int(target))
	for i := 0; i < int(target); i++ {
		go func(group *sync.WaitGroup) {
			for {
				old := atomic.LoadInt64(&count)
				new := old + 1
				if atomic.CompareAndSwapInt64(&count, old, new) {
					break
				}
			}
			group.Done()
		}(group)
	}
	group.Wait()
	return count
}

func mutexlock(target int64) int64 {
	var mu sync.Mutex
	var count int64 = 0
	group := &sync.WaitGroup{}
	group.Add(int(target))
	for i := 0; i < int(target); i++ {
		go func(group *sync.WaitGroup) {
			mu.Lock()
			count++
			mu.Unlock()
			group.Done()
		}(group)
	}
	group.Wait()
	return count
}
