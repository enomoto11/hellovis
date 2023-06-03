package utils

func MapSlice[T, O any](i []T, f func(T) O) []O {
	n := make([]O, len(i))
	for i, e := range i {
		n[i] = f(e)
	}
	return n
}

func MapSliceWithError[T, O any](i []T, f func(T) (O, error)) ([]O, error) {
	n := make([]O, len(i))
	for i, e := range i {
		var err error
		n[i], err = f(e)
		if err != nil {
			return nil, err
		}
	}
	return n, nil
}
