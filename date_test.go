package date

import (
	"encoding/json"
	"fmt"
)

func ExampleDate() {

	type Stu struct {
		Name     string `json:"name"`
		BillDate Date   `json:"billDate"`
	}

	d, err := New("2018-04-01")
	fmt.Println(err)

	stu := Stu{Name: "A", BillDate: *d}
	b, _ := json.Marshal(stu)
	fmt.Println(string(b))

	stu = Stu{Name: "A"}
	b, _ = json.Marshal(stu)
	fmt.Println(string(b))

	stu2 := Stu{}
	data := []byte(`{"name": "W5"}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.BillDate.IsZero())

	data = []byte(`{"name": "W5", "billDate": "2019-09-12"}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.BillDate)

	// Output:
	// <nil>
	// {"name":"A","billDate":"2018-04-01"}
	// {"name":"A","billDate":null}
	// true
	// 2019-09-12
}
