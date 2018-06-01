package queue

import (
	"container/heap"
	"errors"
	"fmt"
)

type (
	// Heapable reflects an interface that an item must implement in order to
	// be placed into a priority queue. If the priority queue does not require
	// modifications (i.e. updates or removals), these can be implemented as
	// no-ops.
	Heapable interface {
		Priority(other interface{}) bool
		Index() int
		SetIndex(i int)
	}

	// items is the internal container for a priority queue. It's implementation
	// should not be exposed publically.
	items []Heapable

	// PriorityQueue implements a priority queue of Heapable items. Each item
	// must implement the Heapable interface and is responsible for giving each
	// item priority respective to every other item. The main benefit of this
	// implementation is to abstract away any direct heap usage.
	PriorityQueue struct {
		queue *items
	}
)

// NewPriorityQueue returns a reference to a new initialized PriorityQueue.
func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{queue: &items{}}

	heap.Init(pq.queue)
	return pq
}

// Push adds a Sortable item to the priority queue. The complexity is
// logarithmic.
func (pq *PriorityQueue) Push(h Heapable) {
	heap.Push(pq.queue, h)
}

// Pop removes a Sortable item from the priority queue with the highest
// priority and returns it. The complexity is logarithmic. If the index is
// beyond the size of the priority queue, an error is returned.
func (pq *PriorityQueue) Pop() (res Heapable, err error) {
	if pq.Size() > 0 {
		res = heap.Pop(pq.queue).(Heapable)
	} else {
		err = errors.New("cannot pop from an empty priority queue")
	}

	return
}

// Size returns the size of the priority queue.
func (pq *PriorityQueue) Size() int {
	return pq.queue.Len()
}

// Update re-establishes the priority queue ordering after the Heapable element
// has changed its value. It is up to the caller to update the element
// accordingly. The complexity is logarithmic. If the index is beyond the size
// of the priority queue, an error is returned.
func (pq *PriorityQueue) Update(h Heapable) (err error) {
	if h.Index() < pq.Size() {
		heap.Fix(pq.queue, h.Index())
	} else {
		err = fmt.Errorf("invalid element index: %d", h.Index())
	}

	return
}

// Remove removes a Heapable item from the priority queue. The item is
// retrieved by fetching it's index. Items are then "fixed" into their
// appropriate order once the item is removed. The complexity is logarithmic.
// If the index is beyond the size of the priority queue, an error is returned.
func (pq *PriorityQueue) Remove(h Heapable) (err error) {
	if h.Index() < pq.Size() {
		heap.Remove(pq.queue, h.Index())
	} else {
		err = fmt.Errorf("invalid element index: %d", h.Index())
	}

	return
}

// Len implements the sort.Interface.
func (it items) Len() int {
	return len(it)
}

// Less implements the sort.Interface.
func (it items) Less(i, j int) bool {
	return it[i].Priority(it[j])
}

// Swap implements the sort.Interface.
func (it items) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
	it[i].SetIndex(i)
	it[j].SetIndex(j)
}

// Push implements the heap interface.
func (it *items) Push(x interface{}) {
	n := len(*it)
	item := x.(Heapable)
	item.SetIndex(n)
	*it = append(*it, item)
}

// Pop implements the heap interface.
func (it *items) Pop() interface{} {
	old := *it
	n := len(old)
	item := old[n-1]
	*it = old[0 : n-1]

	return item
}
