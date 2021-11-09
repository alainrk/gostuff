package main

import "testing"

type user struct {
	name  string
	phone string
	age   uint8
}

type fiscalcode = string

func TestMaps1(t *testing.T) {
	var users map[fiscalcode]user = make(map[fiscalcode]user)

	users["12345678901"] = user{"John", "12345678901", 30}
	users["12345678902"] = user{"Mary", "12345678902", 25}
	users["12345678903"] = user{"Mike", "12345678903", 35}

	if users["12345678901"].name != "John" {
		t.Error("Expected John, got ", users["12345678901"].name)
	}

	if len(users) != 3 {
		t.Error("Expected 3, got ", len(users))
	}
}
