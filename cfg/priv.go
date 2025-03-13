package cfg

import "github.com/jgilman1337/gotils/cfg/marshaler"

// Merges two marshaler arrays, removing any duplicates (identical path values).
func mergeMarshalerArrays(m, n []marshaler.Marshaler) []marshaler.Marshaler {
	//Create a map to act as a set; pre-allocated to save on resource usage
	seen := make(map[string]struct{}, len(m)+len(n))

	//Add existing items to the map
	for _, item := range m {
		seen[item.Path()] = struct{}{}
	}

	//Append unique items from n to m
	for _, item := range n {
		if _, exists := seen[item.Path()]; !exists {
			m = append(m, item)
			seen[item.Path()] = struct{}{}
		}
	}

	return m
}
