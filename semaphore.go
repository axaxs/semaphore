package semaphore

import (
	"sync"
)

// Semaphore represents the parent object
// Do not use this directly, as it will have no internals and panic
// Use NewSemaphore to create
type Semaphore struct {
	semaChan chan struct{}
	wg       sync.WaitGroup
}

// Grab grabs a spot in the semaphore
func (s *Semaphore) Grab() {
	s.wg.Add(1)
	<-s.semaChan
}

// Release frees a spot in the semaphore
func (s *Semaphore) Release() {
	s.semaChan <- struct{}{}
	s.wg.Done()
}

// Wait is used to wait for all workers to be finished
func (s *Semaphore) Wait() {
	s.wg.Wait()
}

// NewSemaphore creates a new semaphore with maxConcurrent runners
func NewSemaphore(maxConcurrent int) *Semaphore {
	ch := make(chan struct{}, maxConcurrent)
	for i := 0; i < maxConcurrent; i++ {
		ch <- struct{}{}
	}
	return &Semaphore{semaChan: ch, wg: sync.WaitGroup{}}
}
