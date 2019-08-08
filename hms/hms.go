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
	IsMidnight bool `json:"isMidnight"`
}

const (
	timeLayout = "15:04:05"
	midnight24 = "24:00:00"
	midnight   = "00:00:00"
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

	is24 := false
	if str == midnight24 {
		is24 = true
		str = midnight
	}

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return nil, err
	}

	return &Hms{Time: t, IsMidnight: is24}, nil
}

func (hms Hms) String() string {
	if hms.IsMidnight {
		return midnight24
	}
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

	v, ok := value.(Hms)
	if ok {
		*hms = v
		return nil
	}
	return fmt.Errorf("can not convert %v to hms", value)
}
