package id

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"
)

type EntityID struct {
	value int64
}

func NewEntityID(v int64) (EntityID, error) {
	if v <= 0 {
		return EntityID{}, ErrInvalidEntityID
	}
	return EntityID{value: v}, nil
}

func EntityIDFrom(s string) (EntityID, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return EntityID{}, ErrEntityIDIsEmpty
	}

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return EntityID{}, ErrCreateEntityID(s)
	}

	return NewEntityID(n)
}

func (x EntityID) Int64() int64 {
	return x.value
}

func (x EntityID) IsZero() bool {
	return x.value == 0
}

func (x EntityID) String() string {
	return strconv.FormatInt(x.value, 10)
}

func (x EntityID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.value)
}

func (x *EntityID) UnmarshalJSON(b []byte) error {
	var n int64
	if err := json.Unmarshal(b, &n); err == nil {
		v, err := NewEntityID(n)
		if err != nil {
			return err
		}
		*x = v
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, err := EntityIDFrom(s)
	if err != nil {
		return err
	}
	*x = v
	return nil
}

func (x EntityID) Value() (driver.Value, error) {
	return x.value, nil
}

func (x *EntityID) Scan(src any) error {
	if src == nil {
		*x = EntityID{}
		return nil
	}

	switch t := src.(type) {
	case int64:
		v, err := NewEntityID(t)
		if err != nil {
			return err
		}
		*x = v
		return nil
	case []byte:
		n, err := strconv.ParseInt(string(t), 10, 64)
		if err != nil {
			return err
		}
		v, err := NewEntityID(n)
		if err != nil {
			return err
		}
		*x = v
		return nil
	case string:
		n, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return err
		}
		v, err := NewEntityID(n)
		if err != nil {
			return err
		}
		*x = v
		return nil
	default:
		return ErrUnsupportedScanType
	}
}
