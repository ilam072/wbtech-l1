package main

import "fmt"

// Создаем структуру Human
type Human struct {
	Name   string
	Age    int
	Gender string
}

// Определяем метод Greeting для структуры Human
func (h Human) Greeting() {
	fmt.Printf("Hi! My name is %s.\n", h.Name)
}

// Создаем структуру Action
// и встраиваем в нее структуру Human
type Action struct {
	Human
}

// Определяем метод Speak для структуры Action,
// который обращается к полям Human
func (a Action) Speak() {
	fmt.Printf("I am %v years old. How old are you?\n", a.Age)
}

func main() {
	Person := Action{
		Human{
			Name:   "John",
			Age:    26,
			Gender: "Male",
		},
	}

	// Можем обращаться к полям встраиваемой структуры
	// и вызывать методы как родительской структуры, так и встраиваемой
	fmt.Println(Person.Name, Person.Age)
	Person.Greeting()
	Person.Speak()
}
