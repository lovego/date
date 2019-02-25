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

func (date Date) String() string {
	return date.Format(timeLayout)
}

func (date *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		return nil
	}
	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}
	*date = Date{Time: t}
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", date.Format(timeLayout))
	return []byte(formatted), nil
}

func (date *Date) Value() (driver.Value, error) {
	var zeroTime time.Time
	if date.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return date.Time, nil
}

func (date *Date) Scan(value interface{}) (err error) {
	v, ok := value.(time.Time)
	if ok {
		*date = Date{Time: v}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}
