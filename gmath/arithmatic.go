package gmath

import "github.com/abferm/go-generics"

// Add casts y to the type of x and returns the sum
// cast to type Z
func Add[X, Y, Z generics.Numeric](x X, y Y) Z {
	return Z(x + X(y))
}

// Sub casts y to the type of x and returns x-y
// cast to type Z
func Sub[X, Y, Z generics.Numeric](x X, y Y) Z {
	return Z(x - X(y))
}

// Mul casts y to the type of x and returns x*y
// cast to type Z
func Mul[X, Y, Z generics.Numeric](x X, y Y) Z {
	return Z(x * X(y))
}

// Div casts y to the type of x and returns x/y
// cast to type Z
func Div[X, Y, Z generics.Numeric](x X, y Y) Z {
	return Z(x / X(y))
}
