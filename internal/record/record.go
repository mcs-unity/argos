package record

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func (r Record) Write(p payload, e event) error {
	if p == nil {
		return errors.New("f parameter can't be nil pointer")
	}
	message := ""

	switch v := p.(type) {
	case error:
		message = v.Error()
	case string:
		message = v
	default:
		return errors.New("unable for log message")
	}

	s := fmt.Sprintf("[%s] %s %s\n", time.Now().UTC(), e, message)

	fmt.Print(s)
	if _, err := r.Writer.Write([]byte(s)); err != nil {
		return err
	}

	return nil
}

func New(path string) (IRecord, error) {
	if s := strings.Trim(path, ""); s == "" {
		return nil, errors.New("path can't be an empty string")
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/events.log", path), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return &Record{f}, nil
}
