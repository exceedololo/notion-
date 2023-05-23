package key_value_storage

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

type Storage struct {
	data map[string]int
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]int),
	}
}

func (s *Storage) Get(key string) (int, error) {

	value, ok := s.data[key]
	if !ok {
		return 0, ErrKeyNotFound
	}

	return value, nil
}

func (s *Storage) Set(key string, value int) {
	s.data[key] = value
}
