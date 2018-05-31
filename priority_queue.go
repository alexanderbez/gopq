package queue

import "container/heap"

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

// Push adds a Sortable item to the priority queue.
func (pq *PriorityQueue) Push(h Heapable) {
	heap.Push(pq.queue, h)
}

// Pop removes a Sortable item from the priority queue with the highest
// priority and returns it.
func (pq *PriorityQueue) Pop() Heapable {
	return heap.Pop(pq.queue).(Heapable)
}

// Size returns the size of the priority queue.
func (pq *PriorityQueue) Size() int {
	return pq.queue.Len()
}

// Update re-establishes the priority queue ordering after the Heapable element
// has changed its value. It is up to the caller to update the element
// accordingly.
func (pq *PriorityQueue) Update(h Heapable) {
	heap.Fix(pq.queue, h.Index())
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
	// oldIdx := item.Index()
	// item.SetIndex(oldIdx - 1)
	*it = old[0 : n-1]

	return item
}
