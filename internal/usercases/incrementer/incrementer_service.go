package incrementer

import (
	"github.com/exceedololo/notion-/internal/repository/key_value_storage"
	"notion-/internal/models"
)

func Increment(req models.KeyValueIncrementRequest) int {
	currentValue := key_value_storage.GetValue(req.Key)
	if currentValue == nil {
		key_value_storage.SetValue(req.Key, req.Value)
		return req.Value
	}

	newValue := *currentValue + req.Value
	key_value_storage.SetValue(req.Key, newValue)
	return newValue
}
