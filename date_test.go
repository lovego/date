package date

import (
	"encoding/json"
	"fmt"
)

func ExampleNew() {
	fmt.Println(New("2018-04-01"))
	// Output: 2018-04-01 <nil>
}

func ExampleMarshalJSON() {
	var day = &Date{}
	b, err := json.Marshal(day)
	fmt.Println(string(b), err)

	day, _ = New("2018-04-01")
	b, err = json.Marshal(day)
	fmt.Println(string(b), err)
	// Output:
	// null <nil>
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
	var day = &Date{}
	fmt.Println(day.Value())

	day, _ = New("2018-04-01")
	fmt.Println(day.Value())

	// Output:
	// NULL <nil>
	// 2018-04-01 <nil>
}
