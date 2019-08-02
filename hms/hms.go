package hms

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Hms struct {
	time.Time
}

const (
	timeLayout = "15:04:05"
)

func New(str string) (*Hms, error) {

	if str == "" || str == "null" {
		return &Hms{}, nil
	}

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return nil, err
	}

	return &Hms{t}, nil
}

func (hms Hms) String() string {
	return hms.Format(timeLayout)
}

func (hms *Hms) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), "\"")

	t, err := New(str)
	if err != nil {
		return err
	}

	*hms = *t
	return nil
}

func (hms Hms) MarshalJSON() ([]byte, error) {
	if hms.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", hms.Format(timeLayout))), nil
}

func (hms Hms) Value() (driver.Value, error) {
	return hms.String(), nil
}

func (hms *Hms) Scan(value interface{}) (err error) {

	if reflect.TypeOf(value) == nil {
		*hms = Hms{}
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*hms = Hms{v}
		return nil
	}
	return fmt.Errorf("can not convert %v to time", value)
}
