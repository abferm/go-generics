package generics

type Float interface {
	~float32 | ~float64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type SignedNumber interface {
	Float | Int
}

type Numeric interface {
	Float | Uint | Int
}

type String interface {
	~string
}
