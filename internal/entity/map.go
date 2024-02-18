package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type String string

const StringNull String = "\x00"

// implements driver.Valuer, will be invoked automatically when written to the db
func (s String) Value() (driver.Value, error) {
	if s == StringNull {
		return nil, nil
	}
	return []byte(s), nil
}

// implements sql.Scanner, will be invoked automatically when read from the db
func (s *String) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		*s = String(v)
	case []byte:
		*s = String(v)
	case nil:
		*s = StringNull
	}
	return nil
}

func (s String) String() string {
	return string(s)
}

func (s *String) Set(v string) {
	if s == nil {
		s = new(String)
	}
	*s = String(v)
}

type Map map[string]any

var ErrKeyNotFound = errors.New("key not found")

func (p Map) Get(key string) (any, error) {
	v, ok := p[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return v, nil
}

func (p Map) Exist(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Map) GetString(key string) (string, error) {
	v, ok := p[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return fmt.Sprint(v), nil
}

func (p Map) GetInt(key string) (int64, error) {
	v, ok := p[key]
	if !ok {
		return 0, ErrKeyNotFound
	}
	value, err := decimal.NewFromString(fmt.Sprint(v))
	if err != nil {
		return 0, err
	}

	return value.IntPart(), nil
}

func (p Map) Set(key, value string) {
	p[key] = value
}

func (p Map) Delete(key string) {
	delete(p, key)
}

func (p Map) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// copies the src into the receiver [map], overwriting any existing data
func (p *Map) Scan(src any) error {
	if src == nil {
		return nil
	}

	switch srcType := src.(type) {
	case []byte:
		return json.Unmarshal(srcType, p)
	case string:
		return json.Unmarshal([]byte(srcType), p)
	default:
		return errors.New("incompatible type for map")
	}
}

// copies the receiver [map] into the destination [ptr]
func (p Map) Unmarshal(ptr any) error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, ptr)
}
