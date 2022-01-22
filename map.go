package generics

func NewMap[KEY comparable, VALUE any](size int) Map[KEY, VALUE] {
	return make(Map[KEY, VALUE], size)
}

type Map[KEY comparable, VALUE any] map[KEY]VALUE

func (m Map[KEY, VALUE]) Len() int {
	return len(m)
}

func (m Map[KEY, VALUE]) Add(key KEY, value VALUE) {
	m[key] = value
}

func (m Map[KEY, VALUE]) Delete(key KEY) {
	delete(m, key)
}

func (m Map[KEY, VALUE]) Get(key KEY) VALUE {
	return m[key]
}

func (m Map[KEY, VALUE]) Has(key KEY) bool {
	_, ok := m[key]
	return ok
}

func (m Map[KEY, VALUE]) Keys() []KEY {
	keys := make([]KEY, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[KEY, VALUE]) Values() []VALUE {
	values := make([]VALUE, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m Map[KEY, VALUE]) Entries() []struct {
	Key   KEY
	Value VALUE
} {
	entries := make([]struct {
		Key   KEY
		Value VALUE
	}, 0, len(m))
	for k, v := range m {
		entries = append(entries, struct {
			Key   KEY
			Value VALUE
		}{k, v})
	}
	return entries
}
