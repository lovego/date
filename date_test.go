package date

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleDate() {

	type Stu struct {
		Name     string `json:"name"`
		BillDate Date   `json:"billDate"`
	}

	stu := Stu{Name: "A", BillDate: Date{time.Date(2018, 4, 1, 0, 0, 0, 0, time.Local)}}
	b, _ := json.Marshal(stu)
	fmt.Println(string(b))

	stu2 := Stu{}
	data := []byte(`{"name": "W5"}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.BillDate.IsZero())

	data = []byte(`{"name": "W5", "billDate": "2019-09-12"}`)
	json.Unmarshal(data, &stu2)
	fmt.Println(stu2.BillDate)

	// Output:
	// {"name":"A","billDate":"2018-04-01"}
	// true
	// 2019-09-12
}
