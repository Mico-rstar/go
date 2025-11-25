package queue


type SimpleQueue[T any] struct {
	inner []T
}

func (q SimpleQueue[T]) Len() int {
	return len(q.inner)
}

func (q SimpleQueue[T]) Cap() int {
	return cap(q.inner)
}

func (q SimpleQueue[T]) Front() (*T, error) {
	if len(q.inner) == 0 {
		return nil, QueueEmptyError{}
	}
	return &(q.inner[0]), nil
}

func (q SimpleQueue[T]) Back() (*T, error) {
	length := len(q.inner)
	if length == 0 {
		return nil, QueueEmptyError{}
	}
	return &(q.inner[length-1]), nil
}

func (q *SimpleQueue[T]) Push(t T) {
	q.inner = append(q.inner, t)
}

func (q *SimpleQueue[T]) Pop() {
	if len(q.inner) >= 1 {
		q.inner = q.inner[1:]
	}
}

func (q SimpleQueue[T]) Empty() bool {
	return len(q.inner) == 0
}

func New[T any](ranges ...T) Queue[T] {
	sl := make([]T, len(ranges))
	copy(sl, ranges)
	return &SimpleQueue[T]{sl}
}
