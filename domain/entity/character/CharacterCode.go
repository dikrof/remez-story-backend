package character

import (
	"database/sql/driver"
	"regexp"
	"remez_story/infrastructure/errors"
	"strings"
)

var codeRe = regexp.MustCompile(`^[A-Z0-9_]{1,64}$`)

type CharacterCode struct {
	value string
}

func NewCharacterCode(s string) (CharacterCode, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "-", "_")
	s = strings.ToUpper(s)
	if !codeRe.MatchString(s) {
		return CharacterCode{}, errors.NewError("CHAR-001", "Invalid character code")
	}
	return CharacterCode{value: s}, nil
}

func MustCharacterCode(s string) CharacterCode {
	c, err := NewCharacterCode(s)
	if err != nil {
		panic(err)
	}
	return c
}

func (c CharacterCode) String() string               { return c.value }
func (c CharacterCode) IsZero() bool                 { return c.value == "" }
func (c CharacterCode) Value() (driver.Value, error) { return c.value, nil }
