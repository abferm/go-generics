package factory_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/abferm/go-generics/factory"
	"github.com/stretchr/testify/assert"
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

	myPizza, _, err := registry.Build(context.Background(), "pizza", []string{"cheese", "pepperoni", "sausage"})
	fmt.Println(myPizza.Describe(), err)

	myIcecream, _, err := registry.Build(context.Background(), "icecream", []string{"chocolate"})
	fmt.Println(myIcecream.Describe(), err)

	_, _, err = registry.Build(context.Background(), "icecream", []string{"chocolate", "vanilla"})
	fmt.Println(err)

	// Output: factory already registered for icecream
	// Pizza piled high with: [cheese pepperoni sausage] <nil>
	// Sweet chocolate icecream <nil>
	// our icecream must have exactly one flavor
}

func TestRegistry(t *testing.T) {
	registry := factory.NewRegistry[string, []string, food]()

	err := registry.Register("pizza", buildPizza)
	assert.NoError(t, err)

	err = registry.Register("icecream", buildIcecream)
	assert.NoError(t, err)

	_, _, err = registry.Build(context.Background(), "pizza", []string{"cheese", "pepperoni", "sausage"})
	assert.NoError(t, err)

	registry.Deregister("pizza")
	_, _, err = registry.Build(context.Background(), "pizza", []string{"cheese", "pepperoni", "sausage"})
	assert.EqualError(t, err, "no factory registered for ID pizza", "Pizza factory should have been removed from the registry")
	_, err = registry.GetConfiguredFactory("pizza", []string{"cheese", "pepperoni", "sausage"})
	assert.EqualError(t, err, "no factory registered for ID pizza", "Pizza factory should have been removed from the registry")

}

func TestContextWithFactory(t *testing.T) {
	registry := factory.NewRegistry[string, []string, food]()

	err := registry.Register("pizza", buildPizza)
	assert.NoError(t, err)

	err = registry.Register("icecream", buildIcecream)
	assert.NoError(t, err)

	favoriteFood, err := registry.GetConfiguredFactory("pizza", []string{"cheese", "pepperoni", "sausage"})
	assert.NoError(t, err)
	ctx, err := factory.ContextWithFactory(context.Background(), favoriteFood)
	assert.NoError(t, err)

	f, err := factory.FromContext[food](ctx)
	assert.NoError(t, err)

	food1, _, _ := favoriteFood(context.Background())
	food2, _, _ := f(context.Background())
	assert.Equal(t, food1, food2, "Expected the same food")

	favoriteDesert, err := registry.GetConfiguredFactory("icecream", []string{"peach"})
	assert.NoError(t, err)

	ctx, err = factory.ContextWithFactory(ctx, favoriteDesert)
	assert.EqualError(t, err, "factory.ConfiguredFactory[factory.food] already registered on context")

	f, err = factory.FromContext[food](ctx)
	assert.NoError(t, err)
	food2, _, _ = f(context.Background())
	assert.Equal(t, food1, food2, "stored factory should not have changed")

	_, err = factory.FromContext[int](ctx)
	assert.EqualError(t, err, "no factory.ConfiguredFactory[int] registered on context")
}

type food interface {
	Hot() bool
	Entree() bool
	Describe() string
}

type pizza struct {
	toppings []string
}

func buildPizza(_ context.Context, toppings []string) (food, func() error, error) {
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

func buildIcecream(_ context.Context, flavor []string) (food, func() error, error) {
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
