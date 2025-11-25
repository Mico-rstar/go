package queue

type QueueEmptyError struct {}

func (e QueueEmptyError) Error() string {
	return "failed to get value from empty queue"
}

