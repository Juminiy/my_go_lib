package simple

import "errors"

type MyQueue struct {
	queue []interface{}
}

func (q *MyQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Index is out of bounds,queue len = 0! ")
	}
	return q.queue[0], nil
}

func (q *MyQueue) Back() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Index is out of bounds,queue len = 0! ")
	}
	return q.queue[len(q.queue)-1], nil
}

func (q *MyQueue) Push(value interface{}) {
	if value != nil {
		q.queue = append(q.queue, value)
	}
}

func (q *MyQueue) Pop() error {
	if q.IsEmpty() {
		return errors.New("Index is out of bounds,queue len = 0! ")
	}
	q.queue = q.queue[1:]
	return nil
}

func (q *MyQueue) Traverse() {

}
