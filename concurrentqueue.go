package concurrentqueue

import "sync"

// Generic Struct
type ConcurrentQueue[T comparable] struct {
	// array of items
	items []T
	// Mutual exclusion lock -> Allow only one worker access to this data.
	lock sync.Mutex
	// Cond is used to pause mulitple goroutines and wait -> Example when enqueue and dequeue.
	cond *sync.Cond
   }
   
   // Initialize ConcurrentQueue
   func New[T comparable]() *ConcurrentQueue[T] {
	q := &ConcurrentQueue[T]{}
	q.cond = sync.NewCond(&q.lock)
	return q
   }
   
   // Put the item in the queue
   func (q *ConcurrentQueue[T]) Enqueue(item T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
	// Cond signals other go routines to execute
	q.cond.Signal()
   }
   
   // Gets the item from queue
   func (q *ConcurrentQueue[T]) Dequeue() T {
	q.lock.Lock()
	defer q.lock.Unlock()
	// if Get is called before Put, then cond waits until the Put signals.
	for len(q.items) == 0 {
	 q.cond.Wait()
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
   }
   
   func (q *ConcurrentQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
   }