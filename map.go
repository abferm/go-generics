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
