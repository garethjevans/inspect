package util

// AllKeys returns a list of all unique keys from all maps.
func AllKeys(maps ...map[string]string) []string {
	keys := []string{}
	for _, m := range maps {
		for k := range m {
			if !Contains(keys, k) {
				keys = append(keys, k)
			}
		}
	}
	return keys
}
