package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("The strings must be the same length.")
		return 0, err
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
