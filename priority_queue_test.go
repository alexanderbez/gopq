package queue

import (
	"reflect"
	"testing"
)

type testSortable struct {
	priority, index, value int
}

func (ts *testSortable) Priority(other interface{}) bool {
	if t, ok := other.(*testSortable); ok {
		return ts.priority > t.priority
	}

	return false
}

func (ts *testSortable) Index() int {
	return ts.index
}

func (ts *testSortable) SetIndex(i int) {
	ts.index = i
}

func TestPriorityQueue(t *testing.T) {
	testCases := []struct {
		set             []*testSortable
		expectedLength  int
		expectedResults []*testSortable
	}{
		{
			set:             []*testSortable{},
			expectedLength:  0,
			expectedResults: []*testSortable{},
		},
		{
			set: []*testSortable{
				&testSortable{priority: 1, index: 0, value: 0},
			},
			expectedLength: 1,
			expectedResults: []*testSortable{
				&testSortable{priority: 1, index: 0, value: 0},
			},
		},
		{
			set: []*testSortable{
				&testSortable{priority: 10, index: 0, value: 0},
				&testSortable{priority: 8, index: 1, value: 1},
				&testSortable{priority: 12, index: 2, value: 2},
				&testSortable{priority: 9, index: 3, value: 3},
			},
			expectedLength: 4,
			expectedResults: []*testSortable{
				&testSortable{priority: 12, index: 3, value: 2},
				&testSortable{priority: 10, index: 2, value: 0},
				&testSortable{priority: 9, index: 1, value: 3},
				&testSortable{priority: 8, index: 0, value: 1},
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
		for j < pq.Size() {
			e := pq.Pop()

			if !reflect.DeepEqual(e, tc.expectedResults[j]) {
				t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedResults[j], e)
			}

			j++
		}
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	testCases := []struct {
		set             []*testSortable
		updateIndex     int
		updatePriority  int
		expectedLength  int
		expectedResults []*testSortable
	}{
		{
			set: []*testSortable{
				&testSortable{priority: 1, index: 0, value: 0},
			},
			updateIndex:    0,
			updatePriority: 10,
			expectedLength: 1,
			expectedResults: []*testSortable{
				&testSortable{priority: 10, index: 0, value: 0},
			},
		},
		{
			set: []*testSortable{
				&testSortable{priority: 10, index: 0, value: 0},
				&testSortable{priority: 8, index: 1, value: 1},
				&testSortable{priority: 12, index: 2, value: 2},
				&testSortable{priority: 9, index: 3, value: 3},
			},
			updateIndex:    1,
			updatePriority: 20,
			expectedLength: 4,
			expectedResults: []*testSortable{
				&testSortable{priority: 20, index: 3, value: 1},
				&testSortable{priority: 12, index: 2, value: 2},
				&testSortable{priority: 10, index: 1, value: 0},
				&testSortable{priority: 9, index: 0, value: 3},
			},
		},
	}

	for _, tc := range testCases {
		pq := NewPriorityQueue()

		for _, e := range tc.set {
			pq.Push(e)
		}

		tc.set[tc.updateIndex].priority = tc.updatePriority
		pq.Update(tc.set[tc.updateIndex])

		if pq.Size() != tc.expectedLength {
			t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedLength, pq.Size())
		}

		j := 0
		for j < pq.Size() {
			e := pq.Pop()

			if !reflect.DeepEqual(e, tc.expectedResults[j]) {
				t.Errorf("incorrect result: expected: %v, got: %v", tc.expectedResults[j], e)
			}

			j++
		}
	}
}
