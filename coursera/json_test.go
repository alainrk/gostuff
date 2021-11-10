package main

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson1(t *testing.T) {
	people := []Person{}
	jsonstr := "[{\"name\":\"Foo\",\"age\":20},{\"name\":\"Bar\",\"age\":42}]"

	json.Unmarshal([]byte(jsonstr), &people)
	jsonstrExp, err := json.Marshal(people)

	if err != nil {
		panic(err)
	}
	if jsonstr != string(jsonstrExp) {
		t.Error("jsonstr != jsonstrExp")
	}
}
