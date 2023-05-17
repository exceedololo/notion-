package key_value_storage

import "sync"

var storageMap = make(map[string]*int)
var mutex = &sync.RWMutex{}

func GetValue(key string) *int {
	mutex.RLock()
	defer mutex.RUnlock()

	value, ok := storageMap[key]
	if !ok {
		return nil
	}

	return value
}

func SetValue(key string, value int) {
	mutex.Lock()
	defer mutex.Unlock()

	storageMap[key] = &value
}
