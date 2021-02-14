package util

// Contains returns true if the string is contained within this list.
func Contains(in []string, val string) (ret bool) {
	for _, i := range in {
		if i == val {
			ret = true
		}
	}
	return
}
