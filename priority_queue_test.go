package queue

import (
	"reflect"
	"testing"
)

type testHeapable struct {
	priority, index, value int
}

func (th *testHeapable) Priority(other interface{}) bool {
	if t, ok := other.(*testHeapable); ok {
		return th.priority > t.priority
	}

	return false
}

func (th *testHeapable) Index() int {
	return th.index
}

func (th *testHeapable) SetIndex(i int) {
	th.index = i
}

func TestPriorityQueue(t *testing.T) {
	testCases := []struct {
		set             []*testHeapable
		expectedLength  int
		expectedResults []*testHeapable
	}{
		{
			set:             []*testHeapable{},
			expectedLength:  0,
			expectedResults: []*testHeapable{},
		},
		{
			set: []*testHeapable{
				&testHeapable{priority: 1, index: 0, value: 0},
			},
			expectedLength: 1,
			expectedResults: []*testHeapable{
				&testHeapable{priority: 1, index: 0, value: 0},
			},
		},
		{
			set: []*testHeapable{
				&testHeapable{priority: 10, index: 0, value: 0},
				&testHeapable{priority: 8, index: 1, value: 1},
				&testHeapable{priority: 12, index: 2, value: 2},
				&testHeapable{priority: 9, index: 3, value: 3},
			},
			expectedLength: 4,
			expectedResults: []*testHeapable{
				&testHeapable{priority: 12, index: 3, value: 2},
				&testHeapable{priority: 10, index: 2, value: 0},
				&testHeapable{priority: 9, index: 1, value: 3},
				&testHeapable{priority: 8, index: 0, value: 1},
			},
		},
	}

	for _, tc := range testCases {
		pq := NewPriorityQueue()

		for _, e := range tc.set {
			pq.Push(e)
		}

		if pq.Size() != tc.expectedLength {
			t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedLength, pq.Size())
		}

		j := 0
		for pq.Size() != 0 {
			res, err := pq.Pop()

			if err != nil {
				t.Errorf("incorrect result: unexpected error: %v", err)
			}

			if !reflect.DeepEqual(res, tc.expectedResults[j]) {
				t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedResults[j], res)
			}

			j++
		}
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	testCases := []struct {
		set             []*testHeapable
		updateIndex     int
		updatePriority  int
		expectedLength  int
		expectedResults []*testHeapable
	}{
		{
			set: []*testHeapable{
				&testHeapable{priority: 1, index: 0, value: 0},
			},
			updateIndex:    0,
			updatePriority: 10,
			expectedLength: 1,
			expectedResults: []*testHeapable{
				&testHeapable{priority: 10, index: 0, value: 0},
			},
		},
		{
			set: []*testHeapable{
				&testHeapable{priority: 10, index: 0, value: 0},
				&testHeapable{priority: 8, index: 1, value: 1},
				&testHeapable{priority: 12, index: 2, value: 2},
				&testHeapable{priority: 9, index: 3, value: 3},
			},
			updateIndex:    1,
			updatePriority: 20,
			expectedLength: 4,
			expectedResults: []*testHeapable{
				&testHeapable{priority: 20, index: 3, value: 1},
				&testHeapable{priority: 12, index: 2, value: 2},
				&testHeapable{priority: 10, index: 1, value: 0},
				&testHeapable{priority: 9, index: 0, value: 3},
			},
		},
	}

	for _, tc := range testCases {
		pq := NewPriorityQueue()

		for _, e := range tc.set {
			pq.Push(e)
		}

		tc.set[tc.updateIndex].priority = tc.updatePriority

		if err := pq.Update(tc.set[tc.updateIndex]); err != nil {
			t.Errorf("incorrect result: unexpected error: %v", err)
		}

		if pq.Size() != tc.expectedLength {
			t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedLength, pq.Size())
		}

		j := 0
		for pq.Size() != 0 {
			res, err := pq.Pop()

			if err != nil {
				t.Errorf("incorrect result: unexpected error: %v", err)
			}

			if !reflect.DeepEqual(res, tc.expectedResults[j]) {
				t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedResults[j], res)
			}

			j++
		}
	}
}

func TestPriorityQueueRemove(t *testing.T) {
	testCases := []struct {
		set             []*testHeapable
		removeIndex     int
		expectedLength  int
		expectedResults []*testHeapable
	}{
		{
			set: []*testHeapable{
				&testHeapable{priority: 1, index: 0, value: 0},
			},
			removeIndex:     0,
			expectedLength:  0,
			expectedResults: []*testHeapable{},
		},
		{
			set: []*testHeapable{
				&testHeapable{priority: 10, index: 0, value: 0},
				&testHeapable{priority: 8, index: 1, value: 1},
				&testHeapable{priority: 12, index: 2, value: 2},
				&testHeapable{priority: 9, index: 3, value: 3},
			},
			removeIndex:    2,
			expectedLength: 3,
			expectedResults: []*testHeapable{
				&testHeapable{priority: 10, index: 2, value: 0},
				&testHeapable{priority: 9, index: 1, value: 3},
				&testHeapable{priority: 8, index: 0, value: 1},
			},
		},
	}

	for _, tc := range testCases {
		pq := NewPriorityQueue()

		for _, e := range tc.set {
			pq.Push(e)
		}

		if err := pq.Remove(tc.set[tc.removeIndex]); err != nil {
			t.Errorf("incorrect result: unexpected error: %v", err)
		}

		if pq.Size() != tc.expectedLength {
			t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedLength, pq.Size())
		}

		j := 0
		for pq.Size() != 0 {
			res, err := pq.Pop()

			if err != nil {
				t.Errorf("incorrect result: unexpected error: %v", err)
			}

			if !reflect.DeepEqual(res, tc.expectedResults[j]) {
				t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedResults[j], res)
			}

			j++
		}
	}
}

func TestPriorityQueueErr(t *testing.T) {
	pq := NewPriorityQueue()

	if _, err := pq.Pop(); err == nil {
		t.Errorf("incorrect result: expected error")
	}

	pq.Push(&testHeapable{priority: 1, index: 0, value: 0})
	pq.Pop()

	if _, err := pq.Pop(); err == nil {
		t.Errorf("incorrect result: expected error")
	}
}

func TestPriorityQueueUpdateErr(t *testing.T) {
	pq := NewPriorityQueue()

	e1 := &testHeapable{priority: 10, index: 0, value: 0}
	e2 := &testHeapable{priority: 5, index: 1, value: 1}

	if err := pq.Update(e1); err == nil {
		t.Errorf("incorrect result: expected error")
	}

	pq.Push(e1)

	if err := pq.Update(e2); err == nil {
		t.Errorf("incorrect result: expected error")
	}
}

func TestPriorityQueueRemoveErr(t *testing.T) {
	pq := NewPriorityQueue()

	e1 := &testHeapable{priority: 10, index: 0, value: 0}
	e2 := &testHeapable{priority: 5, index: 1, value: 1}

	if err := pq.Remove(e1); err == nil {
		t.Errorf("incorrect result: expected error")
	}

	pq.Push(e1)

	if err := pq.Remove(e2); err == nil {
		t.Errorf("incorrect result: expected error")
	}
}
