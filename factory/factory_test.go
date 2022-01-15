package factory_test

import (
	"fmt"

	"github.com/abferm/go-generics/factory"
)

func ExampleFactory() {
	registry := factory.NewRegistry[string, []string, food]()

	err := registry.Register("pizza", buildPizza)
	if err != nil {
		panic(err)
	}

	err = registry.Register("icecream", buildIcecream)
	if err != nil {
		panic(err)
	}

	err = registry.Register("icecream", buildIcecream)
	fmt.Println(err)

	myPizza, _, err := registry.Build("pizza", []string{"cheese", "pepperoni", "sausage"})
	fmt.Println(myPizza.Describe(), err)

	myIcecream, _, err := registry.Build("icecream", []string{"chocolate"})
	fmt.Println(myIcecream.Describe(), err)

	_, _, err = registry.Build("icecream", []string{"chocolate", "vanilla"})
	fmt.Println(err)

	// Output: factory already registered for icecream
	// Pizza piled high with: [cheese pepperoni sausage] <nil>
	// Sweet chocolate icecream <nil>
	// our icecream must have exactly one flavor
}

type food interface {
	Hot() bool
	Entree() bool
	Describe() string
}

type pizza struct {
	toppings []string
}

func buildPizza(toppings []string) (food, func() error, error) {
	return &pizza{toppings}, func() error { return nil }, nil
}

func (p pizza) Hot() bool {
	return true
}

func (p pizza) Entree() bool {
	return true
}

func (p pizza) Describe() string {
	return fmt.Sprintf("Pizza piled high with: %s", p.toppings)
}

type icecream struct {
	flavor string
}

func buildIcecream(flavor []string) (food, func() error, error) {
	if len(flavor) != 1 {
		return nil, nil, fmt.Errorf("our icecream must have exactly one flavor")
	}
	return &icecream{flavor: flavor[0]}, func() error { return nil }, nil
}

func (i icecream) Hot() bool {
	return false
}

func (i icecream) Entree() bool {
	return false
}

func (i icecream) Describe() string {
	return fmt.Sprintf("Sweet %s icecream", i.flavor)
}
