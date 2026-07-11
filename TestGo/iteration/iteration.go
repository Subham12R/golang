package iteration

func Result(s string) string {
	result := ""
	for i := 0; i < 5; i++ {
		result += s
	}

	return result
}
