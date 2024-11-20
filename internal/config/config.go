package config

import (
	_ "embed"
	"encoding/json"
	"reflect"
	"sync"
)

//go:embed config.json
var file []byte
var lock sync.Locker
var configuration list

func isExpectedDataType(v any, dataType string) bool {
	t := reflect.TypeOf(v)
	return t.String() == dataType
}

func Get(k key, dataType string) *Variable {
	lock.Lock()
	defer lock.Unlock()

	if v, ok := configuration[k]; ok && isExpectedDataType(v.Value, dataType) {
		return &v
	}

	return nil
}

func Load() error {
	lock.Lock()
	defer lock.Unlock()

	l := make(list, 1)
	if err := json.Unmarshal([]byte(file), &l); err != nil {
		return err
	}
	configuration = l
	return nil
}

func init() {
	lock = &sync.Mutex{}
}
