package letter

func ConcurrentFrequency(input []string) FreqMap {

	ch := make(chan FreqMap)

	for _, s := range input {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	output := <-ch

	for i := 1; i < len(input); i++ {
		for k, v := range <-ch {
			output[k] += v
		}
	}

	return output
}