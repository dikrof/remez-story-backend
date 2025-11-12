package node

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type SceneLabel struct {
	value string
}

func NewSceneLabel(s string) (SceneLabel, error) {
	s = strings.TrimSpace(s)
	if len(s) > 128 {
		return SceneLabel{}, errors.New("SceneLabel too long")
	}
	return SceneLabel{value: s}, nil
}

func (l SceneLabel) String() string {
	return l.value
}

func (l SceneLabel) IsZero() bool {
	return l.value == ""
}

func (l SceneLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.value)
}

func (l *SceneLabel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	x, err := NewSceneLabel(s)
	if err != nil {
		return err
	}
	*l = x
	return nil
}

func (l SceneLabel) Value() (driver.Value, error) {
	return l.value, nil
}

func (l *SceneLabel) Scan(src any) error {
	if src == nil {
		*l = SceneLabel{}
		return nil
	}

	switch t := src.(type) {
	case string:
		x, err := NewSceneLabel(t)
		if err != nil {
			return err
		}
		*l = x
		return nil
	case []byte:
		x, err := NewSceneLabel(string(t))
		if err != nil {
			return err
		}
		*l = x
		return nil
	default:
		return fmt.Errorf("SceneLabel: unsupported Scan type %T", src)
	}
}
