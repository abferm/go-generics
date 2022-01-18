package factory

import (
	"context"
	"fmt"

	"github.com/abferm/go-generics"
)

type ReleaseFunc = func() error

type Factory[CONFIG any, TARGET any] func(context.Context, CONFIG) (TARGET, ReleaseFunc, error)
type ConfiguredFactory[TARGET any] func(context.Context) (TARGET, ReleaseFunc, error)

func NewSingletonFactory[TARGET any](singleton TARGET) ConfiguredFactory[TARGET] {
	return func(context.Context) (TARGET, ReleaseFunc, error) {
		return singleton, func() error { return nil }, nil
	}
}

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

func (r Registry[ID, CONFIG, TARGET]) Build(ctx context.Context, id ID, config CONFIG) (TARGET, ReleaseFunc, error) {
	// zil is a zero value for the target to return in case of an error
	zil := generics.Zero[TARGET]()
	f, ok := r[id]
	if !ok {
		return zil, nil, fmt.Errorf("no factory registered for ID %v", id)
	}

	return f(ctx, config)
}

func (r Registry[ID, CONFIG, TARGET]) GetConfiguredFactory(id ID, config CONFIG) (ConfiguredFactory[TARGET], error) {
	f, ok := r[id]
	if !ok {
		return nil, fmt.Errorf("no factory registered for ID %v", id)
	}

	return func(ctx context.Context) (TARGET, ReleaseFunc, error) {
		return f(ctx, config)
	}, nil
}

type contextKey[TARGET any] struct{}

// ContextWithFactory attaches a ConfiguredFactory to the context to be passed to resources that use FromContext
// 			Incase of an error, the returned context will be the provided context, it is always safe to overwrite
//			the provided context with the resulting context.
func ContextWithFactory[TARGET any](ctx context.Context, factory ConfiguredFactory[TARGET]) (context.Context, error) {
	val := ctx.Value(contextKey[TARGET]{})
	if val != nil {
		return ctx, fmt.Errorf("%T already registered on context", factory)
	}
	return context.WithValue(ctx, contextKey[TARGET]{}, factory), nil
}

// FromContext fetches the ConfiguredFactory stored on the context
func FromContext[TARGET any](ctx context.Context) (ConfiguredFactory[TARGET], error) {
	var f ConfiguredFactory[TARGET]
	val := ctx.Value(contextKey[TARGET]{})
	if val == nil {
		return nil, fmt.Errorf("no %T registered on context", f)
	}
	f, ok := val.(ConfiguredFactory[TARGET])
	if !ok {
		return nil, fmt.Errorf("value for %T was not a %T: %T", contextKey[TARGET]{}, f, val)
	}
	return f, nil
}
