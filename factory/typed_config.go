package factory

import "context"

type TypedConfig[ID comparable] interface {
	Type() ID
}

type TypedConfigRegistry[ID comparable, CONFIG TypedConfig[ID], TARGET any] struct {
	Registry[ID, CONFIG, TARGET]
}

func NewTypedConfigRegistry[ID comparable, CONFIG TypedConfig[ID], TARGET any]() TypedConfigRegistry[ID, CONFIG, TARGET] {
	return TypedConfigRegistry[ID, CONFIG, TARGET]{
		Registry: NewRegistry[ID, CONFIG, TARGET](),
	}
}

func (r TypedConfigRegistry[ID, CONFIG, TARGET]) Build(ctx context.Context, config CONFIG) (TARGET, ReleaseFunc, error) {
	return r.Registry.Build(ctx, config.Type(), config)
}

func (r TypedConfigRegistry[ID, CONFIG, TARGET]) GetConfiguredFactory(config CONFIG) (ConfiguredFactory[TARGET], error) {
	return r.Registry.GetConfiguredFactory(config.Type(), config)
}
