package generics_test

import (
	"fmt"

	"github.com/abferm/go-generics"
)

func ExampleCastSlice() {
	float32s := []float32{
		1, 2, 3, 4.5,
	}

	float64s, err := generics.CastSlice[float32, float64](float32s)
	fmt.Printf("%T : %f\n", float32s, float32s)
	fmt.Printf("%T : %f\n", float64s, float64s)
	fmt.Println(err)

	stringerInts := []stringerInt{
		1, 2, 99, 364,
	}

	stringers, err := generics.CastSlice[stringerInt, fmt.Stringer](stringerInts)
	fmt.Printf("%T : %v\n", stringerInts, stringerInts)
	fmt.Printf("%T : %v\n", stringers, stringers)
	fmt.Println(err)

	// Output: []float32 : [1.000000 2.000000 3.000000 4.500000]
	// []float64 : [1.000000 2.000000 3.000000 4.500000]
	// <nil>
	// []generics_test.stringerInt : [1 2 99 364]
	// []fmt.Stringer : [1 2 99 364]
	// <nil>
}

type stringerInt int

func (i stringerInt) String() string {
	return fmt.Sprintf("%d", i)
}
