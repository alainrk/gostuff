package main

import (
	"testing"
)

func TestFindian(t *testing.T) {
	valid := []string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}
	invalid := []string{"ihhhhhn", "ina", "xian"}

	for _, v := range valid {
		if !IsIAN(v) {
			t.Error("Failed on valid", v)
		}
	}

	for _, v := range invalid {
		if IsIAN(v) {
			t.Error("Failed on invalid", v)
		}
	}
}
