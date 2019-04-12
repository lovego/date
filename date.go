package date

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const (
	timeLayout = "2006-01-02"
)

func New(str string) (*Date, error) {

	if str == "" || str == "null" {
		return &Date{}, nil
	}

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return nil, err
	}

	return &Date{t}, nil
}

func (date Date) String() string {
	return date.Format(timeLayout)
}

func (date *Date) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), "\"")

	t, err := New(str)
	if err != nil {
		return err
	}

	*date = *t
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	if date.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", date.Format(timeLayout))), nil
}

func (date Date) Value() (driver.Value, error) {
	return date.String(), nil
}

func (date *Date) Scan(value interface{}) (err error) {

	if reflect.TypeOf(value) == nil {
		*date = Date{}
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*date = Date{v}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}
