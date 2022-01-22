package generics_test

import (
	"fmt"
	"testing"

	"github.com/abferm/go-generics"
	"github.com/stretchr/testify/assert"
)

func ExampleMap() {
	// you can make a Map of any comparable type to any type
	m := generics.NewMap[string, string](0)
	m["hello"] = "world"
	m.Add("good", "morning")
	fmt.Println(m)
	fmt.Println(m.Has("good"))
	m.Delete("good")
	fmt.Println(m.Has("good"))
	fmt.Println(m.Get("hello"))
	fmt.Println(m.Get("good"))
	fmt.Println(m.Len())

	// Output: map[good:morning hello:world]
	// true
	// false
	// world
	//
	// 1
}

func TestMap(t *testing.T) {
	type predator string
	type prey string

	m := generics.NewMap[predator, prey](5)

	assert.Equal(t, 0, m.Len())
	assert.False(t, m.Has("fox"))

	m.Add("fox", "rabbit")

	assert.Equal(t, 1, m.Len())
	assert.True(t, m.Has("fox"))

	m.Add("whale", "krill")
	m.Add("shark", "fish")

	assert.Equal(t, 3, m.Len())
	assert.ElementsMatch(t, []predator{"fox", "whale", "shark"}, m.Keys())
	assert.ElementsMatch(t, []prey{"krill", "fish", "rabbit"}, m.Values())
	assert.ElementsMatch(t, []struct {
		Key   predator
		Value prey
	}{{"whale", "krill"}, {"shark", "fish"}, {"fox", "rabbit"}}, m.Entries())
}
