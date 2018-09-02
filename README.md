# gopq

[![GoDoc](https://godoc.org/github.com/alexanderbez/gopq?status.svg)](https://godoc.org/github.com/alexanderbez/gopq)
[![Build Status](https://travis-ci.org/alexanderbez/gopq.svg?branch=master)](https://travis-ci.org/alexanderbez/gopq)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexanderbez/gopq)](https://goreportcard.com/report/github.com/alexanderbez/gopq)

A simple and generic priority queue implementation in Golang.

## API

Any such type that implements the `Heapable` interface may be used in a `PriorityQueue`.
The underlying element essentially must know how to give priority when compared to another
element of the same type.

If the priority queue does not require modifications (i.e. updates or removals), the type
implementing `Heapable` does not need to support indexing.

```golang
type myType struct {
	priority int
}

func (mt *myType) Index() (i int) { return }
func (mt *myType) SetIndex(_ int) {}

func (mt *myType) Priority(other interface{}) bool {
	if t, ok := other.(*myType); ok {
		return th.priority > t.priority
	}

	return false
}

pq := NewPriorityQueue()

pq.Push(&myType{priority: 1})
pq.Push(&myType{priority: 3})
pq.Push(&myType{priority: 2})

for pq.Size() != 0 {
    res, err := pq.Pop()

    // ...
}
```

## Tests

```shell
$ go test -v ./...
```

## Contributing

1. [Fork it](https://github.com/alexanderbez/gopq/fork)
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request
