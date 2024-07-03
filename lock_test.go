package lock

import (
	"math/rand"
	"testing"
	"time"
)

func TestMutexLock(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	target := r.Int63n(1000)
	result := mutexlock(target)
	if result != target {
		t.Error("Invalid result", target, result)
	}

}

func TestCASLock(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	target := r.Int63n(1000)
	result := caslock(target)
	if result != target {
		t.Error("Invalid result", target, result)
	}
}

func BenchmarkMutex_10(b *testing.B) {
	benchmutex(b, 10)
}

func BenchmarkMutex_100(b *testing.B) {
	benchmutex(b, 100)
}

func BenchmarkMutex_200(b *testing.B) {
	benchmutex(b, 200)
}

func BenchmarkMutex_1000(b *testing.B) {
	benchmutex(b, 1000)
}

func BenchmarkCAS_10(b *testing.B) {
	benchcas(b, 10)
}

func BenchmarkCAS_100(b *testing.B) {
	benchcas(b, 100)
}

func BenchmarkCAS_200(b *testing.B) {
	benchcas(b, 200)
}

func BenchmarkCAS_1000(b *testing.B) {
	benchcas(b, 1000)
}

func benchmutex(b *testing.B, target int64) {
	for i := 0; i < b.N; i++ {
		mutexlock(target)
	}
}

func benchcas(b *testing.B, target int64) {
	for i := 0; i < b.N; i++ {
		caslock(target)
	}
}
