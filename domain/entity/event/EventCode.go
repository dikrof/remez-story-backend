package event

import (
	"database/sql/driver"
	"encoding/json"
	"regexp"
	"strings"
)

var codeRe = regexp.MustCompile(`^[A-Z0-9_]{1,64}$`)

type EventCode struct {
	value string
}

func NewEventCode(s string) (EventCode, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ToUpper(s)
	if !codeRe.MatchString(s) {
		return EventCode{}, ErrInvalidEventCode
	}
	return EventCode{value: s}, nil
}

func NewCode(s string) (EventCode, error) {
	return NewEventCode(s)
}

func MustEventCode(s string) EventCode {
	c, err := NewEventCode(s)
	if err != nil {
		panic(err)
	}
	return c
}

func MustCode(s string) EventCode {
	return MustEventCode(s)
}

func (c EventCode) String() string {
	return c.value
}

func (c EventCode) IsZero() bool {
	return c.value == ""
}

func (c EventCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *EventCode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewEventCode(s)
	if err != nil {
		return err
	}
	*c = x
	return nil
}

func (c EventCode) Value() (driver.Value, error) {
	return c.value, nil
}

func (c *EventCode) Scan(src any) error {
	if src == nil {
		*c = EventCode{}
		return nil
	}

	switch t := src.(type) {
	case string:
		x, err := NewEventCode(t)
		if err != nil {
			return err
		}
		*c = x
		return nil
	case []byte:
		x, err := NewEventCode(string(t))
		if err != nil {
			return err
		}
		*c = x
		return nil
	default:
		return ErrEventCodeUnsupportedScanType
	}
}
