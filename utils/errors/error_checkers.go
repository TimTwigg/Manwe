package error_utils

func ValidateKeyExistance[T comparable](dict map[T]any, keys []T) *T {
	for _, key := range keys {
		if _, exists := dict[key]; !exists {
			return &key
		}
	}
	return nil
}
