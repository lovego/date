package date

import (
	"database/sql/driver"
	"fmt"
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
	str = strings.Trim(str, "\"")
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

func (date Date) MarshalJSON() ([]byte, error) {
	if date.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + date.Format(timeLayout) + `"`), nil
}

func (date *Date) UnmarshalJSON(b []byte) error {
	t, err := New(string(b))
	if err != nil {
		return err
	}

	*date = *t
	return nil
}

func (date Date) Value() (driver.Value, error) {
	return date.String(), nil
}

func (date *Date) Scan(value interface{}) error {
	if value == nil {
		*date = Date{}
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*date = Date{v}
		return nil
	}
	return fmt.Errorf("can not convert %v to date.Date", value)
}
