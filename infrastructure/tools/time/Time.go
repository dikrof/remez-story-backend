package time

import (
	"database/sql/driver"
	"time"
)

const (
	LayoutDateDMYByPoint       = "02.01.2006"
	LayoutDateDMYBySlash       = "02/01/2006"
	LayoutDateYMDByDash        = "2006-01-02"
	LayoutDateTimeDMYMMByPoint = "02.01.2006 15:04"
	LayoutDateTimeDMYSSByPoint = "02.01.2006 15:04:05"
)

type Time struct {
	time.Time
}

func Now() *Time {
	return &Time{time.Now().UTC()}
}

func Empty() *Time {
	return &Time{}
}

func FromTime(t time.Time) *Time {
	if t.IsZero() {
		return Empty()
	}
	return &Time{t.UTC()}
}

func FromUnixNano(nanoseconds int64) *Time {
	if nanoseconds == 0 {
		return Empty()
	}
	return &Time{time.Unix(0, nanoseconds).UTC()}
}

func FromUnixMillis(milliseconds int64) *Time {
	if milliseconds == 0 {
		return Empty()
	}
	return &Time{time.UnixMilli(milliseconds).UTC()}
}

func Parse(layout string, value string) (*Time, error) {
	parsedValue, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return FromTime(parsedValue), nil
}

func (t *Time) Local() *Time {
	if t == nil || t.IsZero() {
		return Empty()
	}
	return &Time{t.Time.Local()}
}

func (t *Time) Add(duration time.Duration) *Time {
	if t == nil || t.IsZero() {
		return Empty()
	}
	newTime := t.Time.Add(duration)
	return &Time{newTime}
}

func (t *Time) Sub(duration time.Duration) *Time {
	if t == nil || t.IsZero() {
		return Empty()
	}
	newTime := t.Time.Add(-duration)
	return &Time{newTime}
}

func (t *Time) Equal(other *Time) bool {
	if t == nil || other == nil {
		return t == other
	}
	return t.Time.Equal(other.Time)
}

func (t *Time) Before(other *Time) bool {
	if t == nil || other == nil {
		return false
	}
	return t.Time.Before(other.Time)
}

func (t *Time) After(other *Time) bool {
	if t == nil || other == nil {
		return false
	}
	return t.Time.After(other.Time)
}

func (t *Time) Unix() int64 {
	if t == nil || t.IsZero() {
		return 0
	}
	return t.Time.Unix()
}

func (t *Time) UnixMilli() int64 {
	if t == nil || t.IsZero() {
		return 0
	}
	return t.Time.UnixMilli()
}

func (t *Time) UnixNano() int64 {
	if t == nil || t.IsZero() {
		return 0
	}
	return t.Time.UnixNano()
}

func (t *Time) IsZero() bool {
	if t == nil {
		return true
	}
	return t.Time.IsZero()
}

func (t *Time) Value() (driver.Value, error) {
	if t == nil || t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(src any) error {
	if src == nil {
		*t = Time{}
		return nil
	}

	switch v := src.(type) {
	case time.Time:
		*t = *FromTime(v)
		return nil
	default:
		return nil
	}
}
