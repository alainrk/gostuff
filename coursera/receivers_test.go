package main

import (
	"fmt"
	"testing"
)

type Cat struct {
	Name string
	Age  int
}

func (d Cat) Speak() string {
	return "Woof " + d.Name + "!"
}

func (d *Cat) GetOlder() {
	fmt.Printf("From age %d ", d.Age)
	d.Age++
}

func TestReceivers1(t *testing.T) {
	d := Cat{Name: "Fido", Age: 3}
	spoken := d.Speak()
	if spoken != "Woof Fido!" {
		t.Error("Expected 'Woof Fido!' but got ", spoken)
	}
	d.GetOlder()
	if d.Age != 4 {
		t.Error("Expected 4 but got ", d.Age)
	}
}
