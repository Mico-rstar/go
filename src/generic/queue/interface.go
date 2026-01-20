package queue

type Queue[T any] interface {
	// Element access
	Front() (*T, error) 
	Back() (*T, error)
	Empty() bool
	Len() int
	Cap() int
	Push(T)
	Pop()
}
