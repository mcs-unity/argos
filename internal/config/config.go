package config

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"sync"
)

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

func Load(path string) error {
	if configuration != nil {
		return errors.New("variables are already loaded")
	}

	lock.Lock()
	defer lock.Unlock()

	l := make(list, 1)
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &l); err != nil {
		return err
	}
	configuration = l
	return nil
}

func init() {
	lock = &sync.Mutex{}
}
