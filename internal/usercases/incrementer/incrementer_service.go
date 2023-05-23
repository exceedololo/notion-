package incrementer

import (
	"github.com/exceedololo/notion-/internal/repository/key_value_storage"
)

type IncrementerService struct {
	storage key_value_storage.KeyValueStorage
}

func NewIncrementerService(storage key_value_storage.KeyValueStorage) *IncrementerService {
	return &IncrementerService{
		storage: storage,
	}
}

func (s *IncrementerService) Increment(key string, value int) (int, error) {
	currentValue, err := s.storage.Get(key)
	if err != nil && err != key_value_storage.ErrKeyNotFound {
		return 0, err
	}

	var newValue int
	if currentValue != 0 {
		newValue = currentValue + value
	} else {
		newValue = value
	}

	err = s.storage.Set(key, newValue)
	if err != nil {
		return 0, err
	}

	return newValue, nil
}
