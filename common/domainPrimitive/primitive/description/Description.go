package description

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"unicode/utf8"
)

const DescriptionMaxLength = 2000

type Description struct {
	value string
}

func NewDescription(text string) (Description, error) {
	text = strings.TrimSpace(text)

	if text != "" && utf8.RuneCountInString(text) > DescriptionMaxLength {
		return Description{}, ErrDescriptionTooLong
	}

	return Description{value: text}, nil
}

func MustDescription(text string) Description {
	d, err := NewDescription(text)
	if err != nil {
		panic(err)
	}
	return d
}

func (d Description) String() string {
	return d.value
}

func (d Description) IsZero() bool {
	return d.value == ""
}

func (d Description) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.value)
}

func (d *Description) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewDescription(s)
	if err != nil {
		return err
	}
	*d = x
	return nil
}

func (d Description) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.value, nil
}

func (d *Description) Scan(src any) error {
	if src == nil {
		*d = Description{}
		return nil
	}

	switch v := src.(type) {
	case string:
		x, err := NewDescription(v)
		if err != nil {
			return err
		}
		*d = x
		return nil
	case []byte:
		x, err := NewDescription(string(v))
		if err != nil {
			return err
		}
		*d = x
		return nil
	default:
		return ErrDescriptionUnsupportedScanType
	}
}
