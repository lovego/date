package hms

import (
	"encoding/json"
	"fmt"
)

func ExampleHms() {

	type Stu struct {
		Name       string `json:"name"`
		FinishedAt Hms    `json:"finishedAt"`
	}

	d, err := New("14:05:00")
	fmt.Println(err)

	stu := Stu{Name: "A", FinishedAt: *d}
	b, _ := json.Marshal(stu)
	fmt.Println(string(b))

	stu = Stu{Name: "A"}
	b, _ = json.Marshal(stu)
	fmt.Println(string(b))

	stu2 := Stu{}
	data := []byte(`{"name": "W5"}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.FinishedAt.IsZero())

	data = []byte(`{"name": "W5", "finishedAt": ""}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.FinishedAt)

	// Output:
	// <nil>
	// {"name":"A","finishedAt":"14:05:00"}
	// {"name":"A","finishedAt":null}
	// true
	// 00:00:00
}
