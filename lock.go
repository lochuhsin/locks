package lock

import (
	"sync"
	"sync/atomic"
)

/**
 * Simple but sometimes not efficient, since other goroutine should be waiting for lock
 */
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

/**
 * Since cas doesn't acquire lock in general, it is more performant to use.
 * However, the threads that are not able to add value will stuck in the infinite loop
 * which waste a lot of CPU clocks. (Spinning)
 */
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
