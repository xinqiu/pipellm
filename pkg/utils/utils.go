package utils

// MergeKVListToMap merges the list of keys and values to the map
func MergeKVListToMap(keys, values []string) map[string]string {
	if len(keys) != len(values) {
		return nil
	}
	ret := make(map[string]string, 0)

	for i := range keys {
		ret[keys[i]] = values[i]
	}
	return ret
}
