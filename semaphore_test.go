package semaphore

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(1)

	var wg sync.WaitGroup

	tnow := time.Now()
	// test init
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			s.Grab()
			time.Sleep(time.Second)
			s.Release()
			wg.Done()
		}()
	}
	wg.Wait()

	if time.Now().Sub(tnow) < 2*time.Second {
		t.Fatalf("expected 1 maxRunner so 2 seconds, got %f seconds", time.Now().Sub(tnow).Seconds())
	}
}

func TestThreeConcurrent(t *testing.T) {
	s := NewSemaphore(3)

	var wg sync.WaitGroup

	tnow := time.Now()
	// test init
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			s.Grab()
			time.Sleep(time.Second)
			s.Release()
			wg.Done()
		}()
	}
	wg.Wait()

	if time.Now().Sub(tnow) > 2*time.Second {
		t.Fatalf("expected 3 concurrent which should finish in < 2 seconds, got %f seconds", time.Now().Sub(tnow).Seconds())
	}
}
