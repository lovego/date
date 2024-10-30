package date

import (
	"encoding/json"
	"fmt"
)

func ExampleParse() {
	fmt.Println(Parse("2018-04-01"))
	// Output: 2018-04-01 <nil>
}

func ExampleMarshalJSON() {
	var day = Date{}
	b, err := json.Marshal(day)
	fmt.Println(string(b), err)

	day = New(2018, 04, 01)
	b, err = json.Marshal(day)
	fmt.Println(string(b), err)
	// Output:
	// "" <nil>
	// "2018-04-01" <nil>
}

func ExampleUnmarshalJSON() {
	var day = Date{}
	err := json.Unmarshal([]byte(`"2019-09-12"`), &day)
	fmt.Println(day, err)

	day = Date{}
	err = json.Unmarshal([]byte(`2019-09-12`), &day)
	fmt.Println(day, err)

	// Output:
	// 2019-09-12 <nil>
	//  invalid character '-' after top-level value
}

func ExampleValue() {
	b, err := Date{}.Value()
	fmt.Println(string(b.([]byte)), err)

	day2 := New(2018, 04, 01)
	b, err = day2.Value()
	fmt.Println(string(b.([]byte)), err)

	// Output:
	// NULL <nil>
	// '2018-04-01' <nil>
}

func ExampleCompare() {
	day1 := New(2021, 10, 1)
	day2 := New(2021, 10, 2)
	day3 := New(2021, 10, 2)
	fmt.Println(day1.After(day2))
	fmt.Println(day1.Before(day2))
	fmt.Println(day1.Equal(day2))
	fmt.Println(day2.After(day3))
	fmt.Println(day2.Before(day3))
	fmt.Println(day2.Equal(day3))

	// Output:
	// false
	// true
	// false
	// false
	// false
	// true
}
