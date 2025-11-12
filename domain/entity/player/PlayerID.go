package player

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type PlayerID struct {
	value string
}

func NewPlayerID(s string) (PlayerID, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return PlayerID{}, ErrPlayerIDEmpty
	}
	if len(s) > 128 {
		return PlayerID{}, ErrPlayerIDTooLong
	}
	return PlayerID{value: s}, nil
}

func MustPlayerID(s string) PlayerID {
	id, err := NewPlayerID(s)
	if err != nil {
		panic(err)
	}
	return id
}

func (id PlayerID) String() string {
	return id.value
}

func (id PlayerID) IsZero() bool {
	return id.value == ""
}

func (id PlayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *PlayerID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewPlayerID(s)
	if err != nil {
		return err
	}
	*id = x
	return nil
}

func (id PlayerID) Value() (driver.Value, error) {
	return id.value, nil
}

func (id *PlayerID) Scan(src any) error {
	if src == nil {
		*id = PlayerID{}
		return nil
	}

	switch t := src.(type) {
	case string:
		x, err := NewPlayerID(t)
		if err != nil {
			return err
		}
		*id = x
		return nil
	case []byte:
		x, err := NewPlayerID(string(t))
		if err != nil {
			return err
		}
		*id = x
		return nil
	default:
		return ErrPlayerIDUnsupportedScanType
	}
}
