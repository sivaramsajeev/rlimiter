package storage

import "errors"

type IBucket interface {
	SetMaxSize(int) error
	ReduceByOne() error
}

type InMemory struct {
	currentSize int
}

func (m *InMemory) SetMaxSize(maxSize int) error {
	m.currentSize = maxSize
	return nil
}

func (m *InMemory) ReduceByOne() error {
	if m.currentSize-1 < 0 {
		return errors.New("EXHAUSTED LIMIT")
	}
	m.currentSize--
	return nil
}

func NewInMemBucket() IBucket {
	return &InMemory{}
}
