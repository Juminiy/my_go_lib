package simple

import "errors"

type MyStack struct {
	stack []interface{}
}

func (s *MyStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *MyStack) Push(value interface{}) {
	if value != nil {
		s.stack = append(s.stack, value)
	}
}

func (s *MyStack) Pop() error {
	if s.IsEmpty() {
		return errors.New("Index is out of bounds,stack len = 0! ")
	}
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}

func (s *MyStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Index is out of bounds,stack len = 0! ")
	}
	value := s.stack[len(s.stack)-1]
	return value, nil
}
