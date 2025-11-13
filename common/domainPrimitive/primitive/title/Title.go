package title

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"unicode/utf8"
)

const TitleMaxLength = 240

type Title struct {
	value string
}

func NewTitle(text string) (Title, error) {
	text = strings.TrimSpace(text)

	if text == "" {
		return Title{}, ErrTitleIsEmpty
	}

	if utf8.RuneCountInString(text) > TitleMaxLength {
		return Title{}, ErrTitleTooLong
	}

	return Title{value: text}, nil
}

func MustTitle(text string) Title {
	t, err := NewTitle(text)
	if err != nil {
		panic(err)
	}
	return t
}

func (t Title) String() string {
	return t.value
}

func (t Title) IsZero() bool {
	return t.value == ""
}

func (t Title) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value)
}

func (t *Title) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewTitle(s)
	if err != nil {
		return err
	}
	*t = x
	return nil
}

func (t Title) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.value, nil
}

func (t *Title) Scan(src any) error {
	if src == nil {
		*t = Title{}
		return nil
	}

	switch v := src.(type) {
	case string:
		x, err := NewTitle(v)
		if err != nil {
			return err
		}
		*t = x
		return nil
	case []byte:
		x, err := NewTitle(string(v))
		if err != nil {
			return err
		}
		*t = x
		return nil
	default:
		return ErrTitleUnsupportedScanType
	}
}
