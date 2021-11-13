package main

import "testing"

type Food interface {
	GetTaste() string
}

type Pizza struct {
	name        string
	taste       string
	cookingTime int
}

type Burger struct {
	name        string
	taste       string
	cookingTime int
	price       int
}

func (p Pizza) GetTaste() string {
	return "This pizza is " + p.taste
}

func (b Burger) GetTaste() string {
	return "This burger is " + b.taste
}

func TestInterface1(t *testing.T) {
	pizza := Pizza{
		name:        "Marcherita",
		taste:       "good",
		cookingTime: 2,
	}
	burger := Burger{
		name:        "Cheeseburger",
		taste:       "excellent",
		cookingTime: 10,
		price:       100,
	}

	var food Food
	food = pizza
	if food.GetTaste() != "This pizza is good" {
		t.Errorf("Expected 'This pizza is good', got %s", food.GetTaste())
	}
	food = burger
	if food.GetTaste() != "This burger is excellent" {
		t.Errorf("Expected 'This burger is excellent', got %s", food.GetTaste())
	}
}
