package golangz

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func TestJson() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}
	var val map[string]interface{}
	var input = `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012",
		"name": "Tests",
		"number": 45
	}`

	if err := json.Unmarshal([]byte(input), &val); err != nil {
			panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
			fmt.Println(k, "[", reflect.TypeOf(v), "]", v)
	}
	assertedTypeVal := val["number"].(float64)
	fmt.Println(assertedTypeVal)

	// our target will be of type map[string]interface{}, which is a pretty generic type
	// that will give us a hashtable whose keys are strings, and whose values are of
	// type interface{}
	var data map[string]Timestamp
	var input2 = `
	{
		"created_at": "Thu May 31 00:00:01 +0000 2012"
	}`
	if err := json.Unmarshal([]byte(input2), &data); err != nil {
		panic(err)
	}

	fmt.Println(time.Time(data["created_at"]))
}

func Test() {
	TestJson()
}