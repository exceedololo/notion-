package incrementer

import (
	"github.com/exceedololo/notion-/internal/models"
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

func (s *IncrementerService) Increment(req models.KeyValueIncrementRequest) (int, error) {
	currentValue, err := s.storage.Get(req.Key)
	if err != nil && err != key_value_storage.ErrKeyNotFound {
		return 0, err
	}

	var newValue int
	if currentValue != nil {
		newValue = *currentValue + req.Value
	} else {
		newValue = req.Value
	}

	err = s.storage.Set(req.Key, newValue)
	if err != nil {
		return 0, err
	}

	return newValue, nil
}
