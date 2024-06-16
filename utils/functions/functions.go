package utils

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func GetOptional[T any](dict map[string]any, key string, default_value T) T {
	switch any(default_value).(type) {
	case int:
		if value, ok := dict[key]; ok {
			return any(int(any(value).(float64))).(T)
		}
	default:
		if value, ok := dict[key]; ok {
			return value.(T)
		}
	}
	return default_value
}
