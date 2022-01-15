package generics_test

import (
	"fmt"

	"github.com/abferm/go-generics"
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
