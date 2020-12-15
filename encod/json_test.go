package encod

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func TestJson(t *testing.T) {
	person := Person{Name: "张三", Age: 35}
	value, err := json.Marshal(person)
	if err == nil {
		fmt.Println("value:", string(value))
	}

	per := new(Person)
	er := json.Unmarshal(value, per)
	if er == nil {
		fmt.Println("name:", per.Name, "age:", per.Age)
	}

}
