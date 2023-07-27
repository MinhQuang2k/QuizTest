package utils

func FindUint(a []uint, x uint) uint {
	for _, n := range a {
		if x == n {
			return x
		}
	}
	return 0
}
