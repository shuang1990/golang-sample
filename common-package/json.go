package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Student struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Score float32 `json:"score"`
	Courses []string `json:"courses"`
	Extra map[string]float32 `json:"extra"`
}

func main() {
	student := Student{
		ID: 1001,
		Name: "zhangsan",
		Age: 16,
		Score: 85.4,
	}

	courses := make([]string, 0)
	courses = append(courses, "Chinese")
	courses = append(courses, "English")
	courses = append(courses, "Computer")

	bodyInfo := make(map[string]float32)
	bodyInfo["height"] = 174.5
	bodyInfo["weight"] = 60.4

	student.Courses = courses
	student.Extra = bodyInfo

	//convert Student struct to json data
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("student encode to json failed:", err)
		return
	}
	fmt.Println(string(data))

	//convert json data to Student struct
	s := Student{}
	err = json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println("decode json failed:", err)
		return
	}

	fmt.Printf("%+v %T\n", s, s)

	//convert slice of map to json
	orderList := make([]map[string]string, 5)
	for i := 0; i < len(orderList); i++ {
		order := make(map[string]string)
		order["order_id"] = "201802011234" + strconv.Itoa(i)
		order["cash"] = strconv.Itoa(i + 1 * 100)
		orderList[i] = order
	}

	fmt.Println(orderList)
	data, _ = json.Marshal(orderList)
	fmt.Println(string(data))

	//convert json data to slice of map
	orderListDecode := []map[string]string{}
	err = json.Unmarshal(data, &orderListDecode)
	fmt.Println(orderListDecode)
}