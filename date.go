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

func Today() Date {
	now := time.Now()
	return New(now.Year(), int(now.Month()), now.Day())
}

func (date Date) Add(day int) Date {
	return Date{Time: date.Time.AddDate(0, 0, day)}
}

func New(year, month, day int) Date {
	return Date{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)}
}

func Parse(str string) (Date, error) {
	str = strings.Trim(str, "\"")
	if str == "" || str == "null" {
		return Date{}, nil
	}

	t, err := time.Parse(timeLayout, str)
	if err != nil {
		return Date{}, err
	}

	return Date{t}, nil
}

func (date Date) String() string {
	if date.Time.IsZero() {
		return ""
	}
	return date.Format(timeLayout)
}

func (date Date) MarshalJSON() ([]byte, error) {
	if date.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + date.Format(timeLayout) + `"`), nil
}

func (date *Date) UnmarshalJSON(b []byte) error {
	d, err := Parse(string(b))
	if err != nil {
		return err
	}

	*date = d
	return nil
}

func (date Date) Value() (driver.Value, error) {
	if date.Time.IsZero() {
		return []byte("NULL"), nil
	}
	return []byte("'" + date.Format(timeLayout) + "'"), nil
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

func (t Date) After(u Date) bool {
	return t.Time.After(u.Time)
}

func (t Date) Before(u Date) bool {
	return t.Time.Before(u.Time)
}

func (t Date) Equal(u Date) bool {
	return t.Time.Equal(u.Time)
}

func (t Date) Sub(u Date) time.Duration {
	return t.Time.Sub(u.Time)
}
