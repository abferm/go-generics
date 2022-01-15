package factory

import (
	"fmt"

	"github.com/abferm/go-generics"
)

type ReleaseFunc = func() error

type Factory[CONFIG any, TARGET any] func(CONFIG) (TARGET, ReleaseFunc, error)

type Registry[ID comparable, CONFIG any, TARGET any] map[ID]Factory[CONFIG, TARGET]

func NewRegistry[ID comparable, CONFIG any, TARGET any]() Registry[ID, CONFIG, TARGET] {
	return make(Registry[ID, CONFIG, TARGET])
}

func (r Registry[ID, CONFIG, TARGET]) Register(id ID, factory Factory[CONFIG, TARGET]) error {
	_, conflict := r[id]
	if conflict {
		return fmt.Errorf("factory already registered for %v", id)
	}

	r[id] = factory

	return nil
}

func (r Registry[ID, CONFIG, TARGET]) Deregister(id ID) {
	delete(r, id)
}

func (r Registry[ID, CONFIG, TARGET]) Build(id ID, config CONFIG) (TARGET, ReleaseFunc, error) {
	// zil is a zero value for the target to return in case of an error
	zil := generics.Zero[TARGET]()
	f, ok := r[id]
	if !ok {
		return zil, nil, fmt.Errorf("No factory registered for ID %v", id)
	}

	return f(config)
}
