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

// Unique returns a unique array of strings.
func Unqiue(in []string) []string {
	out := []string{}
	for _, i := range in {
		if !Contains(out, i) {
			out = append(out, i)
		}
	}
	return out
}
