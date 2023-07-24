package utils

func IsUniqueArray(arr []string) bool {
	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	return len(m) == len(arr)
}
