package utils

func Reverse(b []byte) (result []byte) {
	result = make([]byte, 0)
	for i := len(b) - 1; i >= 0; i-- {
		result = append(result, b[i])
	}
	return result
}
