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
	midnight24 = "24:00:00"
	midnight   = "23:59:59"
)

func (hms Hms) Today() time.Time {
	if hms.Time.IsZero() {
		return time.Now()
	}
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		hms.Hour(), hms.Minute(), hms.Second(),
		0, now.Location())
}

func New(str string) (*Hms, error) {

	if str == "" || str == "null" {
		return &Hms{}, nil
	}

	if str == midnight24 {
		str = midnight
	}

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return nil, err
	}

	return &Hms{t}, nil
}

func (hms Hms) String() string {
	f := hms.Format(timeLayout)
	if f == midnight {
		return midnight24
	}
	return f
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

	return []byte(fmt.Sprintf("\"%s\"", hms.String())), nil
}

func (hms Hms) Value() (driver.Value, error) {
	if hms.String() == midnight24 {
		return midnight, nil
	}
	return hms.Format(timeLayout), nil
}

func (hms *Hms) Scan(value interface{}) error {

	if reflect.TypeOf(value) == nil {
		*hms = Hms{}
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*hms = Hms{v}
		return nil
	}
	return fmt.Errorf("can not convert %v to hms", value)
}
